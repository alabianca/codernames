package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alabianca/codernames/core"
	"github.com/alabianca/codernames/core/bolt"
	"github.com/alabianca/codernames/core/client"
	"github.com/alabianca/codernames/core/pb"
	"github.com/alabianca/codernames/core/server"
	"github.com/alabianca/codernames/core/ui"
	"google.golang.org/grpc"
	boltdb "github.com/boltdb/bolt"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	flagServe        *bool
	flagCreate       *bool
	flagJoin         *string
	flagJoinAsMaster *string
	flagSource       *string
)

func main() {
	flagServe = flag.Bool("serve", false, "If this is set the server will start")
	flagCreate = flag.Bool("create", false, "Create a game and join it as a spy master")
	flagJoinAsMaster = flag.String("master", "", "join a given game as a spy master")
	flagJoin = flag.String("join", "", "join a given game")
	flagSource = flag.String("source", "", "Use words at the given path as the words being used in the game")
	flag.Parse()

	if *flagServe {
		os.Exit(server_main())
	} else {
		os.Exit(client_main(*flagCreate, *flagJoin, *flagJoinAsMaster, *flagSource))
	}
}

func client_main(create bool, join, joinAsMaster, source string) int {
	// open up the connection to the server
	// @todo add security
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Error dialing %s\n", err)
		return 1
	}

	defer conn.Close()

	grpcClient := pb.NewGameServiceClient(conn)

	var file *os.File
	if create {
		var err error
		// try to open a file of words that was provided.
		// @todo provide a default?
		file, err = os.Open(source)
		if err != nil {
			log.Printf("Could not open provided file %v\n", err)
			return 1
		}

		defer file.Close()
	}

	// build up the client options.
	var clientOpts []client.Opt
	if create {
		clientOpts = append(clientOpts, client.CreateGame(), client.SourceReader(file))
	} else {
		var roomId string
		if len(join) > 0 {
			roomId = join
		} else {
			roomId = joinAsMaster
		}
		clientOpts = append(clientOpts, client.JoinId(roomId, len(joinAsMaster) > 0))
	}

	// create the client.
	// the client is responsible for interacting with the grpc server
	c := client.NewClient(
		grpcClient,
		clientOpts...,
	)

	// create the board
	board := ui.NewBoard(create)
	// the done channel is closed once the ui is closed.
	// closing the done channel will end the client.Run goroutine
	done := make(chan struct{})
	// renderc is a channel for the ui goroutine and the client's run goroutine to communicate.
	// every time a client sends a message to the rencerc channel, the ui is re-rendered with the updates
	renderc := make(chan core.Card)
	// gameCreated provides information to the ui about the game id for this particular game
	gameCreated := make(chan string, 1)
	// the keys channel is a way for the ui to communictate with the client. messages received by keys
	// are words that are displayed in the board. The client will use the message and send an UpdateGame rpc with it
	keys := make(chan string)

	go board.Render(renderc, gameCreated, keys)
	go c.Run(context.Background(), renderc, keys, gameCreated, done)



	<-board.Done()
	close(done)

	return 0
}

func server_main() int {
	db, err := boltdb.Open("codernames.db", 0600, nil)
	if err != nil {
		log.Printf("Error Creating Database %s", err)
		return 1
	}
	defer db.Close()

	// create the games bucket
	err = db.Update(func(tx *boltdb.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("games"))
		return err
	})
	if err != nil {
		log.Printf("Error Creating the Games bucket %s", err)
		return 1
	}
	//// set up mongodb connection
	//client, err := mongod.NewClient(options.Client().ApplyURI(mongoURI()))
	//if err != nil {
	//	log.Printf("Error Creating Database client (mongodb) %v\n", err)
	//	return 1
	//}
	//
	//defer client.Disconnect(context.Background())
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	//log.Println("Connecting to database")
	//if err := client.Connect(ctx); err != nil {
	//	log.Printf("Could not connect to database (mongodb) %v\n", err)
	//	return 1
	//}
	//
	//if err := client.Ping(context.Background(), nil); err != nil {
	//	log.Fatal("Could not connect to mongodb with provided server selection timeout")
	//	return 1
	//}
	//
	//collection := client.Database("codernames").Collection("games")

	// set up the tcp listener required by the grpc server later
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("Error trying to listen %v\n", err)
		return 1
	}

	// set up the grpc server and service
	grpcSrv := grpc.NewServer()
	//var dal mongo.GameDAL
	var dal bolt.GameDAL
	dal.BucketName = "games"
	dal.DB = db
	//dal.Collection = collection
	service := server.NewGameService(&dal)
	pb.RegisterGameServiceServer(grpcSrv, service)

	go func() {
		log.Println("Listening ...")
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("Error serving %v\n", err)
		}
	}()

	// wait for ctrl-c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping server")
	grpcSrv.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Stopped")

	return 0

}

func mongoURI() string {
	host := getEnv("db_host", "localhost")
	port := getEnv("db_port", "27017")

	return fmt.Sprintf("mongodb://%s:%s", host, port)
}

func getEnv(key, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}

	return env
}

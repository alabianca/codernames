// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/pb/game.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Card_Type int32

const (
	Card_BLUE     Card_Type = 0
	Card_RED      Card_Type = 1
	Card_ASSASSIN Card_Type = 2
	Card_CIVILIAN Card_Type = 3
)

var Card_Type_name = map[int32]string{
	0: "BLUE",
	1: "RED",
	2: "ASSASSIN",
	3: "CIVILIAN",
}

var Card_Type_value = map[string]int32{
	"BLUE":     0,
	"RED":      1,
	"ASSASSIN": 2,
	"CIVILIAN": 3,
}

func (x Card_Type) String() string {
	return proto.EnumName(Card_Type_name, int32(x))
}

func (Card_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{0, 0}
}

type Card struct {
	X1                   int32     `protobuf:"varint,1,opt,name=X1,proto3" json:"X1,omitempty"`
	Y1                   int32     `protobuf:"varint,2,opt,name=Y1,proto3" json:"Y1,omitempty"`
	X2                   int32     `protobuf:"varint,3,opt,name=X2,proto3" json:"X2,omitempty"`
	Y2                   int32     `protobuf:"varint,4,opt,name=Y2,proto3" json:"Y2,omitempty"`
	Active               bool      `protobuf:"varint,5,opt,name=Active,proto3" json:"Active,omitempty"`
	Content              string    `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	Type                 Card_Type `protobuf:"varint,7,opt,name=type,proto3,enum=game.Card_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{0}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetX1() int32 {
	if m != nil {
		return m.X1
	}
	return 0
}

func (m *Card) GetY1() int32 {
	if m != nil {
		return m.Y1
	}
	return 0
}

func (m *Card) GetX2() int32 {
	if m != nil {
		return m.X2
	}
	return 0
}

func (m *Card) GetY2() int32 {
	if m != nil {
		return m.Y2
	}
	return 0
}

func (m *Card) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Card) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Card) GetType() Card_Type {
	if m != nil {
		return m.Type
	}
	return Card_BLUE
}

type Game struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Cards                []*Card  `protobuf:"bytes,2,rep,name=cards,proto3" json:"cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{1}
}

func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Game) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

type CreateGameRequest struct {
	Game                 *Game    `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGameRequest) Reset()         { *m = CreateGameRequest{} }
func (m *CreateGameRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGameRequest) ProtoMessage()    {}
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{2}
}

func (m *CreateGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGameRequest.Unmarshal(m, b)
}
func (m *CreateGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGameRequest.Marshal(b, m, deterministic)
}
func (m *CreateGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGameRequest.Merge(m, src)
}
func (m *CreateGameRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGameRequest.Size(m)
}
func (m *CreateGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGameRequest proto.InternalMessageInfo

func (m *CreateGameRequest) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type CreateGameResponse struct {
	Game                 *Game    `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGameResponse) Reset()         { *m = CreateGameResponse{} }
func (m *CreateGameResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGameResponse) ProtoMessage()    {}
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{3}
}

func (m *CreateGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGameResponse.Unmarshal(m, b)
}
func (m *CreateGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGameResponse.Marshal(b, m, deterministic)
}
func (m *CreateGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGameResponse.Merge(m, src)
}
func (m *CreateGameResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGameResponse.Size(m)
}
func (m *CreateGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGameResponse proto.InternalMessageInfo

func (m *CreateGameResponse) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type JoinGameRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGameRequest) Reset()         { *m = JoinGameRequest{} }
func (m *JoinGameRequest) String() string { return proto.CompactTextString(m) }
func (*JoinGameRequest) ProtoMessage()    {}
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{4}
}

func (m *JoinGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGameRequest.Unmarshal(m, b)
}
func (m *JoinGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGameRequest.Marshal(b, m, deterministic)
}
func (m *JoinGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGameRequest.Merge(m, src)
}
func (m *JoinGameRequest) XXX_Size() int {
	return xxx_messageInfo_JoinGameRequest.Size(m)
}
func (m *JoinGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGameRequest proto.InternalMessageInfo

func (m *JoinGameRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JoinGameResponse struct {
	Game                 *Game    `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGameResponse) Reset()         { *m = JoinGameResponse{} }
func (m *JoinGameResponse) String() string { return proto.CompactTextString(m) }
func (*JoinGameResponse) ProtoMessage()    {}
func (*JoinGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{5}
}

func (m *JoinGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGameResponse.Unmarshal(m, b)
}
func (m *JoinGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGameResponse.Marshal(b, m, deterministic)
}
func (m *JoinGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGameResponse.Merge(m, src)
}
func (m *JoinGameResponse) XXX_Size() int {
	return xxx_messageInfo_JoinGameResponse.Size(m)
}
func (m *JoinGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGameResponse proto.InternalMessageInfo

func (m *JoinGameResponse) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type UpdateGameRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGameRequest) Reset()         { *m = UpdateGameRequest{} }
func (m *UpdateGameRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGameRequest) ProtoMessage()    {}
func (*UpdateGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{6}
}

func (m *UpdateGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGameRequest.Unmarshal(m, b)
}
func (m *UpdateGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGameRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGameRequest.Merge(m, src)
}
func (m *UpdateGameRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGameRequest.Size(m)
}
func (m *UpdateGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGameRequest proto.InternalMessageInfo

func (m *UpdateGameRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateGameRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdateGameResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGameResponse) Reset()         { *m = UpdateGameResponse{} }
func (m *UpdateGameResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateGameResponse) ProtoMessage()    {}
func (*UpdateGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d922638272626698, []int{7}
}

func (m *UpdateGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGameResponse.Unmarshal(m, b)
}
func (m *UpdateGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGameResponse.Marshal(b, m, deterministic)
}
func (m *UpdateGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGameResponse.Merge(m, src)
}
func (m *UpdateGameResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateGameResponse.Size(m)
}
func (m *UpdateGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGameResponse proto.InternalMessageInfo

func (m *UpdateGameResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("game.Card_Type", Card_Type_name, Card_Type_value)
	proto.RegisterType((*Card)(nil), "game.Card")
	proto.RegisterType((*Game)(nil), "game.Game")
	proto.RegisterType((*CreateGameRequest)(nil), "game.CreateGameRequest")
	proto.RegisterType((*CreateGameResponse)(nil), "game.CreateGameResponse")
	proto.RegisterType((*JoinGameRequest)(nil), "game.JoinGameRequest")
	proto.RegisterType((*JoinGameResponse)(nil), "game.JoinGameResponse")
	proto.RegisterType((*UpdateGameRequest)(nil), "game.UpdateGameRequest")
	proto.RegisterType((*UpdateGameResponse)(nil), "game.UpdateGameResponse")
}

func init() { proto.RegisterFile("core/pb/game.proto", fileDescriptor_d922638272626698) }

var fileDescriptor_d922638272626698 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x14, 0x5c, 0x7b, 0xbd, 0x1f, 0x7d, 0x45, 0x6d, 0xfa, 0x24, 0x8a, 0xd5, 0x03, 0x0a, 0x86, 0x43,
	0x4e, 0x5b, 0xd6, 0x05, 0x89, 0x4b, 0x0f, 0xe9, 0x52, 0xa1, 0xa0, 0xaa, 0x07, 0x87, 0xa2, 0x2d,
	0xb7, 0x6c, 0x62, 0xa1, 0x3d, 0x34, 0x09, 0x49, 0xa8, 0xb4, 0x7f, 0x13, 0xf1, 0x83, 0x90, 0x9d,
	0x84, 0x84, 0xec, 0x81, 0xde, 0x32, 0x99, 0x37, 0x7e, 0x33, 0x63, 0x19, 0x30, 0xce, 0x0a, 0x7d,
	0x9e, 0x6f, 0xce, 0xbf, 0x47, 0x0f, 0x7a, 0x91, 0x17, 0x59, 0x95, 0x21, 0x33, 0xdf, 0xe2, 0x37,
	0x01, 0xb6, 0x8a, 0x8a, 0x04, 0x8f, 0x80, 0xae, 0x97, 0x9c, 0xb8, 0xc4, 0x9b, 0x28, 0xba, 0x5e,
	0x1a, 0x7c, 0xbf, 0xe4, 0xb4, 0xc6, 0xf7, 0x16, 0xaf, 0x25, 0x1f, 0x37, 0xbc, 0xb4, 0xbc, 0xe4,
	0xac, 0xe1, 0x25, 0x9e, 0xc2, 0xd4, 0x8f, 0xab, 0xed, 0xa3, 0xe6, 0x13, 0x97, 0x78, 0x73, 0xd5,
	0x20, 0xe4, 0x30, 0x5b, 0x65, 0x69, 0xa5, 0xd3, 0x8a, 0x4f, 0x5d, 0xe2, 0x1d, 0xa8, 0x16, 0xe2,
	0x6b, 0x60, 0xd5, 0x2e, 0xd7, 0x7c, 0xe6, 0x12, 0xef, 0x48, 0x1e, 0x2f, 0xac, 0x37, 0xe3, 0x65,
	0xf1, 0x65, 0x97, 0x6b, 0x65, 0x49, 0xf1, 0x1e, 0x98, 0x41, 0x38, 0x07, 0x76, 0x75, 0x73, 0x77,
	0xed, 0x8c, 0x70, 0x06, 0x63, 0x75, 0xfd, 0xd1, 0x21, 0xf8, 0x0c, 0xe6, 0x7e, 0x18, 0xfa, 0x61,
	0x18, 0xdc, 0x3a, 0xd4, 0xa0, 0x55, 0xf0, 0x35, 0xb8, 0x09, 0xfc, 0x5b, 0x67, 0x2c, 0x3e, 0x00,
	0xfb, 0x14, 0x3d, 0x68, 0xe3, 0x32, 0x48, 0x6c, 0xaa, 0x03, 0x45, 0x83, 0x04, 0x5d, 0x98, 0xc4,
	0x51, 0x91, 0x94, 0x9c, 0xba, 0x63, 0xef, 0x50, 0x42, 0xb7, 0x54, 0xd5, 0x84, 0xb8, 0x80, 0x93,
	0x55, 0xa1, 0xa3, 0x4a, 0x1b, 0xbd, 0xd2, 0x3f, 0x7e, 0xea, 0xb2, 0xc2, 0x97, 0x60, 0xdb, 0xb2,
	0x07, 0xfd, 0x55, 0xd9, 0x81, 0xba, 0xc5, 0x77, 0x80, 0x7d, 0x51, 0x99, 0x67, 0x69, 0xa9, 0xff,
	0xab, 0x7a, 0x05, 0xc7, 0x9f, 0xb3, 0x6d, 0xda, 0x5f, 0x34, 0xf0, 0x2b, 0x24, 0x38, 0xdd, 0xc8,
	0x13, 0x8f, 0xbd, 0x84, 0x93, 0xbb, 0x3c, 0x19, 0x24, 0x18, 0x16, 0xd1, 0xbb, 0x16, 0xfa, 0xcf,
	0xb5, 0x88, 0x37, 0x80, 0x7d, 0x79, 0xb3, 0x74, 0xa0, 0x97, 0xbf, 0x08, 0x1c, 0x9a, 0x81, 0x50,
	0x17, 0x8f, 0xdb, 0x58, 0xa3, 0x0f, 0xd0, 0x35, 0x80, 0x2f, 0x9a, 0x5e, 0x87, 0x45, 0x9e, 0xf1,
	0x7d, 0xa2, 0x5e, 0x20, 0x46, 0x78, 0x09, 0xf3, 0x36, 0x2b, 0x3e, 0xaf, 0xe7, 0x06, 0xf5, 0x9c,
	0x9d, 0x0e, 0x7f, 0xb7, 0xe2, 0xb7, 0xc4, 0x38, 0xe8, 0x7c, 0xb7, 0x0e, 0xf6, 0x8a, 0x68, 0x1d,
	0xec, 0x47, 0x14, 0xa3, 0x2b, 0xf6, 0x8d, 0xe6, 0x9b, 0xcd, 0xd4, 0xbe, 0x8f, 0x8b, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0xa0, 0xe7, 0x5c, 0x35, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (GameService_JoinGameClient, error)
	UpdateGame(ctx context.Context, in *UpdateGameRequest, opts ...grpc.CallOption) (*UpdateGameResponse, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (GameService_JoinGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GameService_serviceDesc.Streams[0], "/game.GameService/JoinGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceJoinGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_JoinGameClient interface {
	Recv() (*JoinGameResponse, error)
	grpc.ClientStream
}

type gameServiceJoinGameClient struct {
	grpc.ClientStream
}

func (x *gameServiceJoinGameClient) Recv() (*JoinGameResponse, error) {
	m := new(JoinGameResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) UpdateGame(ctx context.Context, in *UpdateGameRequest, opts ...grpc.CallOption) (*UpdateGameResponse, error) {
	out := new(UpdateGameResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/UpdateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error)
	JoinGame(*JoinGameRequest, GameService_JoinGameServer) error
	UpdateGame(context.Context, *UpdateGameRequest) (*UpdateGameResponse, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) CreateGame(ctx context.Context, req *CreateGameRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (*UnimplementedGameServiceServer) JoinGame(req *JoinGameRequest, srv GameService_JoinGameServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (*UnimplementedGameServiceServer) UpdateGame(ctx context.Context, req *UpdateGameRequest) (*UpdateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGame not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_JoinGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinGameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).JoinGame(m, &gameServiceJoinGameServer{stream})
}

type GameService_JoinGameServer interface {
	Send(*JoinGameResponse) error
	grpc.ServerStream
}

type gameServiceJoinGameServer struct {
	grpc.ServerStream
}

func (x *gameServiceJoinGameServer) Send(m *JoinGameResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GameService_UpdateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).UpdateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/UpdateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).UpdateGame(ctx, req.(*UpdateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _GameService_CreateGame_Handler,
		},
		{
			MethodName: "UpdateGame",
			Handler:    _GameService_UpdateGame_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinGame",
			Handler:       _GameService_JoinGame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "core/pb/game.proto",
}
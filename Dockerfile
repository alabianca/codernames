FROM golang:1.14.4-alpine3.12 as builder

RUN apk add --no-cache git mercurial
# add gcc and g++ ro run tests
RUN apk add --update gcc
RUN apk add --update g++

RUN git clone https://github.com/alabianca/codernames.git
WORKDIR codernames/cmd/codernames

COPY . .

RUN go build -o codernames

CMD ["./codernames", "-serve"]
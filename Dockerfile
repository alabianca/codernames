FROM golang:1.14.4-alpine3.12 as builder

RUN apk add --no-cache git mercurial
# add gcc and g++ ro run tests
RUN apk add --update gcc
RUN apk add --update g++

RUN go get github.com/alabianca/codernames
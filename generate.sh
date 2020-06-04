#!/bin/bash

protoc core/pb/game.proto --go_out=plugins=grpc:.
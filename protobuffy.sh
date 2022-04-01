#!/bin/bash
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

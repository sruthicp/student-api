#!/bin/bash

protoc -I proto/ proto/student.proto --go-grpc_out=proto/student --go_out=proto/student --go-grpc_opt=require_unimplemented_servers=false
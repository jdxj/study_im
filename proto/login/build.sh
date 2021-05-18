#!/usr/bin/env bash

proto_path='/home/jdxj/workspace/study_im/proto'
protoc --proto_path=${proto_path}:. --micro_out=. --go_out=. login.proto

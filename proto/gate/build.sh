#!/usr/bin/env bash

protoc --proto_path=/home/jdxj/workspace/study_im/proto:. --micro_out=. --go_out=. gate.proto

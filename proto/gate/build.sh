#!/usr/bin/env bash

protoc --proto_path=. --micro_out=. --go_out=. gate.proto
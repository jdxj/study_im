package main

import (
	"github.com/jdxj/study_im/proto"
	"github.com/name5566/leaf/network/protobuf"
)

func registerMsg() *protobuf.Processor {
	p := protobuf.NewProcessor()
	for _, msg := range proto.SortedMsg() {
		p.Register(msg)
	}
	return p
}

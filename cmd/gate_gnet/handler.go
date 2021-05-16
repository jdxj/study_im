package main

import (
	"fmt"
	"time"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
	"github.com/jdxj/study_im/proto/head"
)

func handle(agent *Agent, msg interface{}) []byte {
	var data []byte
	var err error
	switch v := msg.(type) {
	case *head.Head:
		data, err = handleHead(agent, v)
	}

	if err != nil {
		logger.Errorf("handle: %s", err)
	}
	return data
}

func handleHead(agent *Agent, h *head.Head) ([]byte, error) {
	fmt.Printf("req: %s\n", h)
	h.Seq += 1
	h.Timestamp = time.Now().Unix()
	return protobuf.Marshal(h)
}

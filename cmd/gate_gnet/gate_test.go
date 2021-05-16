package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/proto/head"
)

func run(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer conn.Close()

	// 编码内容
	h := &head.Head{
		Version:   1,
		Seq:       1,
		Timestamp: time.Now().Unix(),
	}
	data, err := protobuf.Marshal(h)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	// 封帧
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(len(data)))
	_, err = conn.Write(lenBuf)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	_, err = conn.Write(data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	// 解帧
	_, err = io.ReadFull(conn, lenBuf)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	length := binary.BigEndian.Uint32(lenBuf)
	data = make([]byte, length)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	// 解码内容
	_, msg, err := protobuf.Unmarshal(data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("resp: %s\n", msg)
}

func TestGate_Serve(t *testing.T) {
	run(t)
}

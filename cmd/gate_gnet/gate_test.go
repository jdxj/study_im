package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"

	"github.com/jdxj/study_im/proto/login"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/proto/head"
)

func GenerateFrame(data []byte) []byte {
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(len(data)))

	frameBuf := make([]byte, 4+len(data))
	copy(frameBuf, lenBuf)
	copy(frameBuf[4:], data)
	return frameBuf
}

func ResolveFrame(r io.Reader) ([]byte, error) {
	lenBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lenBuf)
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lenBuf)
	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	return data, err
}

var (
	conn net.Conn
)

func TestMain(m *testing.M) {
	var err error
	conn, err = NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

func NewClient() (net.Conn, error) {
	return net.Dial("tcp", "127.0.0.1:9000")
}

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

func TestAuthReq(t *testing.T) {
	req := &login.AuthRequest{
		Token: "abc",
		Uid:   123,
	}
	data, err := protobuf.Marshal(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	frame := GenerateFrame(data)
	_, err = conn.Write(frame)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	data, err = ResolveFrame(conn)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	_, msg, err := protobuf.Unmarshal(data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", msg)
}

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/proto/head"

	"github.com/panjf2000/gnet"
)

type MyCodec struct {
	num int
}

// len | data

func (mc *MyCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	length := len(buf)

	data := make([]byte, 4+length)
	binary.BigEndian.PutUint32(data, uint32(length))
	copy(data[4:], buf)
	return data, nil
}

func (mc *MyCodec) Decode(c gnet.Conn) ([]byte, error) {
	mc.num++
	fmt.Printf("decode called: %d, num: %d\n", time.Now().UnixNano(), mc.num)
	size, buf := c.ReadN(4)
	if size != 4 {
		return nil, fmt.Errorf("abc")
	}

	length := binary.BigEndian.Uint32(buf)
	size, buf = c.ReadN(4 + int(length))
	if size != 4+int(length) {
		return nil, fmt.Errorf("data len not enough")
	}
	c.ShiftN(size)

	// todo: 不知道是否有 GC 问题
	return buf[4:], nil
}

type echoServer struct {
	*gnet.EventServer
	p *protobuf.Processor
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("es pointer: %p\n", es)
	fmt.Printf("receive: %s\n", frame)
	_, msg, err := es.p.Unmarshal(frame)
	if err != nil {
		fmt.Printf("Unmarshal: %s\n", err)
		return nil, 0
	}

	head := msg.(*head.Head)
	head.Seq++
	data, err := es.p.Marshal(head)
	if err != nil {
		fmt.Printf("Marshal: %s\n", err)
		return nil, 0
	}
	return data, 0
}

func StartServer() {
	echo := new(echoServer)
	echo.p = protobuf.NewProcessor()
	echo.p.Register(0, &head.Head{})
	echo.EventServer.React(nil, nil)

	err := gnet.Serve(echo, "tcp://:9000",
		gnet.WithMulticore(true),
		gnet.WithCodec(&MyCodec{}))
	if err != nil {
		log.Println(err)
	}
}

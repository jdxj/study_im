package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("%v\n", frame)
	v := binary.LittleEndian.Uint32(frame)
	fmt.Println(v)
	out = frame
	return
}

func StartServer() {
	echo := new(echoServer)
	err := gnet.Serve(echo, "tcp://:9000",
		gnet.WithMulticore(true),
		gnet.WithCodec(&gnet.LineBasedFrameCodec{}))
	if err != nil {
		log.Println(err)
	}
}

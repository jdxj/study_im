package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"

	"github.com/jdxj/study_im/proto/head"
	"github.com/jdxj/study_im/proto/protobuf"
)

type client struct {
	conn net.Conn
}

func (c *client) Connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	p := protobuf.NewProcessor()
	p.Register(0, &head.Head{})

	head := &head.Head{
		Version:   1,
		Seq:       0,
		Timestamp: time.Now().Unix(),
	}
	data, err := p.Marshal(head)
	if err != nil {
		log.Fatalln(err)
	}

	mc := &MyCodec{}
	data, _ = mc.Encode(nil, data)
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}

	lenBuf := make([]byte, 4)
	_, err = io.ReadFull(conn, lenBuf)
	if err != nil {
		log.Fatalln(err)
	}

	length := binary.BigEndian.Uint32(lenBuf)
	frameBuf := make([]byte, length)
	_, err = io.ReadFull(conn, frameBuf)
	if err != nil {
		log.Fatalln(err)
	}

	_, msg, err := p.Unmarshal(frameBuf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("resp: %s\n", msg)

	time.Sleep(time.Minute)
}

func (c *client) read() {
	buf := make([]byte, 1024)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("client receive: %s\n, %v", buf[:n], buf[:n])
	}
}

func StartClient() {
	c := &client{}
	c.Connect()
}

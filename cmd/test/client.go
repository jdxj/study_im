package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
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

	// write "hello" to server
	content := []byte("hello")
	length := len(content)
	data := make([]byte, 4+length)
	binary.BigEndian.PutUint32(data, uint32(length))
	copy(data[4:], content)
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}

	// read response from server
	data = data[:4]
	n, err := io.ReadFull(conn, data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("read length1: %d\n", n)

	length = int(binary.BigEndian.Uint32(data))
	data = make([]byte, length)
	n, err = io.ReadFull(conn, data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("read length2: %d\n", n)
	log.Printf("result: %s\n", data)
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

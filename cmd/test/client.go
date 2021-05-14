package main

import (
	"log"
	"net"
	"strconv"
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

	c.conn = conn
	go c.read()

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		data := strconv.Itoa(i) + "\n"
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Fatalln(conn)
		}
	}
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

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

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
	return buf, nil
}

func readN(c gnet.Conn, size int) []byte {
	fmt.Printf("want to read: %d\n", size)
	left := size
	n := 0
	var buf []byte
	data := make([]byte, 0, size)

	for ; left > 0; left -= n {
		n, buf = c.ReadN(left)
		fmt.Printf("read: %d\n", n)
		time.Sleep(2 * time.Second)
		c.ShiftN(n)
		data = append(data, buf...)

	}
	return data
}

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("es pointer: %p\n", es)
	fmt.Printf("receive: %s\n", frame)
	res := fmt.Sprintf("%s %s", frame, "world")
	out = []byte(res)
	return
}

func StartServer() {
	echo := new(echoServer)
	err := gnet.Serve(echo, "tcp://:9000",
		gnet.WithMulticore(true),
		gnet.WithCodec(&MyCodec{}))
	if err != nil {
		log.Println(err)
	}
}

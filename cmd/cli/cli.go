package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/jdxj/study_im/codec/protobuf"
)

func NewCli() *Cli {
	addr := flag.String("addr", "127.0.0.1:9000", "server address")
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

	return &Cli{conn: conn}
}

type Cli struct {
	conn net.Conn
}

func (cli *Cli) ReadLoop() {
	for {
		data, err := ResolveFrame(cli.conn)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}

		_, msg, err := protobuf.Unmarshal(data)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}
		log.Printf("%s\n", msg)
	}
}

func (cli *Cli) WriteLoop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		// \n: LF 10
		line, err := reader.ReadString(10)
		if err != nil {
			log.Printf("ReadString: %s\n", err)
			continue
		}
		line = strings.ReplaceAll(line, "\n", "")
		if len(line) == 0 {
			continue
		}

		args := strings.Split(line, " ")
		parser, ok := commands[args[0]]
		if !ok {
			log.Printf("cmd '%s' not found\n", args[0])
			continue
		}

		data, err := parser.Parse(args[1:])
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}
		if data == nil {
			continue
		}

		frame := GenerateFrame(data)
		_, err = cli.conn.Write(frame)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}
	}
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

func GenerateFrame(data []byte) []byte {
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(len(data)))

	frameBuf := make([]byte, 4+len(data))
	copy(frameBuf, lenBuf)
	copy(frameBuf[4:], data)
	return frameBuf
}

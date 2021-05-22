package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jdxj/study_im/codec/protobuf"
)

func NewCli() *Cli {
	addr := flag.String("addr", "127.0.0.1:9000", "server address")
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

	return &Cli{
		conn:     conn,
		sendQ:    make(chan interface{}, 1000),
		sendSign: make(chan struct{}),
		ackQ:     make(map[uint32]struct{}),
	}
}

type Cli struct {
	conn net.Conn

	sendQ chan interface{}

	seq      uint32
	sendSign chan struct{}

	ackQMutex sync.RWMutex
	ackQ      map[uint32]struct{}

	token  string
	userID uint32
}

func (cli *Cli) ReadLoop() {
	for {
		data, err := ResolveFrame(cli.conn)
		if err != nil {
			if err == io.EOF {
				log.Fatalln("server already stopped")
			}
			log.Printf("%s\n", err)
			continue
		}

		rawMsg, err := protobuf.Unmarshal(data)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}
		cli.handle(rawMsg)
	}
}

func (cli *Cli) handle(rawMsg *protobuf.RawMsg) {
	cli.ackMsg(rawMsg.Ack)

	// todo: 处理各种状态
	log.Printf("%s\n\n", rawMsg)
}

func (cli *Cli) ackMsg(ack uint32) {
	cli.ackQMutex.Lock()
	_, ok := cli.ackQ[ack]
	if ok {
		log.Printf("ack: %d\n", ack)
		delete(cli.ackQ, ack)
		select {
		case cli.sendSign <- struct{}{}:
		default: // 避免阻塞
		}
	}
	cli.ackQMutex.Unlock()
}

func (cli *Cli) SendMessage() {
	go cli.write()

	reader := bufio.NewReader(os.Stdin)
	for {
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

		select {
		case cli.sendQ <- data:
		default:
			log.Printf("队列已满, 稍后再试\n")
		}
	}
}

func (cli *Cli) write() {
	for {
		content := <-cli.sendQ
		seq++

		data, err := protobuf.Marshal(seq, 0, content)
		if err != nil {
			log.Printf("Marshal: %s\n", err)
			continue
		}

		// 注意: 必须先添加到 ackQ, 再发送消息
		cli.ackQMutex.Lock()
		cli.ackQ[seq] = struct{}{}
		cli.ackQMutex.Unlock()

		frame := GenerateFrame(data)
		_, err = cli.conn.Write(frame)
		if err != nil {
			log.Printf("Write: %s\n", err)

			// 撤回 ack
			cli.ackQMutex.Lock()
			delete(cli.ackQ, seq)
			cli.ackQMutex.Unlock()
			continue
		}

		// 等超时, 为了简化逻辑, 重发动作交给用户
		timer := time.NewTimer(3 * time.Second)
		select {
		case <-cli.sendSign:
		case <-timer.C:
			log.Printf("timeout, seq: %d, msg: %s\n", seq, content)
			// 撤回 ack
			cli.ackQMutex.Lock()
			delete(cli.ackQ, seq)
			cli.ackQMutex.Unlock()
		}
		timer.Stop()
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

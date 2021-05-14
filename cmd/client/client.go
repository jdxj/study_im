package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/name5566/leaf/network"
)

func NewClient() *Client {
	c := &Client{
		Processor:   registerMsg(),
		TCPAddr:     "127.0.0.1:9000",
		LenMsgLen:   4,
		sendChan:    make(chan interface{}),
		receiveChan: make(chan interface{}),
		version:     1,
	}
	return c
}

type Client struct {
	Processor network.Processor

	// tcp
	TCPAddr   string
	LenMsgLen int

	// cmd
	sendChan    chan interface{}
	receiveChan chan interface{}
	version     uint32
	seq         uint32
}

func (c *Client) Run(closeSig chan bool) {
	go c.RunCmd(closeSig)

	var tcpClient *network.TCPClient
	if c.TCPAddr != "" {
		tcpClient = new(network.TCPClient)
		tcpClient.Addr = c.TCPAddr
		tcpClient.ConnNum = 1
		tcpClient.ConnectInterval = 3 * time.Second
		tcpClient.PendingWriteNum = 100
		tcpClient.LenMsgLen = c.LenMsgLen
		tcpClient.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &agent{
				conn:        conn,
				client:      c,
				sendChan:    c.sendChan,
				receiveChan: c.receiveChan,
			}
			return a
		}
	}

	if tcpClient != nil {
		tcpClient.Start()
	}

	<-closeSig

	if tcpClient != nil {
		tcpClient.Close()
	}
}

func (c *Client) RunCmd(closeSig chan bool) {
	go c.handle()

	cli := &CmdLine{client: c}
	cli.Init()

	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-closeSig:
			log.Println("stop RunCmd")
		default:
		}

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
		msg, err := cli.Parse(args)
		if err != nil {
			log.Printf("Parse: %s\n", err)
			continue
		}
		if msg != nil {
			c.sendChan <- msg
		}
	}
}

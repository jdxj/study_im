package gate

import (
	"fmt"

	"github.com/name5566/leaf/network"
)

func New(host string, port int) *Gate {
	addr := fmt.Sprintf("%s:%d", host, port)
	gate := &Gate{
		Processor:    registerMsg(),
		TCPAddr:      addr,
		LenMsgLen:    4,
		LittleEndian: false,
	}
	return gate
}

type Gate struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor

	// tcp
	TCPAddr      string
	tcpServer    *network.TCPServer
	LenMsgLen    int
	LittleEndian bool
}

func (gate *Gate) Run() {
	tcpServer := new(network.TCPServer)
	tcpServer.Addr = gate.TCPAddr
	tcpServer.MaxConnNum = gate.MaxConnNum
	tcpServer.PendingWriteNum = gate.PendingWriteNum
	tcpServer.LenMsgLen = gate.LenMsgLen
	tcpServer.MaxMsgLen = gate.MaxMsgLen
	tcpServer.LittleEndian = gate.LittleEndian
	tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
		a := &agent{conn: conn, gate: gate}
		return a
	}
	gate.tcpServer = tcpServer
	tcpServer.Start()
}

func (gate *Gate) Stop() {
	gate.tcpServer.Close()
}

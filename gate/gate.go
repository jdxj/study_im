package gate

import (
	"github.com/name5566/leaf/network"
)

func New() *Gate {
	gate := &Gate{
		MaxConnNum:      0,
		PendingWriteNum: 0,
		MaxMsgLen:       0,
		Processor:       registerMsg(),
		TCPAddr:         ":9000",
		LenMsgLen:       4,
		LittleEndian:    false,
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
	LenMsgLen    int
	LittleEndian bool
}

func (gate *Gate) Run(closeSig chan bool) {
	var tcpServer *network.TCPServer
	if gate.TCPAddr != "" {
		tcpServer = new(network.TCPServer)
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
	}

	if tcpServer != nil {
		tcpServer.Start()
	}

	<-closeSig

	if tcpServer != nil {
		tcpServer.Close()
	}
}

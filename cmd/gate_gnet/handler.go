package main

import (
	"context"

	"github.com/jdxj/study_im/proto/chat"

	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
	pbGate "github.com/jdxj/study_im/proto/gate"
	"github.com/jdxj/study_im/proto/head"
	"github.com/jdxj/study_im/proto/login"
)

func (gate *Gate) handle(conn gnet.Conn, rawMsg *protobuf.RawMsg) []byte {
	var data []byte
	var err error
	switch rawMsg.Cmd {
	case protobuf.Heartbeat:
		data, err = gate.handleHeartbeat(conn, rawMsg)
	case protobuf.AuthRequest:
		data, err = gate.handleAuthRequest(conn, rawMsg)
	}

	if err != nil {
		logger.Errorf("handle: %s", err)
	}
	return data
}

func (gate *Gate) handleHeartbeat(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	logger.Debugf("%v", *rawMsg)
	return protobuf.Marshal(gate.sm.NextSeq(), &head.Heartbeat{})
}

func (gate *Gate) handleAuthRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	clientID, ok := conn.Context().(int64)
	if !ok { // 没有登录过
		clientID = gate.nextID()
	}

	req := rawMsg.Msg.(*login.AuthRequest)
	req.Identity = &pbGate.Identity{
		NodeId:   uint32(gate.nodeID),
		ClientId: clientID,
	}
	resp, err := loginService.Auth(context.Background(), req)
	if err != nil {
		return nil, err
	}

	gate.am.AddAgent(clientID, conn)
	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

func (gate *Gate) handleC2CSendRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	req := rawMsg.Msg.(*chat.C2CSendRequest)
	logger.Debugf("req: %s", req)

	resp := &chat.C2CSendResponse{MsgId: 888}
	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

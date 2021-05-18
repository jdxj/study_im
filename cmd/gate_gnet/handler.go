package main

import (
	"github.com/jdxj/study_im/proto/chat"

	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
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
	req := rawMsg.Msg.(*login.AuthRequest)
	logger.Debugf("req: %s", req)
	if conn.Context() != nil {
		return nil, nil
	}

	// 目前只是简单的添加到集合中, 没有认证
	// todo: 认证

	agentID := gate.nextID()
	gate.am.AddAgent(agentID, conn)

	resp := &login.AuthResponse{
		Status: 1,
		ErrMsg: "ok",
	}
	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

func (gate *Gate) handleC2CSendRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	req := rawMsg.Msg.(*chat.C2CSendRequest)
	logger.Debugf("req: %s", req)

	resp := &chat.C2CSendResponse{MsgId: 888}
	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

package main

import (
	"time"

	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
	"github.com/jdxj/study_im/proto/head"
	"github.com/jdxj/study_im/proto/login"
)

func (gate *Gate) handle(conn gnet.Conn, msg interface{}) []byte {
	var data []byte
	var err error
	switch v := msg.(type) {
	case *head.Heartbeat:
		data, err = gate.handleHeartbeat(conn, v)
	case *login.AuthRequest:
		data, err = gate.handleAuthRequest(conn, v)
	}

	if err != nil {
		logger.Errorf("handle: %s", err)
	}
	return data
}

func (gate *Gate) handleHeartbeat(conn gnet.Conn, h *head.Heartbeat) ([]byte, error) {
	logger.Debugf("req: %s", h)
	if conn.Context() == nil {
		return nil, nil
	}

	h.Seq += 1
	h.Timestamp = time.Now().Unix()
	return protobuf.Marshal(h)
}

func (gate *Gate) handleAuthRequest(conn gnet.Conn, req *login.AuthRequest) ([]byte, error) {
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
	return protobuf.Marshal(resp)
}

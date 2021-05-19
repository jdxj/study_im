package main

import (
	"context"

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
	case protobuf.LogoutRequest:
		data, err = gate.handleLogoutRequest(conn, rawMsg)
	case protobuf.C2CSendRequest:
		data, err = gate.handleC2CSendRequest(conn, rawMsg)
	default:
		logger.Warnf("not handled: %d", rawMsg.Cmd)
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
	clientID := conn.Context().(int64)
	req := rawMsg.Msg.(*login.AuthRequest)
	req.Identity = gate.identify(clientID)
	resp, err := loginService.Auth(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

func (gate *Gate) handleLogoutRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	clientID := conn.Context().(int64)
	req := rawMsg.Msg.(*login.LogoutRequest)
	req.Identity = gate.identify(clientID)

	resp, err := loginService.Logout(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

func (gate *Gate) handleC2CSendRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	clientID := conn.Context().(int64)
	req := rawMsg.Msg.(*chat.C2CSendRequest)
	req.Identity = gate.identify(clientID)

	resp, err := c2cService.C2CSend(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return protobuf.Marshal(gate.sm.NextSeq(), resp)
}

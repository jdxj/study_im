package main

import (
	"context"

	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
	"github.com/jdxj/study_im/proto/chat"
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
	case protobuf.LogoutRequest:
		data, err = gate.handleLogoutRequest(conn, rawMsg)
	case protobuf.C2CMsgR:
		data, err = gate.handleC2CMsg(conn, rawMsg)
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
	return protobuf.Marshal(gate.nextSeq(), rawMsg.Seq, &head.Heartbeat{})
}

func (gate *Gate) handleAuthRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	connID := conn.Context().(int64)
	req := rawMsg.Msg.(*login.AuthRequest)
	req.Identity = &pbGate.Identity{
		NodeId:    gate.nodeID,
		ConnId:    connID,
		GateSeq:   gate.nextSeq(),
		ClientSeq: rawMsg.Seq,
	}

	resp, err := loginService.Auth(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.Code == login.Status_AuthSuccessful ||
		resp.Code == login.Status_KickAuthed {
		gate.cm.AddClient(req.UserID, &Client{
			connID: connID,
			userID: req.UserID,
			conn:   conn,
		})
	}
	return protobuf.Marshal(gate.nextSeq(), rawMsg.Seq, resp)
}

func (gate *Gate) handleLogoutRequest(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	connID := conn.Context().(int64)
	req := rawMsg.Msg.(*login.LogoutRequest)
	req.Identity = &pbGate.Identity{
		NodeId:    gate.nodeID,
		ConnId:    connID,
		GateSeq:   gate.nextSeq(),
		ClientSeq: rawMsg.Seq,
	}

	resp, err := loginService.Logout(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.Code == login.Status_LogoutSuccess {
		gate.cm.DelClient(req.UserID)
	}
	return protobuf.Marshal(gate.nextSeq(), rawMsg.Seq, resp)
}

func (gate *Gate) handleC2CMsg(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	connID := conn.Context().(int64)
	req := rawMsg.Msg.(*chat.C2CMsgR)
	req.Identity = &pbGate.Identity{
		NodeId:    gate.nodeID,
		ConnId:    connID,
		GateSeq:   gate.nextSeq(),
		ClientSeq: rawMsg.Seq,
	}

	resp, err := c2cService.C2CMsg(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return protobuf.Marshal(gate.nextSeq(), rawMsg.Seq, resp)
}

func (gate *Gate) handleC2CAck(conn gnet.Conn, rawMsg *protobuf.RawMsg) ([]byte, error) {
	connID := conn.Context().(int64)
	req := rawMsg.Msg.(*chat.C2CAckR)
	req.Identity = &pbGate.Identity{
		NodeId:    gate.nodeID,
		ConnId:    connID,
		GateSeq:   gate.nextSeq(),
		ClientSeq: rawMsg.Seq,
	}

	_, err := c2cService.C2CAck(context.Background(), req)
	return nil, err
}

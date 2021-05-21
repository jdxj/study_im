package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jdxj/study_im/dao/mysql"
	"github.com/jdxj/study_im/dao/redis"
	"github.com/jdxj/study_im/logger"
	"github.com/jdxj/study_im/proto/chat"
)

type C2CService struct {
}

func (c2c *C2CService) C2CMsg(ctx context.Context, req *chat.C2CMsgR, reply *chat.C2CMsgA) error {
	reply.Code = chat.Status_MessageStored

	sessionFrom := redis.Session{UserID: req.From}
	err := sessionFrom.Get()
	if err != nil {
		logger.Errorf("Get: %s", err)
		reply.Code = chat.Status_InternalError
		return nil
	}
	if sessionFrom.NodeID == 0 {
		reply.Code = chat.Status_NotLoggedIn
		return nil
	}
	if sessionFrom.NodeID != req.Identity.NodeId ||
		sessionFrom.ConnID != req.Identity.ConnId {
		reply.Code = chat.Status_IllegalID
		return nil
	}

	msgID := req.MsgId
	if req.MsgId == 0 { // 不是重发
		// todo: 是否验证 to 的存在
		content, _ := json.Marshal(req.Msg)
		ms := &mysql.MessageSend{
			FromID:   req.From,
			ToID:     req.To,
			Seq:      req.Identity.Seq,
			Content:  content,
			SendTime: time.Now(),
			SendType: 1,
		}
		err := ms.Insert()
		if err != nil {
			logger.Errorf("ms.Insert: %s", err)
			reply.Code = chat.Status_InternalError
			return nil
		}

		mr := &mysql.MessageReceive{
			FromID:    req.From,
			ToID:      req.To,
			MessageID: ms.ID,
		}
		err = mr.Insert()
		if err != nil {
			logger.Errorf("Insert: %s", err)
			reply.Code = chat.Status_InternalError
			return nil
		}
		msgID = ms.ID
	}

	session := redis.Session{UserID: req.To}
	err = session.Get()
	if err != nil || session.NodeID == 0 { // 对方不在线
		if err != nil {
			logger.Errorf("session.Get: %s", err)
		}

		// 发送伪 ackN
		ackN := &chat.C2CAckN{
			From:  req.From,
			To:    req.To,
			MsgId: msgID,
		}
		identity := req.Identity
		err = Publish(identity.NodeId, identity.Seq, req.From, ackN)
		if err != nil {
			logger.Errorf("Publish: %s", err)
		}
		return nil
	}

	msgN := &chat.C2CMsgN{
		From:  req.From,
		Msg:   req.Msg,
		MsgId: msgID,
	}
	err = Publish(session.NodeID, req.Identity.Seq, session.UserID, msgN)
	if err != nil {
		logger.Errorf("Publish: %s", err)
	}
	return nil
}

func (c2c *C2CService) C2CAck(ctx context.Context, req *chat.C2CAckR, reply *chat.C2CAckA) error {
	reply.Code = chat.Status_MsgConfirmed

	sessionTo := redis.Session{UserID: req.To}
	err := sessionTo.Get()
	if err != nil {
		logger.Errorf("Get: %s", err)
		reply.Code = chat.Status_InternalError
		return nil
	}
	if sessionTo.NodeID == 0 {
		reply.Code = chat.Status_NotLoggedIn
		return nil
	}
	if sessionTo.NodeID != req.Identity.NodeId ||
		sessionTo.ConnID != req.Identity.ConnId {
		reply.Code = chat.Status_IllegalID
		return nil
	}

	mr := &mysql.MessageReceive{
		ToID:      req.To,
		MessageID: req.MsgId,
	}
	err = mr.SetRead()
	if err != nil {
		logger.Errorf("mr.SetRead(): %s", err)
		reply.Code = chat.Status_InternalError
		return nil
	}
	reply.MsgId = req.MsgId

	session := &redis.Session{UserID: req.From}
	err = session.Get()
	if err != nil || session.NodeID == 0 { // 发送方不在线就无所谓了
		if err != nil {
			logger.Errorf("session.Get: %s", err)
		}
		return nil
	}

	ackN := &chat.C2CAckN{
		From:  req.From,
		To:    req.To,
		MsgId: req.MsgId,
	}
	err = Publish(session.NodeID, req.Identity.Seq, req.From, ackN)
	if err != nil {
		logger.Errorf("Publish: %s", err)
	}
	return nil
}

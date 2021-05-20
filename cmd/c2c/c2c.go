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
	msgID := req.MsgId
	if req.MsgId != 0 {
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
			//reply.MsgId == 0 表示出错, 发送失败
			logger.Errorf("ms.Insert: %s", err)
			return nil
		}
		msgID = ms.ID
	}

	msgN := &chat.C2CMsgN{
		From:  req.From,
		Msg:   req.Msg,
		MsgId: msgID,
	}

	session := redis.Session{UserID: req.To}
	err := session.Get()
	if err != nil || session.NodeID == 0 {
		if err != nil {
			logger.Errorf("session.Get: %s", err)
		}

		// 发送伪 ackN
		ackN := &chat.C2CAckN{MsgId: msgID}
		identity := req.Identity
		err = Publish(identity.NodeId, identity.Seq, identity.ClientId, ackN)
		if err != nil {
			logger.Errorf("Publish: %s", err)
		}
		return nil
	}

	err = Publish(session.NodeID, req.Identity.Seq, session.ClientID, msgN)
	if err != nil {
		logger.Errorf("Publish: %s", err)
	}
	return nil
}

func (c2c *C2CService) C2CAck(ctx context.Context, req *chat.C2CAckR, reply *chat.C2CAckA) error {
	mr := &mysql.MessageReceive{
		ToID:      req.To,
		MessageID: req.MsgId,
	}
	err := mr.SetRead()
	if err != nil {
		logger.Errorf("mr.SetRead(): %s", err)
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
		MsgId: req.MsgId,
	}
	err = Publish(session.NodeID, req.Identity.Seq, req.Identity.ClientId, ackN)
	if err != nil {
		logger.Errorf("Publish: %s", err)
	}
	return nil
}

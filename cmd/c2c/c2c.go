package main

import (
	"context"

	"github.com/jdxj/study_im/dao/redis"
	"github.com/jdxj/study_im/logger"

	"github.com/jdxj/study_im/proto/chat"
)

type C2CService struct {
	msgID uint64
}

func (c2c *C2CService) C2CSend(ctx context.Context, req *chat.C2CSendRequest, reply *chat.C2CSendResponse) error {
	// todo: 消息存储

	c2c.msgID++
	reply.MsgId = c2c.msgID

	pushReq := &chat.C2CPushRequest{
		From:  req.From,
		Msg:   req.Msg,
		MsgId: c2c.msgID,
	}

	session := redis.Session{UserID: req.To}
	err := session.Get()
	if err != nil {
		logger.Errorf("Get: %s", err)
		return nil
	}
	if session.NodeID == 0 { // 不在线
		return nil
	}

	err = Publish(session.NodeID, session.ClientID, pushReq)
	if err != nil {
		logger.Errorf("Publish: %s", err)
	}
	return nil
}

// todo: 实现
func (c2c *C2CService) C2CPush(ctx context.Context, req *chat.C2CPushResponse, ops *chat.Options) error {
	return nil
}

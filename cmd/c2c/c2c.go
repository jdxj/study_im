package main

import (
	"context"

	"github.com/jdxj/study_im/proto/chat"
)

type C2CService struct {
	msgID uint64
}

func (c2c *C2CService) C2CSend(ctx context.Context, req *chat.C2CSendRequest, reply *chat.C2CSendResponse) error {
	c2c.msgID++
	reply.MsgId = c2c.msgID
	return nil
}

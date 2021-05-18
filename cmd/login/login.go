package main

import (
	"context"

	"github.com/jdxj/study_im/dao/redis"
	"github.com/jdxj/study_im/logger"

	"github.com/jdxj/study_im/proto/login"
)

type LoginService struct {
}

func (ls *LoginService) Auth(ctx context.Context, req *login.AuthRequest, reply *login.AuthResponse) error {
	reply.Status = 1
	session := redis.Session{UserID: req.Uid}
	err := session.Get()
	if err != nil {
		logger.Errorf("Get: %s", err)
		reply.Status = 2
		return nil
	}

	if session.NodeID != 0 {
		reply.Status = 3
		return nil
	}

	session.NodeID = req.Identity.NodeId
	session.ClientID = req.Identity.ClientId
	err = session.Set()
	if err != nil {
		logger.Errorf("Set: %s", err)
		reply.Status = 4
		return nil
	}
	return nil
}

func (ls *LoginService) Logout(ctx context.Context, req *login.LogoutRequest, reply *login.LogoutResponse) error {
	return nil
}

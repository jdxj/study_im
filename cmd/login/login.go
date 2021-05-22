package main

import (
	"context"
	"strconv"

	"github.com/jdxj/study_im/dao/redis"
	"github.com/jdxj/study_im/logger"
	"github.com/jdxj/study_im/proto/login"
)

type LoginService struct {
}

func (ls *LoginService) Auth(ctx context.Context, req *login.AuthRequest, reply *login.AuthResponse) error {
	reply.Code = login.Status_AuthSuccessful

	// todo: 验证 req.Token 中的 UserID 与 req.UserID 的一致性
	if req.Token != strconv.Itoa(int(req.UserID)) {
		reply.Code = login.Status_InvalidToken
		return nil
	}

	// 合法的 token
	session := redis.Session{UserID: req.UserID}
	err := session.Get()
	if err != nil {
		logger.Errorf("session.Get: %s", err)
		reply.Code = login.Status_InternalError
		return nil
	}

	// 从未登录
	if session.NodeID == 0 {
		session.NodeID = req.Identity.NodeId
		session.ConnID = req.Identity.ConnId
		err = session.Set()
		if err != nil {
			logger.Errorf("session.Set: %s", err)
			reply.Code = login.Status_InternalError
		}
		return nil
	}

	// 同一 node, 同一 conn 重复发送登录
	if session.NodeID == req.Identity.NodeId &&
		session.ConnID == req.Identity.ConnId {
		reply.Code = login.Status_RepeatAuth
		return nil
	}

	// 其他情况需要踢人
	reply.Code = login.Status_KickAuthed

	kick := &login.KickOutRequest{
		Reason: login.Reason_OtherLogin,
	}
	err = PublishKickOut(session.NodeID, req.UserID, req.Identity.GateSeq,
		req.Identity.ClientSeq, session.ConnID, kick)
	if err != nil {
		logger.Errorf("PublishKickOut: %s", err)
	}

	// 更新 redis
	session.NodeID = req.Identity.NodeId
	session.ConnID = req.Identity.ConnId
	err = session.Set()
	if err != nil {
		reply.Code = login.Status_InternalError
	}
	return nil
}

func (ls *LoginService) Logout(ctx context.Context, req *login.LogoutRequest, reply *login.LogoutResponse) error {
	reply.Code = login.Status_LogoutSuccess

	// todo: 验证 req.Token 中的 UserID 与 req.UserID 的一致性
	if req.Token != strconv.Itoa(int(req.UserID)) {
		reply.Code = login.Status_InvalidToken
		return nil
	}

	session := &redis.Session{UserID: req.UserID}
	err := session.Del()
	if err != nil {
		logger.Errorf("session.Del: %s", err)
		reply.Code = login.Status_InternalError
	}
	return nil
}

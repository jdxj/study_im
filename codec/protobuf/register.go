package protobuf

import (
	"github.com/jdxj/study_im/proto/chat"
	"github.com/jdxj/study_im/proto/head"
	"github.com/jdxj/study_im/proto/login"
)

var (
	p *Processor
)

func init() {
	p = NewProcessor()

	registerMsg(p)
}

func Marshal(msg interface{}) ([]byte, error) {
	return p.Marshal(msg)
}

func Unmarshal(data []byte) (uint16, interface{}, error) {
	return p.Unmarshal(data)
}

const (
	Heartbeat = iota

	AuthRequest
	AuthResponse
	LogoutRequest
	LogoutResponse
	KickOutRequest
	KickOutResponse

	C2CSendRequest
	C2CSendResponse
	C2CPushRequest
	C2CPushResponse
	C2GSendRequest
	C2GSendResponse
	C2GPushRequest
	C2GPushResponse
	C2SPullMsgRequest
	C2SPullMsgResponse
)

func registerMsg(p *Processor) {
	p.Register(Heartbeat, &head.Heartbeat{})

	p.Register(AuthRequest, &login.AuthRequest{})
	p.Register(AuthResponse, &login.AuthResponse{})
	p.Register(LogoutRequest, &login.LogoutRequest{})
	p.Register(LogoutResponse, &login.LogoutResponse{})
	p.Register(KickOutRequest, &login.KickOutRequest{})
	p.Register(KickOutResponse, &login.KickOutResponse{})

	p.Register(C2CSendRequest, &chat.C2CSendRequest{})
	p.Register(C2CSendResponse, &chat.C2CSendResponse{})
	p.Register(C2CPushRequest, &chat.C2CPushRequest{})
	p.Register(C2CPushResponse, &chat.C2CPushResponse{})
	p.Register(C2GSendRequest, &chat.C2GSendRequest{})
	p.Register(C2GSendResponse, &chat.C2GSendResponse{})
	p.Register(C2GPushRequest, &chat.C2GPushRequest{})
	p.Register(C2GPushResponse, &chat.C2GPushResponse{})
	p.Register(C2SPullMsgRequest, &chat.C2SPullMsgRequest{})
	p.Register(C2SPullMsgResponse, &chat.C2SPullMsgResponse{})
}

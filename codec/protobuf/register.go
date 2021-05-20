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

func Marshal(seq uint32, msg interface{}) ([]byte, error) {
	return p.Marshal(seq, msg)
}

func Unmarshal(data []byte) (*RawMsg, error) {
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

	C2CMsgR
	C2CMsgA
	C2CMsgN
	C2CAckR
	C2CAckA
	C2CAckN

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

	p.Register(C2CMsgR, &chat.C2CMsgR{})
	p.Register(C2CMsgA, &chat.C2CMsgA{})
	p.Register(C2CMsgN, &chat.C2CMsgN{})
	p.Register(C2CAckR, &chat.C2CAckR{})
	p.Register(C2CAckA, &chat.C2CAckA{})
	p.Register(C2CAckN, &chat.C2CAckN{})

	p.Register(C2GSendRequest, &chat.C2GSendRequest{})
	p.Register(C2GSendResponse, &chat.C2GSendResponse{})
	p.Register(C2GPushRequest, &chat.C2GPushRequest{})
	p.Register(C2GPushResponse, &chat.C2GPushResponse{})
	p.Register(C2SPullMsgRequest, &chat.C2SPullMsgRequest{})
	p.Register(C2SPullMsgResponse, &chat.C2SPullMsgResponse{})
}

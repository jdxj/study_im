package main

import (
	"github.com/jdxj/study_im/proto/chat"

	"github.com/jdxj/study_im/proto/login"
)

func (cli *Cli) handleAuthResponse(resp *login.AuthResponse) {
}

func (cli *Cli) handleLogoutResponse(resp *login.LogoutResponse) {
}

func (cli *Cli) handleKickOutRequest(req *login.KickOutRequest) {
}

func (cli *Cli) handleC2CMsgA(resp *chat.C2CMsgA) {
	lastMsgID[resp.To] = resp.MsgId
}

func (cli *Cli) handleC2CMsgN(resp *chat.C2CMsgN) {
	lastMsgID[resp.From] = resp.MsgId
}

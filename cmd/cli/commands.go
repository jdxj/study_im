package main

import (
	"flag"
	"fmt"

	"github.com/jdxj/study_im/proto/chat"

	"github.com/jdxj/study_im/proto/login"
)

const (
	List   = "list"
	Auth   = "auth"
	Logout = "logout"
	Send   = "send"
)

var (
	cmdList = []string{
		List,
		Auth,
		Logout,
		Send,
	}

	commands = make(map[string]Parser)

	// status
	token     string
	userID    uint32
	lastMsgID = make(map[uint32]int64)
)

func init() {
	commands[List] = NewListCmd()
	commands[Auth] = NewAuthCmd()
	commands[Logout] = NewLogoutCmd()
	commands[Send] = NewSendCmd()
}

type Parser interface {
	Parse([]string) (interface{}, error)
}

func NewAuthCmd() *AuthCmd {
	ac := &AuthCmd{
		fs: flag.NewFlagSet(Auth, flag.ContinueOnError),
	}
	ac.token = ac.fs.String("token", "test token", "auth token")
	ac.uid = ac.fs.Uint("uid", 123, "user id")

	return ac
}

type AuthCmd struct {
	fs *flag.FlagSet

	token *string
	uid   *uint
}

func (ac *AuthCmd) Parse(args []string) (interface{}, error) {
	err := ac.fs.Parse(args)
	if err != nil {
		return nil, err
	}

	req := &login.AuthRequest{
		Token:  *ac.token,
		UserID: uint32(*ac.uid),
	}
	token = *ac.token
	userID = uint32(*ac.uid)
	return req, nil
}

func NewLogoutCmd() *LogoutCmd {
	lc := &LogoutCmd{fs: flag.NewFlagSet(Logout, flag.ContinueOnError)}
	lc.token = lc.fs.String("token", "test token", "auth token")
	lc.uid = lc.fs.Uint("uid", 123, "user id")
	return lc
}

type LogoutCmd struct {
	fs *flag.FlagSet

	token *string
	uid   *uint
}

func (lc *LogoutCmd) Parse(args []string) (interface{}, error) {
	err := lc.fs.Parse(args)
	if err != nil {
		return nil, err
	}

	req := &login.LogoutRequest{
		Token:  token,
		UserID: userID,
	}
	return req, nil
}

func NewListCmd() *ListCmd {
	return &ListCmd{}
}

type ListCmd struct {
}

func (lc *ListCmd) Parse(args []string) (interface{}, error) {
	for _, cmd := range cmdList {
		fmt.Printf("- %s\n", cmd)
	}
	return nil, nil
}

func NewSendCmd() *SendCmd {
	sc := &SendCmd{
		fs: flag.NewFlagSet(Send, flag.ContinueOnError),
	}
	sc.from = sc.fs.Uint("from", 123, "from uid")
	sc.to = sc.fs.Uint("to", 456, "to uid")
	sc.msg = sc.fs.String("msg", "hello", "message")
	return sc
}

// SendCmd ???????????????????????????
type SendCmd struct {
	fs *flag.FlagSet

	from *uint
	to   *uint
	msg  *string
}

func (sc *SendCmd) Parse(args []string) (interface{}, error) {
	err := sc.fs.Parse(args)
	if err != nil {
		return nil, err
	}

	req := &chat.C2CMsgR{
		From: uint32(*sc.from),
		To:   uint32(*sc.to),
		Msg:  &chat.Message{Text: *sc.msg},
	}
	return req, nil
}

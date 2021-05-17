package main

import (
	"flag"
	"fmt"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/proto/login"
)

const (
	List = "list"
	Auth = "auth"
)

var (
	cmdList = []string{
		List,
		Auth,
	}

	commands = make(map[string]Parser)
)

func init() {
	commands[List] = NewListCmd()
	commands[Auth] = NewAuthCmd()
}

type Parser interface {
	Parse([]string) ([]byte, error)
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

func (ac *AuthCmd) Parse(args []string) ([]byte, error) {
	err := ac.fs.Parse(args)
	if err != nil {
		return nil, err
	}

	req := &login.AuthRequest{
		Token: *ac.token,
		Uid:   uint32(*ac.uid),
	}
	return protobuf.Marshal(req)
}

func NewListCmd() *ListCmd {
	return &ListCmd{}
}

type ListCmd struct {
}

func (lc *ListCmd) Parse(args []string) ([]byte, error) {
	for _, cmd := range cmdList {
		fmt.Println(cmd)
	}
	return nil, nil
}

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jdxj/study_im/proto/head"
)

type Parser interface {
	Init(*CmdLine)
	Name() string
	Parse([]string) (interface{}, error)
}

type CmdLine struct {
	fsMap  map[string]Parser
	client *Client
}

func (cli *CmdLine) Init() {
	cli.fsMap = make(map[string]Parser)

	RegisterParser(cli, &listFlagSet{})
	RegisterParser(cli, &headFlagSet{})
}

func (cli *CmdLine) Parse(args []string) (interface{}, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("args not enough")
	}

	parser, ok := cli.fsMap[args[0]]
	if !ok {
		return nil, fmt.Errorf("cmd not register: %s, %v", args[0], []byte(args[0]))
	}

	flags := args[1:]
	if len(flags) == 0 {
		flags = []string{"-h"}
	}

	return parser.Parse(flags)
}

const (
	listCmd = "list"
	headCmd = "head"
)

func RegisterParser(cli *CmdLine, parser Parser) {
	parser.Init(cli)
	fsMap := cli.fsMap
	fsMap[parser.Name()] = parser
}

type listFlagSet struct {
	fs  *flag.FlagSet
	cli *CmdLine

	// param
	name *string
}

func (list *listFlagSet) Name() string {
	return listCmd
}

func (list *listFlagSet) Init(cli *CmdLine) {
	list.cli = cli
	fs := flag.NewFlagSet(listCmd, flag.ContinueOnError)
	list.fs = fs

	//fs.Usage =

	list.name = fs.String("name", "abc", "user name")
}

func (list *listFlagSet) Parse(args []string) (interface{}, error) {
	err := list.fs.Parse(args)
	return nil, err
}

type headFlagSet struct {
	fs  *flag.FlagSet
	cli *CmdLine

	// param
	version   *int64
	seq       *int64
	timestamp *int64
}

func (hfs *headFlagSet) Name() string {
	return headCmd
}

func (hfs *headFlagSet) Init(cli *CmdLine) {
	hfs.cli = cli
	fs := flag.NewFlagSet(headCmd, flag.ContinueOnError)
	hfs.fs = fs

	hfs.version = fs.Int64("version", 1, "head.version")
	hfs.seq = fs.Int64("seq", 0, "head.seq")
	hfs.timestamp = fs.Int64("timestamp", time.Now().Unix(), "head.ts")
}

func (hfs *headFlagSet) Parse(args []string) (interface{}, error) {
	err := hfs.fs.Parse(args)
	if err != nil {
		return nil, err
	}

	msg := &head.Head{
		Version:   uint32(*hfs.version),
		Seq:       hfs.cli.client.seq,
		Timestamp: *hfs.timestamp,
	}
	if *hfs.seq != 0 {
		msg.Seq = uint32(*hfs.seq)
	}
	return msg, nil
}

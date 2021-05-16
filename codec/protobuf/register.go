package protobuf

import "github.com/jdxj/study_im/proto/head"

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
	HEAD = iota
)

func registerMsg(p *Processor) {
	p.Register(HEAD, &head.Head{})
}

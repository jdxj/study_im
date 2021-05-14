package proto

import (
	"github.com/golang/protobuf/proto"
	"github.com/jdxj/study_im/proto/head"
)

func SortedMsg() []proto.Message {
	messages := []proto.Message{
		&head.Head{},
	}
	return messages
}

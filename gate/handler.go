package gate

import (
	"time"

	"github.com/jdxj/study_im/proto/head"
)

func handle(a *agent, data interface{}) {
	switch msg := data.(type) {
	case *head.Head:
		handleHead(a, msg)
	}
}

func handleHead(a *agent, msgHead *head.Head) {
	msgHead.Seq += 1
	msgHead.Timestamp = time.Now().Unix()

	a.WriteMsg(msgHead)
}

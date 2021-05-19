package main

import (
	"fmt"
	"math"

	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/dao/rabbit"
)

var (
	broker *rabbit.Broker
)

func InitBroker(user, pass, host, bindingKey string, port int) error {
	broker = rabbit.New(user, pass, host, bindingKey, port)
	return broker.Connect()
}

func Publish(nodeID uint32, clientID int64, msg interface{}) error {
	headers := make(map[string]interface{})
	headers["nodeID"] = int64(nodeID) // rabbitmq driver 提示不支持 uint32
	headers["logicID"] = clientID
	headers["type"] = "c2c"

	// 这里可以用其他编码方式, seq 随便填写.
	// seq 应该由 gate 负责.
	data, err := protobuf.Marshal(math.MaxUint32, msg)
	if err != nil {
		return err
	}

	routingKey := fmt.Sprintf("node.%d", nodeID)
	return broker.Publish(routingKey, headers, data)
}

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

func PublishKickOut(nodeID, userID uint32, connID int64, msg interface{}) error {
	headers := make(map[string]interface{})
	headers["nodeID"] = int64(nodeID) // rabbitmq driver 提示不支持 uint32
	headers["userID"] = int64(userID)
	headers["connID"] = connID
	headers["action"] = "kick"

	data, err := protobuf.Marshal(math.MaxUint32, 0, msg)
	if err != nil {
		return err
	}

	routingKey := fmt.Sprintf("node.%d", nodeID)
	return broker.Publish(routingKey, headers, data)
}

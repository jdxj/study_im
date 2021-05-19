package main

import (
	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/dao/rabbit"
	"github.com/jdxj/study_im/logger"
)

var (
	broker *rabbit.Broker
)

func InitBroker(user, pass, host, bindingKey string, port int) error {
	broker = rabbit.New(user, pass, host, bindingKey, port)
	return broker.Connect()
}

func (gate *Gate) handleBroker(headers map[string]interface{}, body []byte) error {
	if headers == nil {
		logger.Warnf("handleBroker: empty headers")
		return nil
	}

	nodeID, ok := headers["nodeID"].(int64)
	if ok && nodeID != int64(gate.nodeID) {
		logger.Warnf("handleBroker: mismatch node: %d", nodeID)
		return nil
	}

	typ, ok := headers["type"].(string)
	if !ok {
		logger.Warnf("handleBroker: invalid typ: %v", headers["type"])
	}

	logicID, ok := headers["logicID"].(int64)
	if !ok {
		logger.Warnf("handleBroker: logicID not found: %d", headers["logicID"])
		return nil
	}

	rawMsg, err := protobuf.Unmarshal(body)
	if err != nil {
		logger.Errorf("Unmarshal: %s", err)
		return nil
	}

	// todo: 实现单发/群发
	switch typ {
	case "c2c":
		body, err = protobuf.Marshal(gate.sm.NextSeq(), rawMsg.Msg)
		if err != nil {
			// 可能是数据错误, 让 rabbit 重发
			return err
		}
		conn := gate.am.GetAgent(logicID)
		err := conn.AsyncWrite(body)
		if err != nil {
			logger.Errorf("AsyncWrite: %s", err)
		}
	}
	return nil
}

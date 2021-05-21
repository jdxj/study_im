package main

import (
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
		logger.Warnf("handleBroker: mismatch node: %v", headers["nodeID"])
		return nil
	}

	action, ok := headers["action"].(string)
	if !ok {
		logger.Errorf("handleBroker: invalid action: %v", headers["action"])
		return nil
	}

	switch action {
	case "kick":
		userID, ok := headers["userID"].(int64)
		if !ok {
			logger.Errorf("invalid userID: %v", headers["userID"])
			return nil
		}
		connID, ok := headers["connID"].(int64)
		if !ok {
			logger.Errorf("invalid connID: %v", headers["connID"])
			return nil
		}

		gate.cm.DelClient(uint32(userID))
		conn := gate.rm.GetConn(connID)
		if conn != nil {
			err := conn.AsyncWrite(body)
			if err != nil {
				logger.Errorf("AsyncWrite: %s", err)
			}
		}

	case "c2c":
		userID, ok := headers["userID"].(int64)
		if !ok {
			logger.Errorf("invalid userID: %v", headers["userID"])
			return nil
		}
		client := gate.cm.GetClient(uint32(userID))
		if client != nil {
			err := client.conn.AsyncWrite(body)
			if err != nil {
				logger.Errorf("AsyncWrite: %s", err)
			}
		}
	default:
		logger.Warnf("not define action: %s", action)
	}
	return nil
}

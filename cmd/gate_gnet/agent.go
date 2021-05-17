package main

import (
	"sync"

	"github.com/panjf2000/gnet"
)

type Agent struct {
	nodeID  int
	agentID int

	conn gnet.Conn

	// 可能的其他附加信息
}

type AgentManager struct {
	mutex  *sync.RWMutex
	agents map[int64]gnet.Conn
}

func (am *AgentManager) GetAgent(agentID int64) gnet.Conn {
	am.mutex.RLock()
	conn := am.agents[agentID]
	am.mutex.RUnlock()
	return conn
}

func (am *AgentManager) AddAgent(agentID int64, conn gnet.Conn) {
	am.mutex.Lock()
	am.agents[agentID] = conn
	am.mutex.Unlock()

	conn.SetContext(agentID)
}

func (am *AgentManager) DelAgent(agentID int64) {
	am.mutex.Lock()
	delete(am.agents, agentID)
	am.mutex.Unlock()
}

func (am *AgentManager) Range(f func(int64, gnet.Conn)) {
	am.mutex.RLock()
	for agentID, conn := range am.agents {
		f(agentID, conn)
	}
	am.mutex.RUnlock()
}

package main

import (
	"sync"

	"github.com/panjf2000/gnet"
)

type Client struct {
	connID int64
	userID uint32
	conn   gnet.Conn
}

// ClientManager 已登录的连接, todo: 需要配合心跳机制
type ClientManager struct {
	mutex   sync.RWMutex
	clients map[uint32]*Client
}

func (cm *ClientManager) GetClient(userID uint32) *Client {
	cm.mutex.RLock()
	client := cm.clients[userID]
	cm.mutex.RUnlock()
	return client
}

func (cm *ClientManager) AddClient(userID uint32, client *Client) {
	cm.mutex.Lock()
	cm.clients[userID] = client
	cm.mutex.Unlock()
}

func (cm *ClientManager) DelClient(userID uint32) {
	cm.mutex.Lock()
	delete(cm.clients, userID)
	cm.mutex.Unlock()
}

func (cm *ClientManager) Range(f func(uint32, *Client)) {
	cm.mutex.RLock()
	for userID, client := range cm.clients {
		f(userID, client)
	}
	cm.mutex.RUnlock()
}

type Group struct {
	groupID uint32
	cm      *ClientManager
}

type GroupManager struct {
	mutex  sync.RWMutex
	groups map[uint32]*Group
}

func (gm *GroupManager) AddMember(groupID, userID uint32, client *Client) {
	gm.mutex.Lock()
	group, ok := gm.groups[groupID]
	if !ok {
		group = &Group{
			groupID: groupID,
			cm: &ClientManager{
				clients: make(map[uint32]*Client),
			},
		}
		gm.groups[groupID] = group
	}
	gm.mutex.Unlock()

	group.cm.AddClient(userID, client)
}

func (gm *GroupManager) DelMember(groupID, userID uint32) {
	gm.mutex.RLock()
	group := gm.groups[groupID]
	gm.mutex.RUnlock()

	if group != nil {
		group.cm.DelClient(userID)
	}
}

func (gm *GroupManager) Range(groupID uint32, f func(uint32, *Client)) {
	gm.mutex.RLock()
	group := gm.groups[groupID]
	gm.mutex.RUnlock()

	if group != nil {
		group.cm.Range(f)
	}
}

// RelationManager 已与 Gate 建立 tcp 的连接,
// Gate 不应该主动关闭连接
type RelationManager struct {
	mutex       sync.RWMutex
	connections map[int64]gnet.Conn
}

func (rm *RelationManager) GetConn(connID int64) gnet.Conn {
	rm.mutex.RLock()
	conn := rm.connections[connID]
	rm.mutex.RUnlock()
	return conn
}

func (rm *RelationManager) AddConn(connID int64, conn gnet.Conn) {
	rm.mutex.Lock()
	_, ok := rm.connections[connID]
	if !ok {
		rm.connections[connID] = conn
	}
	rm.mutex.Unlock()
}

func (rm *RelationManager) DelConn(connID int64) {
	rm.mutex.Lock()
	delete(rm.connections, connID)
	rm.mutex.Unlock()
}

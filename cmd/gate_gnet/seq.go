package main

import "sync"

type SeqManager struct {
	mutex sync.Mutex
	next  uint32
}

func (sm *SeqManager) NextSeq() uint32 {
	sm.mutex.Lock()
	seq := sm.next
	sm.next++
	sm.mutex.Unlock()
	return seq
}

package util

import (
	"container/list"
	"sync"
)

// SlaveList is the list of slaves
type SlaveList struct {
	list *list.List
	mu   sync.RWMutex
}

// SlaveListElement Each slave contains an address to connect to it and a name which we give it to it
type SlaveListElement struct {
	Address string
	Id      uint32
	Dead    bool
}

// NewSlaveList will create a slave list
func NewSlaveList() *SlaveList {
	return &SlaveList{
		list: list.New(),
	}
}

func (l *SlaveList) AddSlave(slave SlaveListElement) {
	l.mu.Lock()
	l.list.PushBack(&slave)
	l.mu.Unlock()
}

func (l *SlaveList) ToList() []SlaveListElement {
	result := make([]SlaveListElement, 0, l.list.Len())
	l.mu.RLock()
	for e := l.list.Front(); e != nil; e = e.Next() {
		result = append(result, *e.Value.(*SlaveListElement))
	}
	l.mu.RUnlock()
	return result
}

// MakeDead will make the dead parameter of
func (l *SlaveList) MakeDead(address string) {
	l.mu.RLock()
	for e := l.list.Front(); e != nil; e = e.Next() {
		v := e.Value.(*SlaveListElement)
		if v.Address == address {
			v.Dead = true
			break
		}
	}
	l.mu.RUnlock()
}

package synctools

import (
	"sync"
)

type RWMutex struct {
	locker sync.RWMutex
}

func (m *RWMutex) LockDo(fn func()) {
	m.locker.Lock()
	defer m.locker.Unlock()

	fn()
}

func (m *RWMutex) RLockDo(fn func()) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	fn()
}

type Mutex struct {
	locker sync.Mutex
}

func (m *Mutex) LockDo(fn func()) {
	m.locker.Lock()
	defer m.locker.Unlock()

	fn()
}

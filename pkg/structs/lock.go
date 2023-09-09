package structs

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	ErrLockLocked error = fmt.Errorf("lock is locked")
)

type Lock struct {
	isRunning int32
	mu        sync.Mutex
}

func (l *Lock) Lock() error {
	free := atomic.CompareAndSwapInt32(&l.isRunning, 0, 1)

	if !free {
		return ErrLockLocked
	}

	l.mu.Lock()

	return nil
}

func (l *Lock) Unlock() {
	l.mu.Unlock()
	atomic.StoreInt32(&l.isRunning, 0)
}

func (l *Lock) IsLocked() bool {
	return atomic.LoadInt32(&l.isRunning) == 1
}

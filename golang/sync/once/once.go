package once

import (
	"sync"
	"sync/atomic"
)

/**
 *@Author tudou
 *@Date 2021/1/8
 **/

type once struct {
	done uint32
	m sync.Mutex
}
func (o *once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}
func (o *once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	// 双检查
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}



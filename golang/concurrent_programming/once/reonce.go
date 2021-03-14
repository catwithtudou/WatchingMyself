package once

import (
	"sync"
	"sync/atomic"
)

/**
 *@Author tudou
 *@Date 2021/1/8
 **/

type ReOnce struct {
	m sync.Mutex
	done uint32
}


func (o *ReOnce) Do(f func() error) error {
	//fast path
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}


func (o *ReOnce) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}



package mutex

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

/**
 *@Author tudou
 *@Date 2021/1/1
 **/

type recursiveMutex struct{
	sync.Mutex
	owner int64 //持有该锁的goroutine id
	recursion int32 //该goroutine重入的次数
}

func (m *recursiveMutex)lock(){
	gid := goid.Get()
	//若持有该锁的goroutine为这次调用的goroutine则为重入
	if atomic.LoadInt64(&m.owner) == gid{
		m.recursion++
		return
	}
	//若第一次获得锁则需要记录其gid
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner,gid)
	m.recursion = 1
}

func (m *recursiveMutex) unlock()  {
	gid := goid.Get()
	//非持有该锁的goroutine尝试释放锁则不合理
	if atomic.LoadInt64(&m.owner) !=gid{
		panic(fmt.Sprintf("it's wrong owner(%d): %d!",m.owner,gid))
	}
	//调用次数-1
	m.recursion--
	//说明持有该锁的goroutine还有
	if m.recursion!=0{
		return
	}
	atomic.StoreInt64(&m.owner,-1)
	m.Mutex.Unlock()
}
package mutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

/**
 *@Author tudou
 *@Date 2021/1/3
 **/

//TryLock

const(
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct{
	sync.Mutex
}

//获取锁当前持有和等待中的goroutine之和
func (m *Mutex)Count()int{
	v:=atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}

//tryLock实现
func (m *Mutex)TryLock()bool{
	//fast path
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)),0, mutexLocked){
		return true
	}

	//获取当前锁state
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	//若处于其中任何一个状态则放弃竞争返回false
	if old&(mutexLocked|mutexWoken|mutexStarving) !=0{
		return false
	}

	//开始竞争
	re := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)),old,re)
}

//锁是否被持有
func (m *Mutex) IsLocked()bool  {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) == mutexLocked
}

//锁是否有等待者被唤醒
func (m *Mutex) IsWoken()bool  {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) == mutexWoken
}

//锁是否处于饥饿状态
func (m *Mutex) IsStarving()bool  {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) == mutexStarving
}


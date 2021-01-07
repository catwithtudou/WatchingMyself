package cond

import "sync"

/**
 *@Author tudou
 *@Date 2021/1/7
 **/


type condQueue struct{
	val []int
	size int
	cap int
	sync.Cond
}

func NewCondQueue(cap int)*condQueue{
	return &condQueue{
		val:  make([]int,0,cap),
		size: 0,
		cap:  cap,
	}
}

func (c *condQueue)Push(val int){
	//判断容量
	c.size++
	if c.size>=c.cap{
		//若大于容量则加入等待队列
		c.L.Lock()
		c.Wait()
		c.L.Unlock()
		//被唤醒时说明容量足够
	}
	c.val=append(c.val,val)
}

func (c *condQueue)Pop()(int,bool){
	if c.size==0{
		return 0,false
	}
	val:=c.val[0]
	c.val=c.val[1:]
	c.Signal()
	return val,true
}
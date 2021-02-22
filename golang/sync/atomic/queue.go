package main

import (
	"sync/atomic"
	"unsafe"
)

type LKQueue struct{
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct{
	value interface{}
	next unsafe.Pointer
}


func NewLKQueue()*LKQueue{
	n:=unsafe.Pointer(&node{})
	return &LKQueue{
		head: n,
		tail: n,
	}
}

func (q *LKQueue)Enqueue(v interface{}){
	n:=&node{value: v}
	for{
		tail:=load(&q.tail)
		next:=load(&tail.next)
		if tail == load(&q.tail){
			if next == nil{ //尾为空即没有数据入队
				if cas(&tail.next,next,n){ //增加到队尾
					cas(&q.tail,tail,n) //入队成功，移动尾巴指针
					return
				}
			}else{ //尾不为空则需要移动尾指针
				cas(&q.tail,tail,next)
			}
		}
	}
}

func (q *LKQueue)Dequeue()interface{}{
	for{
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head){
			if head == tail{
				if next == nil{ //说明为空队列
					return nil
				}
				// 只是尾指针还没调整，尝试调整它指向下一个
				cas(&q.tail,tail,next)
			}else{
				//读取出队的数据
				v:=next.value
				if cas(&q.head,head,next){
					return v
				}
			}
		}
	}


}


func load(p *unsafe.Pointer)(n *node){
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer,old,new *node)(ok bool){
	return atomic.CompareAndSwapPointer(
		p,unsafe.Pointer(old),unsafe.Pointer(new))
}
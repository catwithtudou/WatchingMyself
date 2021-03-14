package mutex

import "sync"

/**
 *@Author tudou
 *@Date 2021/1/3
 **/


type SliceQueue struct {
	mu sync.Mutex
	data []interface{}
}

func NewSliceQueue(n int)*SliceQueue {
	return &SliceQueue{
		data: make([]interface{},0,n),
	}
}

func (q *SliceQueue)Put(v interface{}){
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data=append(q.data,v)
}

func (q *SliceQueue) Out()interface{}  {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data)==0{
		return nil
	}
	v:=q.data[0]
	q.data = q.data[1:]
	return v
}
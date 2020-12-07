package dispatch

/**
 *@Author tudou
 *@Date 2020/12/7
 **/

type LRUCache struct{
	size,capacity int
	cache map[int]*DLinkedNode
	head,tail *DLinkedNode
}

type DLinkedNode struct{
	key,value int
	prev,next *DLinkedNode
}

func initDLinkNode(key,value int)*DLinkedNode{
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int)LRUCache{
	l:=LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     initDLinkNode(0,0),
		tail:     initDLinkNode(0,0),
	}
	l.head.next=l.tail
	l.tail.prev=l.head
	return l
}

func (l *LRUCache)addToHead(node *DLinkedNode){
	node.prev=l.head
	node.next=l.head.next
	l.head.next.prev=node
	l.head.next=node
}

func (l *LRUCache)removeNode(node *DLinkedNode){
	node.prev.next=node.next
	node.next.prev=node.prev
}

func (l *LRUCache)moveToHead(node *DLinkedNode){
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache)removeTail()*DLinkedNode{
	node:=l.tail.prev
	l.removeNode(node)
	return node
}

func (l *LRUCache)Get(key int)int{
	node,ok:=l.cache[key]
	if !ok{
		return -1
	}
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache)Put(key,value int){
	if node,ok:=l.cache[key];!ok{
		node:=initDLinkNode(key,value)
		l.cache[key]=node
		l.size++
		l.addToHead(node)
		if l.size>l.capacity{
			removed:=l.removeTail()
			delete(l.cache,removed.key)
			l.size--
		}
	}else{
		node.value=value
		l.moveToHead(node)
	}
}
package _7_iterator

import "fmt"

type FaIterator interface {
	Iterator()Iterator
}

type Iterator interface {
	First() interface{}
	Next()interface{}
	IsDone()bool
	IteratorPrint()
}

type Numbers struct{
	Start,End int
}

func NewNumbers(start,end int)*Numbers{
	return &Numbers{
		Start: start,
		End:   end,
	}
}

func (n *Numbers)Iterator()Iterator{
	return &NumbersIterator{
		Numbers:    n,
		NextNumber: n.Start,
	}
}

type NumbersIterator struct{
	Numbers *Numbers
	NextNumber int
}

func (n *NumbersIterator)First()interface{}{
	return n.NextNumber
}

func (n *NumbersIterator)IsDone()bool{
	return n.NextNumber-1 == n.Numbers.End
}

func (n *NumbersIterator)Next()interface{}{
	if !n.IsDone(){
		n.NextNumber++
		next:=n.NextNumber
		return next
	}
	return nil
}

func (n *NumbersIterator)IteratorPrint(){
	for start:=n.First();!n.IsDone();{
		fmt.Println(start)
		start=n.Next()
	}
}
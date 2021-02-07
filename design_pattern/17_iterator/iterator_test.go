package _7_iterator

import (
	"testing"
)

func TestIterator(t *testing.T){
	var i FaIterator
	i=NewNumbers(1,10)
	i.Iterator().IteratorPrint()
}

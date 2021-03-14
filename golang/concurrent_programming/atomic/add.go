package main

import (
	"fmt"
	"sync/atomic"
)

func main(){
	var i uint32 = 3
	atomic.AddUint32(&i,^uint32(0))
	fmt.Println(i)
}

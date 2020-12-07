package dispatch

import (
	"fmt"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/7
 **/


func TestLRU(t *testing.T){
	obj:=Constructor(10)
	obj.Put(1,1)
	fmt.Print(obj.Get(1))
}
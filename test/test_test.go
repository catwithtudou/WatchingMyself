package test

import (
	"fmt"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/24
 **/


func TestUintNum(t *testing.T){
	var a uint8 = 1
	var b uint8 = 255
	fmt.Println("result:",a+b)
}

func TestSliceMap(t *testing.T){
	a:=make(map[string][]string)
	b:=make([]string,0)
	b=append(b,"1")
	a["1"]=b
	_,ok:=a["1"]
	if ok{
		a["1"]=append(a["1"],"2")
	}
	fmt.Println(a["1"])
	fmt.Println(b)
}
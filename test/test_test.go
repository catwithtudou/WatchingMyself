package test

import (
	"fmt"
	"testing"
	"time"
)

/**
 *@Author tudou
 *@Date 2020/12/24
 **/


func TestUintNum(t *testing.T){
	tick:=int64(time.Now().Unix()) / 10
	fmt.Printf("%d\n",tick)
	time.Sleep(7*time.Second)
	tick=int64(time.Now().Unix()) / 10
	fmt.Printf("%d\n",tick)
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


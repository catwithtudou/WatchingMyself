package _8_adapter

import "testing"

/**
 *@Author tudou
 *@Date 2020/12/6
 **/


func TestNewAdapter(t *testing.T) {
	origin:=NewOrigin()
	target:=NewAdapter(origin)
	res:=target.Result()
	if res!="no-test"{
		t.Fatal("it's wrong")
	}
}
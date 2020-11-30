package simple_factory

import "testing"

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

func TestNewThing(t *testing.T) {
	head:=NewThing("head")
	head.Work("head")
}
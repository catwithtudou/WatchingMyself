package _4_mediator

import "testing"

func TestMediator(t *testing.T) {
	me:=GetMediatorInstance()
	me.Input=&Input{}
	me.Process=&Process{}
	me.Out=&Out{}

	me.Input.ReadData("byteDance")

}

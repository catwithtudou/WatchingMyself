package _5_Observer

import "testing"

func TestObserver(t *testing.T){
	r1:=NewReceiver("byte")
	r2:=NewReceiver("dance")
	topic:=NewTopic()
	topic.AddReceiver(r1)
	topic.AddReceiver(r2)
	topic.UpdateContext("testContext")
	topic.UpdateContext("doneContext")
}

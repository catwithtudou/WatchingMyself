package _3_bridge

import "testing"

func TestMessageDevice(t *testing.T){
	//ios qq
	iQ:=NewIosDevice(QqMessage())
	iQ.SendMessage("ios qq sending","ios")

	//ios wx
	iW:=NewIosDevice(WxMessage())
	iW.SendMessage("ios wx sending","ios")

	//win qq
	wQ:=NewWinDevice(QqMessage())
	wQ.SendMessage("win qq sending","win")

	//win wx
	wW:=NewWinDevice(WxMessage())
	wW.SendMessage("win wx sending","win")
}

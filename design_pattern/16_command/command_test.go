package _6_command

import "testing"

func TestCommand(t *testing.T){
	ops:=Operation{}
	exe:=NewExe(NewInputCommand(ops),NewOutCommand(ops))
	exe.PressClick()
	exe.PressCan()
}

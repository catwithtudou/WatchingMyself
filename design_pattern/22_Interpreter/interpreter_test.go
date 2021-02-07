package _2_Interpreter

import "testing"

func TestInterpreter(t *testing.T){
	var e expression
	e.parse("1+2+3-1")
	v:=e.result().interpreter()
	if v!=5{
		t.Fatal("wrong")
	}
}

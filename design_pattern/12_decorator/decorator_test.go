package _2_decorator

import (
	"testing"
)

func TestDecorator(t *testing.T){
	var n Number = &ConcreteNumber{}
	n = WrapAddDecorator(n,10)
	n = WrapMulDecorator(n,8)
	result:=n.Calc()
	if result!=80{
		t.Fatal("error")
	}
}
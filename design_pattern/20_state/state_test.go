package _0_state

import "testing"

func TestState(t *testing.T){
	week:=NewDay()
	week.Today()
	week.Next()
	week.Today()
}
package _1_memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T){
	careTaker:=NewCareTaker()
	originator:=NewOriginator("open")
	fmt.Println(originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("closed")
	fmt.Println(originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(careTaker.getMemento(0))
	fmt.Println(originator.GetState())

	originator.RestoreMemento(careTaker.getMemento(1))
	fmt.Println(originator.GetState())
}

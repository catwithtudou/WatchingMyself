package _0_state

import "fmt"

type WeekDay interface {
	Today()
	Next(day *Day)
}

type Day struct{
	today WeekDay
}

func NewDay()*Day{
	return &Day{today:&Sunday{},}
}

func (d *Day)Today(){
	d.today.Today()
}

func (d *Day)Next(){
	d.today.Next(d)
}


type Sunday struct{}

func (*Sunday)Today(){
	fmt.Println("Sunday")
}

func (*Sunday)Next(day *Day){
	day.today = &Monday{}
}

type Monday struct{}

func (*Monday)Today(){
	fmt.Println("Monday")
}

func (*Monday)Next(day *Day){

}



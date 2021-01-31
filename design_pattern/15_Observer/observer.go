package _5_Observer

import "fmt"

type Topic struct{
	observers []Observer
	context string
}

func NewTopic()*Topic{
	return &Topic{
		observers: make([]Observer,0),
	}
}

func (t *Topic)AddReceiver(r *Receiver){
	t.observers=append(t.observers,r)
}

func (t *Topic)Notify(){
	for _,o:=range t.observers{
		o.Update(t)
	}
}

func (t *Topic)UpdateContext(context string){
	t.context = context
	t.Notify()
}

type Observer interface {
	Update(topic *Topic)
}



type Receiver struct{
	Name string
}

func NewReceiver(name string)*Receiver{
	return &Receiver{
		Name:name,
	}
}

func (r *Receiver)Update(s *Topic){
	fmt.Printf("[receiver:%s] receive %s\n",r.Name,s.context)
}

package _4_mediator

import "fmt"

type Input struct{
	Data string
}

func (i *Input)ReadData(data string){
	i.Data = data
	fmt.Printf("input data:%s\n",i.Data)
	GetMediatorInstance().changed(i)
}


type Process struct{
	Data string
}

func (p *Process)ProcessData(data string){
	p.Data = "[process]"+data
	fmt.Printf("Process data:%s\n",p.Data)
	GetMediatorInstance().changed(p)
}

type Out struct{
	Data string
}

func (o *Out)OutData(data string){
	o.Data = "[out]"+data
	fmt.Printf("Out data:%s\n",o.Data)
}

type Mediator struct{
	Input *Input
	Process *Process
	Out *Out
}

var mediator *Mediator

func GetMediatorInstance()*Mediator{
	if mediator==nil{
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator)changed(i interface{}){
	switch inst:=i.(type) {
	case *Input:
		m.Process.ProcessData(inst.Data)
	case *Process:
		m.Out.OutData(inst.Data)
	}
}
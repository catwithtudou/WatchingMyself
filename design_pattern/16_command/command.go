package _6_command

import "fmt"

type Operation struct{}

func (o *Operation)Start(){
	fmt.Println("start command")
}

func (o *Operation)Process(){
	fmt.Println("process command")
}

func (o *Operation)End(){
	fmt.Println("end command")
}

type Command interface {
	Execute()
}

type InputCommand struct{
	Operations Operation
}

func NewInputCommand(op Operation)*InputCommand{
	return &InputCommand{Operations:op}
}

func (i *InputCommand)Execute(){
	i.Operations.Start()
	i.Operations.Process()
}

type OutCommand struct{
	Operations Operation
}

func NewOutCommand(op Operation)*OutCommand{
	return &OutCommand{Operations:op}
}

func (o *OutCommand)Execute(){
	o.Operations.End()
}

type Exe struct{
	Click Command
	Cancel Command
}

func NewExe(cli,can Command)*Exe{
	return &Exe{
		Click:  cli,
		Cancel: can,
	}
}

func (e *Exe)PressClick(){
	e.Click.Execute()
}

func (e *Exe)PressCan(){
	e.Cancel.Execute()
}
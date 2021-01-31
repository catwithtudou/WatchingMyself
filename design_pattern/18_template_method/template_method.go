package _8_template_method

import "fmt"

type Calculate interface {
	Calc()
}

type CalculateImpl interface {
	Input()
	Process()
}

type Template struct{
	CalculateImpl
	Left int
	Right int
}

func NewTemplate(impl CalculateImpl,l,r int)*Template{
	return &Template{
		CalculateImpl: impl,
		Left:          l,
		Right:         r,
	}
}

func (t *Template)Calc(){
	fmt.Println("calculate begin")
	t.Input()
	t.Process()
	fmt.Println("calculate end")
}

type AddCalc struct{
	*Template
}

func NewAddCalc(l,r int)Calculate{
	calc:=&AddCalc{}
	template:=NewTemplate(calc,l,r)
	calc.Template=template
	return calc
}

func (a *AddCalc)Input(){
	fmt.Printf("input data:%d and %d\n",a.Left,a.Right)
}

func (a *AddCalc)Process(){
	fmt.Printf("process data:%d\n",a.Left+a.Right)
}


type MulCalc struct{
	*Template
}

func NewMulCalc(l,r int)Calculate{
	calc:=&MulCalc{}
	template:=NewTemplate(calc,l,r)
	calc.Template=template
	return calc
}

func (a *MulCalc)Input(){
	fmt.Printf("input data:%d and %d\n",a.Left,a.Right)
}

func (a *MulCalc)Process(){
	fmt.Printf("process data:%d\n",a.Left*a.Right)
}


package _8_template_method

import "testing"

func TestTemplateMethod(t *testing.T){
	//add
	var calc Calculate = NewAddCalc(1,1)
	calc.Calc()

	//mul
	var mulc Calculate = NewMulCalc(1,10)
	mulc.Calc()
}

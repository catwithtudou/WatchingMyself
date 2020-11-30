package factory_method

import (
	"fmt"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

func compute(factory OperatorFactory, a, b int) int {
	//抽象工厂创建抽象类
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T){
	//导入加法工厂
	plus:=compute(PlusOperatorFactory{},1,1)
	fmt.Println(plus)
	//导入减法工厂
	minus:=compute(MinusOperatorFactory{},1,1)
	fmt.Println(minus)
}


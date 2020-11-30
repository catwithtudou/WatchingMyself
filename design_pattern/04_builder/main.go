package builder

import "fmt"

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

/**

 */

//生成器
type Builder interface {
	Step1()
	Step2()
	Step3()
}


type Product struct{
	Builder
}

func NewProduct(builder Builder)*Product{
	return &Product{Builder:builder}
}

func (p *Product)WayMake(){
	p.Builder.Step1()
	p.Builder.Step2()
	p.Builder.Step3()
}

type Product1 struct{
	result string
}

func (p *Product1)Step1(){
	p.result+="开始"
}
func (p *Product1)Step2(){
	p.result+="过程"
}
func (p *Product1)Step3(){
	p.result+="结束"
}

func (p *Product1)Result(){
	fmt.Println(p.result)
}

type Product2 struct{
	result string
}

func (p *Product2)Step1(){
	p.result+="结束"
}
func (p *Product2)Step2(){
	p.result+="过程"
}
func (p *Product2)Step3(){
	p.result+="开始"
}

func (p *Product2)Result(){
	fmt.Println(p.result)
}



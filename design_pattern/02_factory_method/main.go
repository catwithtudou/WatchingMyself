package factory_method

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

//使用子类的方式延迟生成对象到子类中实现
//Go中不存在继承
//使用匿名组合来实现

//被封装的抽象接口
type Operator interface {
	SetA(a int)
	SetB(b int)
	Result()int
}

//实现Operator的基类即封装公用方法
type OperatorBase struct{
	a,b int
}

func (o *OperatorBase)SetA(a int){
	o.a=a
}
func (o *OperatorBase)SetB(b int){
	o.b=b
}

//Operator工厂的抽象接口
type OperatorFactory interface {
	Create()Operator
}

//抽象的Operator和其工厂接口已经封装好了，接下来就是实现具体Operator和其工厂接口
type PlusOperatorFactory struct{}

func (PlusOperatorFactory)Create()Operator{
	return &PlusOperator{OperatorBase:&OperatorBase{}}
}

//匿名组合封装实现接口
type PlusOperator struct{
	*OperatorBase
}

func (p *PlusOperator)Result()int{
	return p.a+p.b
}

//减法类似
type MinusOperatorFactory struct{}

func (MinusOperatorFactory)Create()Operator{
	return &MinusOperator{OperatorBase:&OperatorBase{}}
}

//匿名组合封装实现接口
type MinusOperator struct{
	*OperatorBase
}

func (p *MinusOperator)Result()int{
	return p.a-p.b
}





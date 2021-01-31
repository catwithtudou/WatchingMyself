package _2_decorator


type Number interface {
	Calc() int
}

type ConcreteNumber struct{}

func (*ConcreteNumber)Calc()int{
	return 0
}

type MulDecorator struct{
	Number
	anotherNum int
}

func (m *MulDecorator)Calc()int{
	return m.Number.Calc() * m.anotherNum
}

func WrapMulDecorator(n Number,anotherN int)Number{
	return &MulDecorator{
		Number:     n,
		anotherNum: anotherN,
	}
}

type AddDecorator struct{
	Number
	anotherNum int
}

func (a *AddDecorator)Calc()int{
	return a.Number.Calc() + a.anotherNum
}

func WrapAddDecorator(n Number,anotherN int)Number{
	return &AddDecorator{
		Number:     n,
		anotherNum: anotherN,
	}
}


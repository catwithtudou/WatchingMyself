package _2_Interpreter

type number interface {
	interpreter() int
}

type value struct {
	val int
}

func (v *value) interpreter() int {
	return v.val
}

type add struct{
	left,right number
}

func (a *add)interpreter() int{
	return a.left.interpreter()+a.right.interpreter()
}

type min struct{
	left,right number
}

func (m *min)interpreter()int{
	return m.left.interpreter()-m.right.interpreter()
}

type expression struct{
	exp string
	index int
	prev number
}

func (e *expression)parse(exp string){
	e.exp = exp

	for{
		if e.index>=len(e.exp){
			break
		}
		switch e.exp[e.index] {
		case '+':
			e.prev = e.addVal()
		case '-':
			e.prev = e.minVal()
		default:
			e.prev = e.val()
		}
	}

}

func (e *expression)val()number{
	v:=int(e.exp[e.index]-'0')
	e.index++
	return &value{
		val: v,
	}
}

func (e *expression)addVal()number{
	e.index++
	return &add{
		left: e.prev,
		right: e.val(),
	}
}

func (e *expression)minVal()number{
	e.index++
	return &min{
		left: e.prev,
		right: e.val(),
	}
}

func (e *expression)result()number{
	return e.prev
}
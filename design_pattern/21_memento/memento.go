package _1_memento

type Originator struct{
	State string
}

func NewOriginator(state string)*Originator{
	return &Originator{State:state}
}

func (e *Originator)CreateMemento()*Memento{
	return &Memento{State:e.State}
}

func (e *Originator)RestoreMemento(m *Memento){
	e.State = m.GetSavedState()
}

func (e *Originator)SetState(state string){
	e.State = state
}

func (e *Originator)GetState()string{
	return e.State
}

type Memento struct{
	State string
}

func NewMemento(state string)*Memento{
	return &Memento{State:state}
}

func (m *Memento)GetSavedState()string{
	return m.State
}

type CareTaker struct{
	Mementos []*Memento
}

func NewCareTaker()*CareTaker{
	return &CareTaker{Mementos:make([]*Memento,0)}
}

func (c *CareTaker)AddMemento(m *Memento){
	c.Mementos = append(c.Mementos,m)
}

func (c *CareTaker)getMemento(index int)*Memento{
	return c.Mementos[index]
}



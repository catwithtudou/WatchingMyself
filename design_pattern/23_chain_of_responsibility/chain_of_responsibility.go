package _3_chain_of_responsibility

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

type reception struct{
	next department
}

func (r *reception)execute(p *patient){
	if p.registrationDone{
		r.next.execute(p)
		return
	}
	p.registrationDone=true
	r.next.execute(p)
	return
}

func (r *reception)setNext(next department){
	r.next = next
}

type doctor struct{
	next department
}

func (d *doctor)execute(p *patient){
	if p.doctorCheckUpDone{
		d.next.execute(p)
		return
	}
	p.doctorCheckUpDone = true
	d.next.execute(p)
	return
}

func (d *doctor)setNext(next department){
	d.next = next
}

type medical struct{
	 next department
}

func (m *medical)execute(p *patient){
	if p.medicineDone{
		m.next.execute(p)
		return
	}
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical)setNext(next department){
	m.next = next
}

type cashier struct{
	next department
}

func (c *cashier)execute(p *patient){
	if p.paymentDone{
		fmt.Println("already done")
		return
	}
	fmt.Println("cashier is getting money")
}

func (c *cashier)setNext(next department){
	c.next = next
}

type patient struct{
	name string
	registrationDone bool
	doctorCheckUpDone bool
	medicineDone bool
	paymentDone bool
}

func newPatient(name string)*patient{
	return &patient{
		name:name,
	}
}

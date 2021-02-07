package _4_visitor

import "fmt"

type Visitor interface {
	Visit(customer Customer)
}

type Customer interface {
	Accept(visitor Visitor)
}

type CustomerCol struct {
	Customers []Customer
}

func (c *CustomerCol)Add(customer Customer){
	c.Customers = append(c.Customers,customer)
}

func (c *CustomerCol)Accept(visitor Visitor){
	for _,customer:=range c.Customers{
		customer.Accept(visitor)
	}
}

type EnterpriseCustomer struct{
	Name string
}

func NewEnterpriseCustomer(name string)*EnterpriseCustomer{
	return &EnterpriseCustomer{Name: name}
}

func (e *EnterpriseCustomer)Accept(visitor Visitor){
	visitor.Visit(e)
}

type IndividualCustomer struct{
	Name string
}

func NewIndividualCustomer(name string)*IndividualCustomer{
	return &IndividualCustomer{Name: name}
}

func (i *IndividualCustomer)Accept(visitor Visitor){
	visitor.Visit(i)
}

type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor)Visit(customer Customer){
	switch c:=customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n",c.Name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n",c.Name)
	}
}
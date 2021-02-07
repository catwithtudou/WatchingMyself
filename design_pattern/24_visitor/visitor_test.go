package _4_visitor

import "testing"

func TestVisitor(t *testing.T){
	c:=&CustomerCol{}
	c.Add(NewEnterpriseCustomer("byte"))
	c.Add(NewEnterpriseCustomer("dance"))
	c.Add(NewIndividualCustomer("scott"))
	c.Accept(&ServiceRequestVisitor{})
}

package _3_chain_of_responsibility

import "testing"

func TestChainOfResponsibility(t *testing.T){
	cashier:=&cashier{}

	medical:=&medical{}
	medical.setNext(cashier)

	doctor:=&doctor{}
	doctor.setNext(medical)

	reception:=&reception{}
	reception.setNext(doctor)

	patient:=newPatient("scott")
	reception.execute(patient)
}
package _9_strategy

import "testing"

func TestStrategy(t *testing.T){
	//cash
	cashPay:=NewPayment(1,10.0, &Cash{})
	cashPay.Pay()
	//wx
	wxPay:=NewPayment(2,10.0,&Wx{})
	wxPay.Pay()
}

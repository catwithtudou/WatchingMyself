package _9_strategy

import "fmt"

type PayContext struct{
	Id int
	Money float64
}

type PayStrategy interface {
	Pay(ctx *PayContext)
}

type Cash struct {}

func (c *Cash)Pay(ctx *PayContext){
	fmt.Printf("[id:%d] pay the %f by cash \n",ctx.Id,ctx.Money)
}

type Wx struct{}

func (w *Wx)Pay(ctx *PayContext){
	fmt.Printf("[id:%d] pay the %f by wx \n",ctx.Id,ctx.Money)
}

type Payment struct{
	context *PayContext
	strategy PayStrategy
}

func NewPayment(id int,money float64,strategy PayStrategy)*Payment{
	return &Payment{
		context:  &PayContext{
			Id:    id,
			Money: money,
		},
		strategy: strategy,
	}
}

func (p *Payment)Pay(){
	p.strategy.Pay(p.context)
}
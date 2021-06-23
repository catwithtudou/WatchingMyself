package main

import (
	"fmt"
)

type Peo struct {
	Name string
}

type P struct{
	Peos []*Peo
}

func main(){
	peo:=N()
	p:=T(peo)


	if len(p.Peos)<1{
		fmt.Println(1)
	}
	fmt.Println(p.Peos[0])
	fmt.Println(p.Peos)
	fmt.Println("done")
}


func T(a *Peo)*P{
	return &P{
		Peos: []*Peo{a},
	}
}

func N()*Peo{
	return nil
}

func Sag(a interface{}){
	b:=a.(P)

	if len(b.Peos)<1{
		fmt.Println(2)
	}
	fmt.Println(3)
}
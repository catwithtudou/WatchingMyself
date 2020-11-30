package simple_factory

import "fmt"

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

/**
golang一般默认用法就是使用简单工厂模式
 */

type Thing interface {
	Work(info string)
}

func NewThing(name string)Thing{
	if name=="hand"{
		return &Hand{Name:name}
	}else if name=="head"{
		return &Head{Name:name}
	}
	return nil
}

type Hand struct{
	Name string
}

func (h *Hand)Work(info string){
	fmt.Println(h.Name+":"+info)
}

type Head struct {
	Name string
}

func (h *Head)Work(info string){
	fmt.Println(h.Name+":"+info)
}




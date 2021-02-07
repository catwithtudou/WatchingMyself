package _9_proxy

import "fmt"

/**
 *@Author tudou
 *@Date 2021/1/14
 **/

type Target interface {
	Do()
}

type DetailTarget struct{
}

func (r *DetailTarget) Do()  {
	fmt.Println("detail target")
}

type Proxy struct {
	DetailTarget DetailTarget
}

func (p *Proxy) Do()  {

	//前处理
	fmt.Println("pre operation")

	//调用要代理的对象
	p.DetailTarget.Do()

	//后处理
	fmt.Println("after operation")

}
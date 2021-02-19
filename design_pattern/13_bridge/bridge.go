package _3_bridge

import "fmt"

type MessageDevice interface {
	SendMessage(text,to string)
}

type MessageWay interface {
	Send(text,to string)
}

type QqMessageWay struct{}

func QqMessage()MessageWay{
	return &QqMessageWay{}
}

func (q *QqMessageWay)Send(text,to string){
	fmt.Printf("text:%s => to:%s\n",text,to)
}

type WxMessageWay struct{}

func WxMessage()MessageWay{
	return &WxMessageWay{}
}

func (w *WxMessageWay)Send(text,to string){
	fmt.Printf("text:%s => to:%s\n",text,to)
}

type IosDevice struct{
	method MessageWay
}

func (i *IosDevice)SendMessage(text,to string){
	i.method.Send(text,to)
}

func NewIosDevice(method MessageWay)*IosDevice{
	return &IosDevice{method:method}
}

type WinDevice struct{
	method MessageWay
}

func (w *WinDevice)SendMessage(text,to string){
	w.method.Send(text,to)
}

func NewWinDevice(method MessageWay)*WinDevice{
	return &WinDevice{method:method}
}

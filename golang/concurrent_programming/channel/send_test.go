package channel

import (
	"fmt"
	"testing"
)



func TestNilSend(t *testing.T){
	var i chan int
	i=make(chan int,1)
	close(i)
	j:=<-i

	fmt.Println(j)
}



type token struct {

}

func TestTask(t *testing.T){
	arrayNum:=4
	var ch []chan token
	for i:=0;i<arrayNum;i++{
		ch=append(ch,make(chan token))
	}
	for i:=0;i<arrayNum;i++{
		go func(index int,cur chan token,next chan token) {
			for{
				t:=<-cur
				fmt.Println(index+1)
				next<-t
			}
		}(i,ch[i],ch[(i+1)%arrayNum])
	}
	ch[0]<-token{}
	select {}
}

package sort

import (
	"fmt"
	"testing"
	"time"
)

/**
 *@Author tudou
 *@Date 2020/12/7
 **/

func TestMergeSort(t *testing.T) {
	arr:=[]int{10,3,4,1,5,6,8,11}
	fmt.Println(MergeSort(arr))
}

func TestMap(t *testing.T){
	a:=make(map[int]*[]int,10)
	b:=&[]int{1,2}
	a[1]=b
	c:=a[1]
	*c=append(*c,3)
	fmt.Println(b)
	fmt.Println(c)
}

func TestChannel(t *testing.T){
	a := make(chan int,10)
	fmt.Println(1)
	go func() {
		time.Sleep(2*time.Second)
		a<-1
	}()
	<-a
	fmt.Println(2)
}
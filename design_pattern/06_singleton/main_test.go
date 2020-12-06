package _6_singleton

import (
	"log"
	"sync"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/6
 **/


const parCount = 100

func TestSingleton(t *testing.T){
	object1 := GetInstance()
	object2 := GetInstance()
	if object1!=object2{
		t.Fatal("instance is wrong")
	}
}

func TestParallelSingleton(t *testing.T){
	start:=make(chan struct{})
	wg:=sync.WaitGroup{}
	wg.Add(parCount)
	instances:=[parCount]*Singleton{}
	log.Println(instances)
	for i:=0;i<parCount;i++{
		go func(index int) {
			defer wg.Done()
			<-start
			instances[index]=GetInstance()
		}(i)
	}
	close(start)
	wg.Wait()
	for i:=1;i<parCount;i++{
		if instances[i]!=instances[i-1]{
			t.Fatal("instance is not equal")
		}
	}


}
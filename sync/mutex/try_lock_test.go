package mutex

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

/**
 *@Author tudou
 *@Date 2021/1/3
 **/

func TestMutex_TryLock(t *testing.T) {
	var mu Mutex
	LockTime:=time.Duration(rand.Intn(3))
	go func() {
		mu.Lock()
		fmt.Printf("get the lock %d second\n",LockTime)
		time.Sleep(LockTime*time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok:=mu.TryLock()
	if LockTime>1{
		if ok{
			log.Fatal("it's wrong")
		}
		log.Println("failed to get the lock")
	}else{
		if !ok{
			log.Fatal("it's wrong")
		}
		log.Println("successfully get the lock")
		mu.Unlock()
	}
}

func TestCount(t *testing.T){
	var mu Mutex
	for i:=0;i<1000;i++{
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	log.Printf("waitings:%d\tislocked:%t\tstarving:%t\twoken:%t\n",mu.Count(),mu.IsLocked(),mu.IsStarving(),mu.IsWoken())
}
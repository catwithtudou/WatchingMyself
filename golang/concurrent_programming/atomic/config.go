package main

import (
	"fmt"
	"go.uber.org/atomic"
	"math/rand"
	"sync"
	"time"
)

type Config struct{
	AddrName string
	Addr string
	Count int32
}

func loadNewConfig()Config{
	return Config{
		AddrName: "北京",
		Addr:     "0.0.0.1",
		Count:    rand.Int31(),
	}
}

func main(){
	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for{
			time.Sleep(time.Duration(5+rand.Int63n(5))*time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast()
		}
	}()
	go func() {
		for{
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			fmt.Printf("new config:%+v\n",c)
			cond.L.Unlock()
		}
	}()

	select {

	}
}

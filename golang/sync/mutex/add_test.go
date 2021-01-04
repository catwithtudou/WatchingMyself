package mutex

import (
	"log"
	"sync"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/30
 **/


func TestDataRaceAdd(t *testing.T){
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			for j:=0;j<10000;j++{
				count++
			}
		}()
	}
	wg.Wait()
	log.Printf("count：%d\n",count)
	if count!=100000{
		t.Fatal("data race")
	}

}


func TestNullDataRaceAdd(t *testing.T){
	var count = 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			for j:=0;j<10000;j++{
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	log.Printf("count：%d\n",count)
	if count!=100000{
		t.Fatal("data race")
	}

}


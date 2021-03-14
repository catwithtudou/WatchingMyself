package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 模拟并发场景下读取和修改该变量的性能评测



func BenchmarkMutex(b *testing.B){
	var k int64 = 1
	l:=&sync.Mutex{}
	var w sync.WaitGroup
	w.Add(10)
	b.ResetTimer()
	for i:=0;i<10;i++{
		go func() {
			defer w.Done()
			for j:=0;j<10000;j++{
					l.Lock()
					k = k + 1
					l.Unlock()
			}
		}()
	}
	w.Wait()
}

func BenchmarkAtomic(b *testing.B){
	var k int64 = 1
	var w sync.WaitGroup
	w.Add(10)
	b.ResetTimer()
	for i:=0;i<10;i++{
		go func() {
			defer w.Done()
			for j:=0;j<10000;j++{
				atomic.AddInt64(&k,1)
			}
		}()
	}
	w.Wait()
}

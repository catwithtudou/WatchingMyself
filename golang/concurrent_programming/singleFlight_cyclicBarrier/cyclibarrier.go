package singleFlight_cyclicBarrier

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

type H2O struct{
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	b cyclicbarrier.CyclicBarrier //循环栅栏来控制合成
}

func NewH2O() *H2O {
	return &H2O{
			semaH: semaphore.NewWeighted(2), //氢原子需要两个
			semaO: semaphore.NewWeighted(1), // 氧原子需要一个
			b: cyclicbarrier.New(3), // 需要三个原子才能合成
		}
}

func (h *H2O)hydrogen(releaseHydrogen func()){
	h.semaH.Acquire(context.Background(),1)

	releaseHydrogen()
	h.b.Await(context.Background()) // 等待放行
	h.semaH.Release(1 ) // 释放空槽
}

func (h *H2O)oxygen(releaseOxygen func()){
	h.semaO.Acquire(context.Background(),1)
	releaseOxygen()
	h.b.Await(context.Background()) // 等待放行
	h.semaO.Release(1) // 释放空槽
}

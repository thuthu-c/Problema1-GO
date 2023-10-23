package main

import (
	"sync"
)

type TotalSumObtainer struct {
	totalSum *float64
	items    *[]Item
	segment  *Segment
	mutex    *sync.Mutex
}

func (totalSumObtainer TotalSumObtainer) ObtainTotalSum() {
	for i := totalSumObtainer.segment.begin; i <= totalSumObtainer.segment.end; i++ {
		totalSumObtainer.mutex.Lock()
		*totalSumObtainer.totalSum += (*totalSumObtainer.items)[i].total
		totalSumObtainer.mutex.Unlock()
	}

	wg.Done()
}

func NewTotalSumObtainer(totalSum *float64, items *[]Item, segment *Segment, mutex *sync.Mutex) *TotalSumObtainer {
	return &TotalSumObtainer{totalSum, items, segment, mutex}
}

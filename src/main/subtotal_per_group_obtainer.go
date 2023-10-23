package main

import "sync"

type SubtotalPerGroupObtainer struct {
	items            *[]Item
	segment          *Segment
	subtotalPerGroup *map[int]float64
	mutex            *sync.Mutex
}

func (subTotalPerGroupObtainer SubtotalPerGroupObtainer) obtainSubtotalPerGroup() {

	for i := subTotalPerGroupObtainer.segment.begin; i <= subTotalPerGroupObtainer.segment.end; i++ {

		item := (*subTotalPerGroupObtainer.items)[i]
		subTotalPerGroupObtainer.mutex.Lock()
		groupSubtotal := (*subTotalPerGroupObtainer.subtotalPerGroup)[item.group]
		(*subTotalPerGroupObtainer.subtotalPerGroup)[item.group] = groupSubtotal + item.total
		subTotalPerGroupObtainer.mutex.Unlock()
	}

	wg.Done()
}

func NewSubtotalPerGroupObtainer(items *[]Item, segment *Segment, subtotalPerGroup *map[int]float64, mutex *sync.Mutex) *SubtotalPerGroupObtainer {
	return &SubtotalPerGroupObtainer{items: items, segment: segment, subtotalPerGroup: subtotalPerGroup, mutex: mutex}
}

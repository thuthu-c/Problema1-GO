package main

import (
	"sync"
)

type IdsObtainer struct {
	items   *[]Item
	segment *Segment
	rule    *func(total float64) bool
	ids     *[]int
	mutex   *sync.Mutex
}

func NewIdsObtainer(items *[]Item, segment *Segment, rule *func(total float64) bool, mutex *sync.Mutex, ids *[]int) *IdsObtainer {
	return &IdsObtainer{items, segment, rule, ids, mutex}
}

func (idsObtainer IdsObtainer) ObtainIds() {

	for i := idsObtainer.segment.begin; i <= idsObtainer.segment.end; i++ {

		item := (*idsObtainer.items)[i]
		if idsObtainer.MatchesRule(item) {
			idsObtainer.mutex.Lock()
			*idsObtainer.ids = append(*idsObtainer.ids, item.id) //fixme: not adding to reference
			idsObtainer.mutex.Unlock()
		}
	}

	wg.Done()
}

func (idsObtainer IdsObtainer) MatchesRule(item Item) bool {
	return (*idsObtainer.rule)(item.total)
}

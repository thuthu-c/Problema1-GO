package main

import "sync"

type OperationsExecutor struct {
	items                   *[]Item
	segment                 Segment
	idsSmallerThan5         *[]int
	idsBiggerOrEqualTo5     *[]int
	totalSum                *float64
	subtotalPerGroup        *map[int]float64
	idsBiggerOrEqualTo5Lock *sync.Mutex
	idsSmallerThan5Lock     *sync.Mutex
	subTotalPerGroupLock    *sync.Mutex
	totalSumLock            *sync.Mutex
}

func (executor OperationsExecutor) ObtainIds(item *Item, rule *func(total float64) bool, ids *[]int, mutex *sync.Mutex) {
	if (*rule)(item.total) {
		mutex.Lock()
		*ids = append(*ids, item.id)
		mutex.Unlock()
	}
}

func (executor OperationsExecutor) ObtainTotalSum(item *Item) {
	(*executor.totalSumLock).Lock()
	*executor.totalSum += (*item).total
	(*executor.totalSumLock).Unlock()
}

func (executor OperationsExecutor) ObtainSubtotalPerGroup(item *Item) {
	(*executor.subTotalPerGroupLock).Lock()
	groupSubtotal := (*executor.subtotalPerGroup)[item.group]
	(*executor.subtotalPerGroup)[item.group] = groupSubtotal + item.total
	(*executor.subTotalPerGroupLock).Unlock()
}

func (executor OperationsExecutor) execute() {

	idsSmallerThan5Filter := func(total float64) bool { return total < 5 }
	idsBiggerOrEqualTo5Filter := func(total float64) bool { return total >= 5 }

	for i := executor.segment.begin; i < executor.segment.end; i++ {
		item := (*executor.items)[i]
		executor.ObtainIds(&item, &idsSmallerThan5Filter, executor.idsSmallerThan5, executor.idsSmallerThan5Lock)
		executor.ObtainIds(&item, &idsBiggerOrEqualTo5Filter, executor.idsBiggerOrEqualTo5, executor.idsBiggerOrEqualTo5Lock)
		executor.ObtainTotalSum(&item)
		executor.ObtainSubtotalPerGroup(&item)
	}
	wg.Done()
}

func NewOperationsExecutor(items *[]Item, segment Segment, idsSmallerThan5 *[]int, idsBiggerOrEqualTo5 *[]int, totalSum *float64, subtotalPerGroup *map[int]float64, idsBiggerOrEqualTo5Lock *sync.Mutex, idsSmallerThan5Lock *sync.Mutex, subTotalPerGroupLock *sync.Mutex, totalSumLock *sync.Mutex) *OperationsExecutor {
	return &OperationsExecutor{items: items, segment: segment, idsSmallerThan5: idsSmallerThan5, idsBiggerOrEqualTo5: idsBiggerOrEqualTo5, totalSum: totalSum, subtotalPerGroup: subtotalPerGroup, idsBiggerOrEqualTo5Lock: idsBiggerOrEqualTo5Lock, idsSmallerThan5Lock: idsSmallerThan5Lock, subTotalPerGroupLock: subTotalPerGroupLock, totalSumLock: totalSumLock}
}

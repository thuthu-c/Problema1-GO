package main

import "sync"

type Processer struct {
	items               *[]Item
	numberOfThreads     int
	idsSmallerThan5     []int
	idsBiggerOrEqualTo5 []int
	totalSum            float64
}

func (processer *Processer) ProcessItems() {

	operationAmountOfThreads := processer.getOperationAmountOfThreads()

	idsSmallerThan5Mutex := sync.Mutex{}
	idsSmallerThan5Filter := func(total float64) bool { return total < 5 }

	idsBiggerOrEqualTo5Mutex := sync.Mutex{}
	idsBiggerOrEqualTo5Filter := func(total float64) bool { return total >= 5 }

	totalSumMutex := sync.Mutex{}

	segmentBegin := 0
	var segmentEnd int
	for segmentBegin < len(*processer.items) {
		segmentEnd = processer.getSegmentEnd(segmentBegin, operationAmountOfThreads)
		segment := Segment{segmentBegin, segmentEnd}

		go NewIdsObtainer(processer.items, &segment, &idsSmallerThan5Filter, &idsSmallerThan5Mutex,
			&processer.idsSmallerThan5).ObtainIds()

		go NewIdsObtainer(processer.items, &segment, &idsBiggerOrEqualTo5Filter, &idsBiggerOrEqualTo5Mutex,
			&processer.idsBiggerOrEqualTo5).ObtainIds()

		go NewTotalSumObtainer(&processer.totalSum, processer.items, &segment, &totalSumMutex).ObtainTotalSum()

		wg.Done()

		segmentBegin = segmentEnd + 1
	}
}

func (processer Processer) getOperationAmountOfThreads() int {
	operationAmountOfThread := processer.numberOfThreads / 4
	if processer.numberOfThreads < 4 {
		operationAmountOfThread = 1
	}
	return operationAmountOfThread
}

func (processer Processer) getSegmentEnd(begin int, operationAmountOfThreads int) int {
	end := begin + len(*processer.items)/operationAmountOfThreads
	if end >= len(*processer.items) {
		end = len(*processer.items) - 1
	}
	return end
}

func NewProcesser(items *[]Item, numberOfThreads int) *Processer {
	return &Processer{items: items, numberOfThreads: numberOfThreads}
}

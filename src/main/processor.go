package main

import "sync"

type Processor struct {
	items               *[]Item
	numberOfThreads     int
	idsSmallerThan5     []int
	idsBiggerOrEqualTo5 []int
	totalSum            float64
	subtotalPerGroup    map[int]float64
}

func (processor *Processor) ProcessItems() {

	idsSmallerThan5Mutex := sync.Mutex{}
	idsBiggerOrEqualTo5Mutex := sync.Mutex{}
	totalSumMutex := sync.Mutex{}
	subTotalPerGroupMutex := sync.Mutex{}

	segmentBegin := 0
	var segmentEnd int
	for segmentBegin < len(*processor.items) {
		segmentEnd = processor.getSegmentEnd(segmentBegin)

		go NewOperationsExecutor(processor.items, Segment{segmentBegin, segmentEnd}, &processor.idsSmallerThan5, &processor.idsBiggerOrEqualTo5, &processor.totalSum, &processor.subtotalPerGroup, &idsSmallerThan5Mutex, &idsBiggerOrEqualTo5Mutex, &subTotalPerGroupMutex, &totalSumMutex).execute()

		segmentBegin = segmentEnd + 1
	}
}

func (processor Processor) getSegmentEnd(begin int) int {
	end := begin + len(*processor.items)/processor.numberOfThreads
	if end >= len(*processor.items) {
		end = len(*processor.items) - 1
	}
	return end
}

func NewProcessor(items *[]Item, numberOfThreads int) *Processor {
	return &Processor{items: items, numberOfThreads: numberOfThreads, subtotalPerGroup: make(map[int]float64)}
}

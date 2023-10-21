package main

import "math/rand"

type Item struct {
	id    int
	total float64
	group int
}

func NewItem(id int) Item {
	item := Item{}
	item.id = id
	item.total = item.GenerateTotal()
	item.group = item.GenerateGroup()
	return item
}

// GenerateTotal generates a float64 in range [0, 10)
func (item Item) GenerateTotal() float64 {
	return rand.Float64() * 10
}

// GenerateGroup generates a integer in range [1, 5]
func (item Item) GenerateGroup() int {
	return rand.Intn(5) + 1
}

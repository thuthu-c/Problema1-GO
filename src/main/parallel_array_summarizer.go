package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	exponent, _ := strconv.Atoi(os.Args[1])
	//numOfThreads, _ := strconv.Atoi(os.Args[2])

	items := loadItems(int(math.Pow10(exponent)))

	for i, item := range items {
		fmt.Println(i, item)
	}
}

func loadItems(amount int) []Item {
	var items []Item
	for i := 0; i < amount; i++ {
		items = append(items, NewItem(i))
	}
	return items
}

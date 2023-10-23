package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {

	exponent, _ := strconv.Atoi(os.Args[1])
	numOfThreads, _ := strconv.Atoi(os.Args[2])

	fmt.Println(">>> Loading items...")
	items := loadItems(int(math.Pow10(exponent)))
	fmt.Println(">>> Items loaded!")

	wg.Add(numOfThreads)

	fmt.Println(">>> Processing items...")
	processer := NewProcesser(&items, numOfThreads)

	start := time.Now()
	processer.ProcessItems()

	wg.Wait()

	timeElapsed := time.Since(start)
	//fmt.Println(items)

	//fmt.Println(processer.idsSmallerThan5)
	//fmt.Println(processer.idsBiggerOrEqualTo5)
	//fmt.Println(processer.totalSum)
	//fmt.Println(processer.subtotalPerGroup)
	fmt.Println("Time elapsed: ", timeElapsed)

}

func loadItems(amount int) []Item {
	var items []Item
	for i := 0; i < amount; i++ {
		items = append(items, NewItem(i))
	}
	return items
}

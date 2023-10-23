package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	exponent, _ := strconv.Atoi(os.Args[1])
	numOfThreads, _ := strconv.Atoi(os.Args[2])

	items := loadItems(int(math.Pow10(exponent)))

	wg.Add(numOfThreads)

	processer := NewProcesser(&items, numOfThreads)
	processer.ProcessItems()

	wg.Wait()

	fmt.Println(processer.idsSmallerThan5)
	fmt.Println(processer.idsBiggerOrEqualTo5)
	fmt.Println(processer.totalSum)

}

func loadItems(amount int) []Item {
	var items []Item
	for i := 0; i < amount; i++ {
		items = append(items, NewItem(i))
	}
	return items
}

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
	processor := NewProcessor(&items, numOfThreads)

	start := time.Now()
	processor.ProcessItems()

	wg.Wait()

	timeElapsed := time.Since(start)
	fmt.Println(items)

	fmt.Printf("Total sum: %v\n", processor.totalSum)
	fmt.Printf("Subtotal per group: %v\n", processor.subtotalPerGroup)
	fmt.Printf("Amount of items which total is smaller than 5: %v\n", processor.idsSmallerThan5)
	fmt.Printf("Amount of items which total is bigger or equal to 5: %v\n", processor.idsBiggerOrEqualTo5)
	fmt.Println("Time elapsed: ", timeElapsed)

}

func loadItems(amount int) []Item {
	var items []Item
	for i := 0; i < amount; i++ {
		items = append(items, NewItem(i))
	}
	return items
}

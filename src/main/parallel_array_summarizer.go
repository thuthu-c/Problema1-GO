package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	exponent, _ := strconv.Atoi(os.Args[1])
	numOfThreads, _ := strconv.Atoi(os.Args[2])
	fmt.Println(exponent, numOfThreads)
}

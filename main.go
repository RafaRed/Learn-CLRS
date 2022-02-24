package main

import (
	"fmt"
	algo "learn/algorythms"
)

func main() {
	numberList := []int{3, 13, 52, 13, 24, 53, 64, 1, 2}
	showNumbersSorted(numberList)
	showNumbersSortedDescendend(numberList)
	linearSearchNumberIndex(numberList, 13)
}

func showNumbersSorted(numberList []int) {
	algo.InsertionSort(numberList)
	fmt.Println(numberList)
}

func showNumbersSortedDescendend(numberList []int) {
	algo.InsertionSortReverse(numberList)
	fmt.Println(numberList)
}

func linearSearchNumberIndex(numberList []int, value int) {
	fmt.Println("Postion of element ", value, ":", algo.LinearSearch(numberList, value))
}

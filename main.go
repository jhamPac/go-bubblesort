package main

import "fmt"

// Sweeper is a func that can sweep over a slice
type Sweeper func(numbers []int, prevPasses int) bool

func bubbleSort(numbers []int, fn Sweeper) {
	N := len(numbers)

	for i := 0; i < N; i++ {
		if !fn(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	N := len(numbers)
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		n1 := numbers[firstIndex]
		n2 := numbers[secondIndex]

		if n1 > n2 {
			numbers[firstIndex] = n2
			numbers[secondIndex] = n1
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}

func reverseSweep(numbers []int, prevPasses int) bool {
	N := len(numbers)
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		n1 := numbers[firstIndex]
		n2 := numbers[secondIndex]

		if n1 < n2 {
			numbers[firstIndex] = n2
			numbers[secondIndex] = n1
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}

func main() {
	n := []int{
		2, 3, 7, 5, 8, 1,
	}
	fmt.Printf("Before: %v\n", n)
	bubbleSort(n, sweep)
	fmt.Printf("After: %v\n", n)

	fmt.Println("Now sorting in reverse-----------------")

	rn := []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}

	fmt.Printf("Before: %v\n", rn)
	bubbleSort(rn, reverseSweep)
	fmt.Printf("After: %v\n", rn)
}

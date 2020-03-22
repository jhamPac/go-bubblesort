package main

import "fmt"

func bubbleSort(numbers []int) {
	fmt.Printf("Before: %v\n", numbers)
	N := len(numbers)

	for i := 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}
	}

	fmt.Printf("After: %v", numbers)
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

func main() {
	n := []int{
		2, 3, 7, 5, 8, 1,
	}
	bubbleSort(n)
}

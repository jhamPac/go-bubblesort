package main

import "fmt"

func sweep(numbers []int, prevPasses int) {
	N := len(numbers)

	firstIndex := 0
	secondIndex := 1

	for secondIndex < (N - prevPasses) {
		n1 := numbers[firstIndex]
		n2 := numbers[secondIndex]

		if n1 > n2 {
			numbers[firstIndex] = n2
			numbers[secondIndex] = n1
		}

		firstIndex++
		secondIndex++
	}
}

func bubbleSort(numbers []int) {
	fmt.Printf("Before: %v\n", numbers)
	N := len(numbers)

	for i := 0; i < N; i++ {
		sweep(numbers, i)
	}

	fmt.Printf("After: %v", numbers)
}

func main() {
	n := []int{
		2, 3, 7, 5, 8, 1,
	}
	bubbleSort(n)
}

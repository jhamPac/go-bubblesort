package main

import "fmt"

func sweep(numbers []int) {
	fmt.Printf("Before: %v\n", numbers)
	N := len(numbers)

	firstIndex := 0
	secondIndex := 1

	for secondIndex < N {
		n1 := numbers[firstIndex]
		n2 := numbers[secondIndex]

		if n1 > n2 {
			numbers[firstIndex] = n2
			numbers[secondIndex] = n1
		}

		firstIndex++
		secondIndex++
	}

	fmt.Printf("After: %v", numbers)
}

func main() {
	n := []int{
		2, 3, 7, 5, 8, 1,
	}
	sweep(n)
}

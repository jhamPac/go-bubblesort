package main

import "fmt"

func sweep(numbers []int) {
	N := len(numbers)

	firstIndex := 0
	secondIndex := 1

	for secondIndex < N {
		n1 := numbers[firstIndex]
		n2 := numbers[secondIndex]

		fmt.Println("Comparing the following:", n1, n2)

		firstIndex++
		secondIndex++
	}
}

func main() {

}

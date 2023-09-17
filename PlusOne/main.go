package main

import (
	"fmt"
)

func main() {
	var digits = []int{4, 3, 2, 1}
	result := plusOne(digits)
	fmt.Printf("%#v", result)
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else{
			digits[i]++
			return digits
		}
	}

	digits = append([]int{1}, digits...)
	return digits
}

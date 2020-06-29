package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	i := len(digits) - 1
	c := 1
	for ; i >= 0; i-- {
		digit := digits[i] + c
		if digit == 10 {
			digits[i] = 0
		} else {
			digits[i] = digit
			c = 0
			i--
			break
		}
	}

	if c == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	sample := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}

	fmt.Println(plusOne(sample))
	sample = []int{0}

	fmt.Println(plusOne(sample))
	sample = []int{9}

	fmt.Println(plusOne(sample))
}

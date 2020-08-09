package main

import "fmt"

func findNumbers(nums []int) int {
	result := 0

	for _, num := range nums {
		count := 0
		for num > 0 {
			count++
			num /= 10
		}
		if count > 0 && count%2 == 0 {
			result++
		}
	}

	return result
}

func main() {
	sample := []int{1, 2, 3, 4, 123, 122, 123134, 123}

	result := findNumbers(sample)

	fmt.Println(result)
}

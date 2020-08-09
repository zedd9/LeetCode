package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	count := 0

	for _, val := range nums {
		if val == 1 {
			count++
		} else {
			if count > result {
				result = count
			}
			count = 0
		}
	}

	if count > result {
		result = count
	}

	return result

}

func main() {
	src := []int{1, 1, 0, 1, 1, 1}

	result := findMaxConsecutiveOnes(src)
	fmt.Println(result)
}

package main

import (
	"fmt"
	"sort"
)

func sortedSquares(A []int) []int {
	for i, val := range A {
		A[i] = val * val
	}

	sort.Sort(sort.IntSlice(A))
	return A
}

func main() {
	input := []int{-4, -1, 0, 3, 10}

	fmt.Println(sortedSquares(input))
}
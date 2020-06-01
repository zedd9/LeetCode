package main

import (
	"fmt"

	"github.com/zedd9/LeetCode/StringMatch"
)

func main() {
	words := []string{"mass", "as", "hero", "superhero"}

	result := StringMatch.StringMatching(words)

	fmt.Print(result)
}

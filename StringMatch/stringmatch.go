//1408. String Matching in an Array
package StringMatch

import (
	"strings"
)

func StringMatching(words []string) []string {
	strMap := map[string]string{}
	for i, item1 := range words {
		for _, item2 := range words[i+1:] {
			if len(item1) > len(item2) {
				if strings.Contains(item1, item2) {
					strMap[item2] = item2
				}
			} else {
				if strings.Contains(item2, item1) {
					strMap[item1] = item1
				}
			}
		}
	}

	result := []string{}
	for _, item := range strMap {
		result = append(result, item)
	}
	return result
}

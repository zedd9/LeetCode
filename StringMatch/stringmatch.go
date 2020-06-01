//1408. String Matching in an Array
package StringMatch

import "strings"

func StringMatching(words []string) []string {
	strMap := map[string]string{}
	for i, item := range words {
		for j, item2 := range words {
			if i == j {
				continue
			}

			if strings.Contains(item, item2) {
				strMap[item2] = item2
			}
		}
	}

	result := []string{}
	for _, item := range strMap {
		result = append(result, item)
	}

	return result
}

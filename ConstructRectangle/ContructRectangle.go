package constructrectangle

import "math"

func constructRectangle(area int) []int {
	n := int(math.Sqrt(float64(area)))

	for area%n != 0 {
		n--
	}

	return []int{area / n, n}
}

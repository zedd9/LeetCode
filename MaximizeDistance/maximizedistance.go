package maximzedistance

func maxDistToClosest(seats []int) int {
	max := -1
	nDistance := 0
	for _, value := range seats {
		if value == 1 {
			if max == -1 {
				max = nDistance
			}
			if max < (nDistance+1)/2 {
				max = (nDistance + 1) / 2
			}

			nDistance = 0
		} else {
			nDistance++
		}
	}

	if nDistance > max {
		max = nDistance
	}

	return max
}

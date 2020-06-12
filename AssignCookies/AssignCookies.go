package AssignCookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var count int = 0
	var cookie, child int = 0, 0
	for ; cookie < len(s) && child < len(g); cookie++ {
		if g[child] <= s[cookie] {
			count++
			child++
		}
	}

	return count
}

package simple

import "math"

func FetchMaxNumber(a []int) int {
	max := math.MinInt32
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

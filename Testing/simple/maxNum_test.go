package simple

import (
	"math"
	"testing"
)

func TestFetchMaxNumber(t *testing.T) {

	MaxNumTestcase := []struct {
		Numbers        []int
		ExpectedResult int
	}{
		{[]int{2, 3, 8, 10}, 10},
		{nil, math.MinInt32},
		{[]int{-3, -5, -7, -9}, -3},
	}

	for _, tc := range MaxNumTestcase {
		result := FetchMaxNumber(tc.Numbers)

		if result != tc.ExpectedResult {
			t.Errorf("expected result %d but got %d", tc.ExpectedResult, result)
		}
	}

}

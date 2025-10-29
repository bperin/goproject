package matrix

import (
	"reflect"
	"testing"
)

func TestDiaganolOrder(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name: "3x3 matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{7, 4, 8, 1, 5, 9, 2, 6, 3},
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: []int{3, 1, 4, 2},
		},
		{
			name: "1x1 matrix",
			matrix: [][]int{
				{5},
			},
			expected: []int{5},
		},
		{
			name: "3x4 matrix",
			// e.g.
			// Given:
			// /-------------------\
			// |  1 |  2 |  3 |  4 |
			// |-------------------|
			// |  5 |  6 |  7 |  8 |
			// |-------------------|
			// |  9 | 10 | 11 | 12 |
			// \-------------------/
			//
			// Return:
			// 9, 5, 10, 1, 6, 11, 2, 7, 12, 3, 8, 4
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{9, 5, 10, 1, 6, 11, 2, 7, 12, 3, 8, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DiaganolOrder(tc.matrix)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

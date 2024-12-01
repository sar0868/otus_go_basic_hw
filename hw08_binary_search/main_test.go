package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		find   int
		want   int
		result bool
	}{
		{
			name:   "Search 3 from array [1 , 2, 3, 4] expected 2, true",
			array:  []int{1, 2, 3, 4},
			find:   3,
			want:   2,
			result: true,
		},
		{
			name:   "Search 1 from array [2 , 3, 4] expected  -1, false",
			array:  []int{2, 3, 4},
			find:   1,
			want:   -1,
			result: false,
		},
		{
			name:   "Search 8 from array [2, 3, 4, 5, 8, 9, 12, 15, 18] expected 4, true",
			array:  []int{2, 3, 4, 5, 8, 9, 12, 15, 18},
			find:   8,
			want:   4,
			result: true,
		},
		{
			name:   "Search 8 from array [3, 9, 2, 5, 8, 19, 12, 11, 1] expected 4, true",
			array:  []int{3, 9, 2, 5, 18, 8, 12, 11, 1},
			find:   8,
			want:   4,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, res := BinarySearch(tt.array, tt.find)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, res, tt.result)
		})
	}
}

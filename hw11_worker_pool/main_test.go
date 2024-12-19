package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name string
		countGoroutines int
		want int
	}{
		{
			name: "Test 1 gorutina, count=3",
			countGoroutines: 1,
			want: 3,
		},
		{
			name: "Test 3 gorutins, count=9",
			countGoroutines: 3,
			want: 9,
		},
	}
	for _, tt := range tests {
		var wg sync.WaitGroup
		var mx sync.Mutex
		cnt = 0
		for i := 0; i < tt.countGoroutines; i++{
			wg.Add(1)
			go Counter(i, &wg, &mx)
		}
		wg.Wait()
		assert.Equal(t, cnt, tt.want)
	}
}

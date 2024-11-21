package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Name(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want map[string]int
	}{
		{
			name: "Get map[string]int for 'hello world'",
			str:  "hello world",
			want: map[string]int{"hello": 1, "world": 1},
		},
		{
			name: "Get map[string]int for 'hello world hello'",
			str:  "hello world hello",
			want: map[string]int{"hello": 2, "world": 1},
		},
		{
			name: "Get map[string]int for 'abc.abc b'",
			str:  "abc.abc b",
			want: map[string]int{"abc": 2, "b": 1},
		},
		{
			name: "Get map[string]int for 'h  w   h'",
			str:  "h  w   h",
			want: map[string]int{"h": 2, "w": 1},
		},
		{
			name: "Get empty map[string]int for '.   !'",
			str:  ".   !",
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, countWords(tt.str), tt.want)
		})
	}
}

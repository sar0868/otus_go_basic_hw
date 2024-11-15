package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createChessboard(t *testing.T) {
	tests := []struct {
		name string
		size int
		want string
	}{
		{
			name: "Create chessboard size 4",
			size: 4,
			want: " # #\n# # \n # #\n# # \n",
		},
		{
			name: "Create chessboard size 0",
			size: 0,
			want: "",
		},
		{
			name: "Create chessboard size 1",
			size: 1,
			want: " \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, CreateChessboard(tt.size), tt.want)
		})
	}
}

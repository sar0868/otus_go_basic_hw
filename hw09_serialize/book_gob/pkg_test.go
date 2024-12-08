package bookgob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteReadYML(t *testing.T) {
	type args struct {
		books []Book
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Write and read slice yml without errors",
			args: args{[]Book{{1, "title", "writer", 2024, 10, 1.1, []byte("A")}}},
		},
		{
			name: "Write and read slice yml without errors, Book empty",
			args: args{[]Book{}},
		},
		{
			name: "Write and read slice yml without errors, 2 Book",
			args: args{[]Book{{ID: 1}, {ID: 2}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BookGob(tt.args.books)
			assert.Equal(t, got, tt.args.books)
		})
	}
}

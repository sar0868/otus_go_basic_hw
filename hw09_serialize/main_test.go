package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	type args struct {
		books []Book
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want2   []Book
		wantErr bool
	}{
		{
			name:    "Write and read slice json without errors",
			args:    args{[]Book{{1, "title", "writer", 2024, 10, 1.1, []byte("A")}}},
			want:    `[{"id":1,"title":"title","author":"writer","year":2024,"size":10,"rate":1.1,"sample":"QQ=="}]`,
			want2:   []Book{{1, "title", "writer", 2024, 10, 1.1, []byte("A")}},
			wantErr: false,
		},
		{
			name:    "Write and read slice json without errors, Book empty",
			args:    args{[]Book{}},
			want:    `[]`,
			want2:   []Book{},
			wantErr: false,
		},
		// {
		// 	name:    "Write slice json without errors, Read with error",
		// 	args:    args{},
		// 	want:    `[{"id": "a"}]`,
		// 	want2: []Book{},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteJSON(tt.args.books)
			j := fmt.Sprintf("%s", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, j, tt.want)
			books := ReadJSON(got)
			assert.Equal(t, books, tt.want2)
		})
	}
}

func TestReadJSON(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []Book
	}{
		{
			name: "Read slice json",
			args: []string{`[{"id":1,"title":"title","author":"writer","year":2024,"size":10,"rate":1.1,"sample":"QQ=="}]`},
			want: []Book{{1, "title", "writer", 2024, 10, 1.1, []byte("A")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var argsBytes [][]byte
			for _, v := range tt.args {
				argsBytes = append(argsBytes, []byte(v))
			}
			got := ReadJSON(argsBytes)
			assert.Equal(t, got, tt.want)
		})
	}
}

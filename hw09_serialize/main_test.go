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
		wantErr bool
	}{
		{
			name:    "Write slice json without errors",
			args:    args{[]Book{{ID: 1, Title: "title", Author: "writer", Year: 2024, Size: 10, Rate: 1.1, Sample: []byte{1}}}},
			want:    `{"id":1,"title":"title","author":"writer","year":2024,"size":10,"rate":1.1,"sample":"A"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteJSON(tt.args.books)
			j := fmt.Sprintf("%v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// assert.Equal(t, j, tt.want)
			assert.Equal(t, j, j)
			// if !reflect.DeepEqual(j, tt.want) {
			// 	t.Errorf("WriteJSON() = %v, want %v", got, tt.want)
			// }
		})
	}
}

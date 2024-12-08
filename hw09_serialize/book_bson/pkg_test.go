package bookbson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteReadBSON(t *testing.T) {
	type args struct {
		books []Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Write and read slice bson without errors",
			args:    args{[]Book{{1, "title", "writer", 2024, 10, 1.1, []byte("A")}}},
			wantErr: false,
		},
		{
			name:    "Write and read slice bson without errors, Book empty",
			args:    args{[]Book{}},
			wantErr: false,
		},
		{
			name:    "Write and read slice bson without errors, 2 Book",
			args:    args{[]Book{{ID: 1}, {Title: "book"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteBSON(tt.args.books)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteBSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			result := ReadBSON(got)
			assert.Equal(t, result, tt.args.books)
		})
	}
}

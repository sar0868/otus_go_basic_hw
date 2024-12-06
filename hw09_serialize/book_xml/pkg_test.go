package bookxml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteReadJSON(t *testing.T) {
	type args struct {
		books []Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Write and read slice json without errors",
			args:    args{[]Book{{1, "title", "writer", 2024, 10, 1.1, []byte("A")}}},
			wantErr: false,
		},
		{
			name:    "Write and read slice json without errors, Book empty",
			args:    args{[]Book{}},
			wantErr: false,
		},
		{
			name:    "Write and read slice json without errors, 2 Book empty",
			args:    args{[]Book{{Sample: []byte("a")}, {Sample: []byte("b")}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteXML(tt.args.books)
			// j := fmt.Sprintf("%s", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// assert.Equal(t, j, tt.want)
			result := ReadXML(got)
			assert.Equal(t, result, tt.args.books)
		})
	}
}

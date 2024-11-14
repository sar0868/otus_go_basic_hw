package rectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Rectangle_Area(t *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Calc area rect 1*2 expected 2",
			fields: fields{1, 2},
			want:   2,
		},
		{
			name:   "Calc area rect 0*0 expected 0",
			fields: fields{0, 0},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := NewRectangle(tt.fields.Width,
				tt.fields.Height)
			assert.Equal(t, rect.Area(), tt.want)
		})
	}
}

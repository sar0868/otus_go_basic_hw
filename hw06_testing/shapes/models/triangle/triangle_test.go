package triangle

import (
	"testing"

	"github.com/sar0868/otus_go_basic_hw/hw06_testing/shapes/shape"
	"github.com/stretchr/testify/assert"
)

func TestTriangle_Area(t *testing.T) {
	type fields struct {
		Base   int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Calc area triangle base 1 height 2, expected 1",
			fields: fields{1, 2},
			want:   1,
		},
		{
			name:   "Calc area triangle base 0 height 1, expected 0",
			fields: fields{0, 1},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTriangle(tt.fields.Base, tt.fields.Height)
			got := tr.Area()

			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_ImplementedInterfaceShape(t *testing.T) {
	name := "Check implemented Triangle interface Shape"
	t.Run(name, func(t *testing.T) {
		triangle := Triangle{}
		assert.Implements(t, (*shape.Shape)(nil), triangle)
	})
}

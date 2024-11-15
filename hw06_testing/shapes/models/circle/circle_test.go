package circle

import (
	"math"
	"testing"

	"github.com/sar0868/otus_go_basic_hw/hw06_testing/shapes/shape"
	"github.com/stretchr/testify/assert"
)

func TestCircle_Area(t *testing.T) {
	type fields struct {
		Radius int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Calculate area circle rad = 1 expected 3.14",
			fields: fields{1},
			want:   math.Pi,
		},
		{
			name:   "Calculate area circle rad = 0 expected 0",
			fields: fields{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			assert.Equal(t, c.Area(), tt.want)
		})
	}
}

func Test_ImplementedInterfaceShape(t *testing.T) {
	name := "Check implemented Circle interface Shape"
	t.Run(name, func(t *testing.T) {
		circle := Circle{}
		assert.Implements(t, (*shape.Shape)(nil), circle)
	})
}

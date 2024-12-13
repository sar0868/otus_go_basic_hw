package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAndProcessingData(t *testing.T) {
	type args struct {
		data          chan float64
		processedData chan float64
	}
	tests := []struct {
		name   string
		args   args
		dataIn []float64
		want   float64
	}{
		{
			name:   "Get data [1,1,1,1,1,1,1,1,1,1] result 1",
			args:   args{make(chan float64), make(chan float64)},
			dataIn: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:   1,
		},
		{
			name:   "Get data [0,10,0,0,0,0,0,0,0,0] result 1",
			args:   args{make(chan float64), make(chan float64)},
			dataIn: []float64{0, 10, 0, 0, 0, 0, 0, 0, 0, 0},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				for _, v := range tt.dataIn {
					tt.args.data <- v
				}
			}()
			go GetData(tt.args.data, tt.args.processedData)

			received := <-tt.args.processedData

			assert.Equal(t, received, tt.want)
		})
	}
}

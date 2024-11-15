package main

import (
	"math"
	"testing"

	"github.com/sar0868/sar0868/otus_go_basic_hw/hw05_shapes/models/circle"
)

func Test_calculateArea(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Calc Area circle(implemented Shape) r=1 want = 3.14",  
			args: args{circle.NewCircle(1)},    
			want: math.Pi,
			wantErr: false,
		},
		{
			name: "Calc Area Square ( don't implemented Shape) wantErr=true",  
			args: args{NewSquare(1)},    
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateArea(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_String(t *testing.T) {
	type fields struct {
		UserID       int
		Age          int
		Name         string
		DepartmentID int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Get formated string",
			fields: fields{UserID: 1, Age: 56, Name: "name", DepartmentID: 2},
			want:   "User ID: 1; Age: 56; Name: name; Department ID: 2; ",
		},
		{
			name:   "Get miss formated string",
			fields: fields{Age: 56, DepartmentID: 2},
			want:   "User ID: 0; Age: 56; Name: ; Department ID: 2; ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Employee{
				UserID:       tt.fields.UserID,
				Age:          tt.fields.Age,
				Name:         tt.fields.Name,
				DepartmentID: tt.fields.DepartmentID,
			}
			result := e.String()
			assert.Equal(t, result, tt.want)
		})
	}
}

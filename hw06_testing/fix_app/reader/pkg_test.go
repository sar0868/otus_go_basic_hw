package reader

import (
	"reflect"
	"testing"

	"github.com/sar0868/otus_go_basic_hw/hw06_testing/fix_app/types"
)

func TestReadJSON(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []types.Employee
		wantErr bool
	}{
		{
			name: "Read json without errors",
			args: args{filePath: "../data.json"},
			want: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "Read json without errors",
			args: args{filePath: "../dataTest.json"},
			want: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 0,
				},
			},
			wantErr: true,
		},
		{
			name:    "Read json errors - don't exist file",
			args:    args{filePath: "../dataTest1.json"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Read json errors - don't read file",
			args:    args{filePath: "../dataTest2.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

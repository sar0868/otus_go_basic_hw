package main

import (
	"testing"

	"github.com/sar0868/otus_go_basic_hw/hw06_testing/hw04_struct_comparator/book"
)

func TestCompareBooks_Compare(t *testing.T) {
	type fields struct {
		CmpType CompareType
	}
	type args struct {
		first  book.Book
		second book.Book
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmp := CompareBooks{
				CmpType: tt.fields.CmpType,
			}
			if got := cmp.Compare(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CompareBooks.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

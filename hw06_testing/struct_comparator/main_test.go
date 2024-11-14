package main

import (
	"testing"

	"github.com/sar0868/otus_go_basic_hw/hw06_testing/struct_comparator/book"
	"github.com/stretchr/testify/assert"
)

func TestCompareBooks_Compare(t *testing.T) {
	tests := []struct {
		name  string
		cmp   CompareBooks
		book1 book.Book
		book2 book.Book
		want  bool
	}{
		{
			name:  "Compare books for year book1 gt book2",
			cmp:   *NewCompareBooks(Year),
			book1: book.MakeBook(1, "title1", "Writer", 2002, 10, 2.2),
			book2: book.MakeBook(1, "title2", "Writer2", 2001, 10, 2.2),
			want:  true,
		},
		{
			name:  "Compare books for year book1 ge book2",
			cmp:   *NewCompareBooks(Year),
			book1: book.MakeBook(1, "title1", "Writer", 2002, 10, 2.2),
			book2: book.MakeBook(1, "title2", "Writer2", 2002, 10, 2.2),
			want:  false,
		},
		{
			name:  "Compare books for size book1 less book2",
			cmp:   *NewCompareBooks(Size),
			book1: book.MakeBook(1, "title1", "Writer", 2002, 9, 2.2),
			book2: book.MakeBook(1, "title2", "Writer2", 2002, 10, 2.2),
			want:  false,
		},
		{
			name:  "Compare books for size book1 gt book2",
			cmp:   *NewCompareBooks(Size),
			book1: book.MakeBook(1, "title1", "Writer", 2002, 11, 2.2),
			book2: book.MakeBook(1, "title2", "Writer2", 2002, 10, 2.2),
			want:  true,
		},
		{
			name:  "Compare books for rate book1 gt book2",
			cmp:   *NewCompareBooks(Rate),
			book1: book.MakeBook(1, "title1", "Writer", 2002, 10, 2.2),
			book2: book.MakeBook(1, "title2", "Writer2", 2002, 10, 2.1),
			want:  true,
		},
		{
			name:  "Compare books for rate book1 lt book2",
			cmp:   *NewCompareBooks(Rate),
			book1: book.MakeBook(1, "title1", "Writer", 2002, 11, 2.2),
			book2: book.MakeBook(1, "title2", "Writer2", 2002, 10, 2.3),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.cmp.Compare(tt.book1, tt.book2), tt.want)
		})
	}
}

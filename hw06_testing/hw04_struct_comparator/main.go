package main

import (
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw06_testing/hw04_struct_comparator/book"
)

type CompareBooks struct {
	CmpType CompareType
}

type CompareType int

const (
	Year CompareType = 0
	Size CompareType = 1
	Rate CompareType = 2
)

func NewCompareBooks(cmpType CompareType) *CompareBooks {
	var cmpBook CompareBooks
	cmpBook.CmpType = cmpType
	return &cmpBook
}

func (cmp CompareBooks) Compare(first book.Book, second book.Book) bool {
	switch cmp.CmpType {
	case Year:
		return first.Year() > second.Year()
	case Size:
		return first.Size() > second.Size()
	case Rate:
		return first.Rate() > second.Rate()
	default:
		return false
	}
}

func main() {
	book1 := book.MakeBook(1, "title Book", "Writer", 2021, 10, 2.2)
	book2 := book.MakeBook(2, "Handbook", "WriterII", 2023, 11, 2.1)
	book1.SetYear(2024)
	cmp := NewCompareBooks(Year)
	fmt.Println(cmp.Compare(book1, book2))
	fmt.Println(NewCompareBooks(Size).Compare(book1, book2))
	fmt.Println(NewCompareBooks(Rate).Compare(book1, book2))
}

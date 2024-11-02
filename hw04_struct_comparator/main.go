package main

import (
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw04_struct_comparator/book"
)

type CompareBooks struct {
	CmpType CompareType
}

type CompareType int

const (
	Year = iota
	Size
	Rate
)

func (cmp CompareBooks) Compare(first book.Book, second book.Book) bool {
	switch cmp.CmpType {
	case 0:
		return first.Year() > second.Year()
	case 1:
		return first.Size() > second.Size()
	case 2:
		return first.Rate() > second.Rate()
	default:
		return false
	}
}

func main() {
	book1 := book.Book{}
	book1.SetYear(2024)
	book2 := book.Book{}
	book2.SetYear(2023)
	book1.SetSize(10)
	book2.SetSize(11)
	book1.SetRate(2.2)
	book2.SetRate(2.1)
	cmp := CompareBooks{CompareType(Year)}
	fmt.Println(cmp.Compare(book1, book2))
	fmt.Println(CompareBooks{CompareType(Size)}.Compare(book1, book2))
	fmt.Println(CompareBooks{CompareType(Rate)}.Compare(book1, book2))
}

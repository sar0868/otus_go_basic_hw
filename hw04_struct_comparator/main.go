package main

import (
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw04_struct_comparator/book"
)

func main() {
	book1 := book.Book{}
	book1.SetYear(2024)
	book2 := book.Book{}
	book2.SetYear(2023)
	book1.SetSize(10)
	book2.SetSize(11)
	book1.SetRate(2.2)
	book2.SetRate(2.1)
	fmt.Println(book1.Compare(book2, book.Year))
	fmt.Println(book1.Compare(book2, book.Size))
	fmt.Println(book1.Compare(book2, book.Rate))
}

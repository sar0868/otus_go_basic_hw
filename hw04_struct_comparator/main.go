package main

import (
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw04_struct_comparator/book"
)

func main() {
	book := book.Book{}
	fmt.Println(book.ID())
	book.SetID(2)
	fmt.Println(book.ID())
}

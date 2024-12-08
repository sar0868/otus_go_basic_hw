package bookgob

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
	Sample []byte
}

func BookGob(books []Book) []Book {
	result := make([]Book, 0)
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	for _, v := range books {
		book := v
		err := enc.Encode(book)
		if err != nil {
			log.Fatalf("encode error: %v\n", err)
		}
	}
	dec := gob.NewDecoder(&network)
	for {
		var book Book
		err := dec.Decode(&book)
		if err != nil {
			break
		}
		result = append(result, book)
	}
	return result
}

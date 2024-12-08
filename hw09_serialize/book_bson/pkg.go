package bookbson

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func WriteBSON(books []Book) ([][]byte, error) {
	result := make([][]byte, 0)
	for _, v := range books {
		book := v
		j, err := bson.Marshal(&book)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil, err
		}
		result = append(result, j)
	}

	return result, nil
}

func ReadBSON(dataSlice [][]byte) []Book {
	result := make([]Book, 0)
	for _, v := range dataSlice {
		var book Book
		err := bson.Unmarshal(v, &book)
		if err != nil {
			continue
		}
		result = append(result, book)
	}
	return result
}

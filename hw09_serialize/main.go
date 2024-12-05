package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
	Sample []byte  `json:"sample"`
}

func WriteJSON(books []Book) ([][]byte, error) {
	// book := Book{
	// 	ID:     1,
	// 	Title:  "title",
	// 	Author: "author",
	// 	Year:   2024,
	// 	Size:   10,
	// 	Rate:   1,
	// 	Sample: []byte{1, 2, 3},
	// }
	result := make([][]byte, 0)
	for _, v := range books {
		book := v
		j, err := json.Marshal(&book)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil, err
		}
		fmt.Printf("%s\n", j)
		result = append(result, j)
	}

	return result, nil
}

func main() {
	book := Book{
		ID:     1,
		Title:  "title",
		Author: "writer",
		Year:   2024,
		Size:   10,
		Rate:   1,
		Sample: []byte{1, 2, 3},
	}
	books := []Book{book, book, book}
	j, err := WriteJSON(books)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", j)
}

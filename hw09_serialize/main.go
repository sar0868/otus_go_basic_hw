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
	result := make([][]byte, 0)
	for _, v := range books {
		book := v
		j, err := json.Marshal(&book)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil, err
		}
		result = append(result, j)
	}

	return result, nil
}

func ReadJSON(jsonSlice [][]byte) []Book {
	result := make([]Book, 0)
	for _, v := range jsonSlice {
		var book Book
		json.Unmarshal(v, &book)
		// if err != nil {
		// 	continue
		// }
		result = append(result, book)
	}
	return result
}

func main() {
	book := Book{
		ID:     1,
		Title:  "title",
		Author: "writer",
		Year:   2024,
		Size:   10,
		Rate:   1,
		Sample: []byte("hello"),
	}
	books := []Book{book, book, book}
	j, err := WriteJSON(books)
	if err != nil {
		fmt.Println(err)
	}
	inp := fmt.Sprintf("%s\n", j)
	// fmt.Printf("%s\n", j)
	fmt.Println(inp)
}

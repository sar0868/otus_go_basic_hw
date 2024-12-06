package bookjson

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

func ReadJSON(dataSlice [][]byte) []Book {
	result := make([]Book, 0)
	for _, v := range dataSlice {
		var book Book
		err := json.Unmarshal(v, &book)
		if err != nil {
			continue
		}
		result = append(result, book)
	}
	return result
}

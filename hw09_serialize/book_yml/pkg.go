package bookyml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Book struct {
	ID     int     `yml:"id"`
	Title  string  `yml:"title"`
	Author string  `yml:"author"`
	Year   int     `yml:"year"`
	Size   int     `yml:"size"`
	Rate   float64 `yml:"rate"`
	Sample []byte  `yml:"sample"`
}

func WriteYML(books []Book) ([][]byte, error) {
	result := make([][]byte, 0)
	for _, v := range books {
		book := v
		j, err := yaml.Marshal(&book)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil, err
		}
		result = append(result, j)
	}

	return result, nil
}

func ReadYML(dataSlice [][]byte) []Book {
	result := make([]Book, 0)
	for _, v := range dataSlice {
		var book Book
		err := yaml.Unmarshal(v, &book)
		if err != nil {
			continue
		}
		result = append(result, book)
	}
	return result
}

package bookxml

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	ID     int     `xml:"id"`
	Title  string  `xml:"title"`
	Author string  `xml:"author"`
	Year   int     `xml:"year"`
	Size   int     `xml:"size"`
	Rate   float64 `xml:"rate"`
	Sample []byte  `xml:"sample"`
}

func WriteXML(books []Book) ([][]byte, error) {
	result := make([][]byte, 0)
	for _, v := range books {
		book := v
		j, err := xml.Marshal(&book)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil, err
		}
		result = append(result, j)
	}

	return result, nil
}

func ReadXML(dataSlice [][]byte) []Book {
	result := make([]Book, 0)
	for _, v := range dataSlice {
		var book Book
		err := xml.Unmarshal(v, &book)
		if err != nil {
			continue
		}
		result = append(result, book)
	}
	return result
}

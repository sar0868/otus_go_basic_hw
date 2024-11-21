package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var text string
	fmt.Scanf("%s", &text)
	fmt.Println(countWords(text))
}

func countWords(str string) map[string]int {
	result := make(map[string]int)
	for _, word := range splitText(str) {
		result[word]++
	}
	return result
}

func splitText(text string) []string {
	text = replacePunctuationMarks(text)
	result := make([]string, 0)
	for _, word := range strings.Split(text, " ") {
		word = strings.Trim(word, " ")
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}

func replacePunctuationMarks(str string) string {
	var result string
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			result += " "
			continue
		}
		result += string(r)
	}
	return result
}

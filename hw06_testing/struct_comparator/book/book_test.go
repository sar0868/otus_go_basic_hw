package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeBook(t *testing.T) {
	name := "Created Book using MakeBook"
	t.Run(name, func(t *testing.T) {
		assert.IsType(t, Book{}, MakeBook(1, "title",
			"author", 2024, 1, 1))
	})
}

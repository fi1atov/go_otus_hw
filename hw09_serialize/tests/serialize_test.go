package main

import (
	"testing"

	// "github.com/fi1atov/go_otus_hw/hw09_serialize"
	"github.com/stretchr/testify/assert"
)

func TestserializationJSON(t *testing.T) {
	t.Parallel()
	bk := Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}

	assert.Equal(t, bk, bk)
}

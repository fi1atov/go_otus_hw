package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestStructJSON(t *testing.T) {
	t.Parallel()
	bk1 := Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bk2 := Book{}

	got := bk1.MarshalJSON()
	want, _ := json.Marshal(&bk1)

	assert.Equal(t, want, got)

	bk2.UnmarshalJSON(got)

	assert.Equal(t, &bk1, &bk2)
}

func TestStructPROTO(t *testing.T) {
	t.Parallel()
	bk1 := Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bk2 := Book{}

	got := bk1.MarshalPROTO()
	want, _ := proto.Marshal(&bk1)

	assert.Equal(t, want, got)

	bk2.UnmarshalPROTO(got)

	assert.Equal(t, bk1.Isbn, bk2.Isbn)
	assert.Equal(t, bk1.Title, bk2.Title)
	assert.Equal(t, bk1.Author, bk2.Author)
	assert.Equal(t, bk1.Year, bk2.Year)
	assert.Equal(t, bk1.Size, bk2.Size)
	assert.Equal(t, bk1.Rate, bk2.Rate)
}

func TestSliceOfStructsJSON(t *testing.T) {
	t.Parallel()
	books := []*Book{
		{
			Isbn:   1,
			Title:  "Книга1",
			Author: "Автор1",
			Year:   1996,
			Size:   145,
			Rate:   0.8,
		},
		{
			Isbn:   2,
			Title:  "Книга2",
			Author: "Автор2",
			Year:   1997,
			Size:   120,
			Rate:   0.6,
		},
	}

	got := serializationJSON(books)
	want, _ := json.Marshal(books)

	assert.Equal(t, want, got)

	booksEmpty := deserializationJSON(got)

	assert.Equal(t, &books, &booksEmpty)
}

func TestSliceOfStructsPROTO(t *testing.T) {
	t.Parallel()
	books := []*Book{
		{
			Isbn:   1,
			Title:  "Книга1",
			Author: "Автор1",
			Year:   1996,
			Size:   145,
			Rate:   0.8,
		},
		{
			Isbn:   2,
			Title:  "Книга2",
			Author: "Автор2",
			Year:   1997,
			Size:   120,
			Rate:   0.6,
		},
	}
	bookList := BooksList{
		Content: books,
	}

	got := SerializationPROTO(books)
	want, _ := proto.Marshal(&bookList)

	assert.Equal(t, want, got)

	booksEmpty := DeserializationPROTO(got)

	assert.Equal(t, books[0].Isbn, booksEmpty[0].Isbn)
	assert.Equal(t, books[0].Title, booksEmpty[0].Title)
	assert.Equal(t, books[0].Author, booksEmpty[0].Author)
	assert.Equal(t, books[0].Year, booksEmpty[0].Year)
	assert.Equal(t, books[0].Size, booksEmpty[0].Size)
	assert.Equal(t, books[0].Rate, booksEmpty[0].Rate)

	assert.Equal(t, books[1].Isbn, booksEmpty[1].Isbn)
	assert.Equal(t, books[1].Title, booksEmpty[1].Title)
	assert.Equal(t, books[1].Author, booksEmpty[1].Author)
	assert.Equal(t, books[1].Year, booksEmpty[1].Year)
	assert.Equal(t, books[1].Size, booksEmpty[1].Size)
	assert.Equal(t, books[1].Rate, booksEmpty[1].Rate)
}

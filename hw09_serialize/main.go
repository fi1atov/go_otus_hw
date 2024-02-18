package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Isbn   int64
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

type Marshaller interface {
	Marshal() []byte
}

type Unmarshaller interface {
	Unmarshal()
}

func (b *Book) Marshal() []byte {
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return res
}

func (b *Book) Unmarshal(data []byte) {
	if err := json.Unmarshal(data, &b); err != nil {
		panic(err)
	}
}

func main() {
	book := Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bookSecond := Book{}

	fmt.Printf("Объект book: %v\n", book)
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)

	res := book.Marshal()
	bookSecond.Unmarshal(res)

	fmt.Printf("book в JSON: %v\n", string(res))
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)
}

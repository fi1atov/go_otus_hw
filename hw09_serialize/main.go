package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// type BookLocal struct {
// 	Isbn   int64
// 	Title  string
// 	Author string
// 	Year   int
// 	Size   int
// 	Rate   float32
// }

type Marshaller interface {
	MarshalJSON() []byte
	MarshalPROTO() []byte
}

type Unmarshaller interface {
	UnmarshalJSON([]byte)
	UnmarshalPROTO([]byte)
}

func (b *Book) MarshalJSON() []byte {
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return res
}

func (b *Book) UnmarshalJSON(data []byte) {
	if err := json.Unmarshal(data, b); err != nil {
		panic(err)
	}
}

func (b *Book) MarshalPROTO() []byte {
	res, err := proto.Marshal(b)
	if err != nil {
		panic(err)
	}
	return res
}

func (b *Book) UnmarshalPROTO(data []byte) {
	if err := proto.Unmarshal(data, b); err != nil {
		panic(err)
	}
}

func SerializationJSON(books []*Book) []byte {
	res, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	return res
}

func DeserializationJSON(data []byte) []*Book {
	books := []*Book{}
	if err := json.Unmarshal(data, &books); err != nil {
		panic(err)
	}
	return books
}

func SerializationPROTO(books []*Book) []byte {
	bookList := BooksList{
		Content: books,
	}
	res, err := proto.Marshal(&bookList)
	if err != nil {
		panic(err)
	}
	return res
}

func DeserializationPROTO(data []byte) []*Book {
	booksTemplate := BooksList{}
	if err := proto.Unmarshal(data, &booksTemplate); err != nil {
		panic(err)
	}
	return booksTemplate.Content
}

func main() {
	book := &Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bookSecond := &Book{}
	bookProtoSecond := &Book{}

	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("-----------one object with JSON--------------\n")

	fmt.Printf("Объект book: %v\n", book)
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)

	res := book.MarshalJSON()
	bookSecond.UnmarshalJSON(res)

	fmt.Printf("book в JSON: %v\n", string(res))
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)

	fmt.Printf("-----------one object with PROTO--------------\n")

	fmt.Printf("Объект bookProto: %v\n", book)
	fmt.Printf("Объект bookProtoSecond: %v\n", bookProtoSecond)

	protoRes := book.MarshalPROTO()
	bookProtoSecond.UnmarshalPROTO(protoRes)

	fmt.Printf("Объект protoRes: %v\n", protoRes)
	fmt.Printf("Объект bookProtoSecond: %v\n", bookProtoSecond)

	fmt.Printf("----------slice objects with JSON-----------\n")
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

	fmt.Printf("Объект books: %v\n", books)

	bookSlice := SerializationJSON(books)
	booksEmpty := DeserializationJSON(bookSlice)

	fmt.Printf("Объект bookSlice: %v\n", string(bookSlice))
	fmt.Printf("Объект booksEmpty: %v\n", booksEmpty)

	fmt.Printf("----------slice objects with PROTO-----------\n")
	fmt.Printf("Объект books: %v\n", books)

	bookPROTO := SerializationPROTO(books)
	deserializedBooks := DeserializationPROTO(bookPROTO)

	fmt.Printf("Объект bookPROTO: %v\n", bookPROTO)
	fmt.Printf("Объект deserializedBooks: %v\n", deserializedBooks)
}

package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
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
	MarshalJSON() []byte
	// MarshalPROTO() []byte
}

type Unmarshaller interface {
	UnmarshalJSON()
	// UnmarshallPROTO()
}

func (b *Book) MarshalJSON() []byte {
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return res
}

func (b *Book) UnmarshalJSON(data []byte) {
	if err := json.Unmarshal(data, &b); err != nil {
		panic(err)
	}
}

// func (b *Book) MarshalPROTO() []byte {
// 	res, err := proto.Marshal(b)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res
// }

// func UnmarshallPROTO([]byte) []Book {
// 	return nil
// }

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

	res := book.MarshalJSON()
	bookSecond.UnmarshalJSON(res)

	fmt.Printf("book в JSON: %v\n", string(res))
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)

	protoRes, _ := proto.Marshal(&book)
	fmt.Printf("Объект protoRes: %v\n", protoRes)
}

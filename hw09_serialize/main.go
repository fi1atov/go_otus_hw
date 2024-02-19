package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type BookLocal struct {
	Isbn   int64
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

type Marshaller interface {
	MarshalJSON() []byte
	MarshalPROTO() []byte
}

type Unmarshaller interface {
	UnmarshalJSON()
	UnmarshalPROTO()
}

func (b *BookLocal) MarshalJSON() []byte {
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return res
}

func (b *BookLocal) UnmarshalJSON(data []byte) {
	if err := json.Unmarshal(data, &b); err != nil {
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

func main() {
	book := BookLocal{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bookSecond := BookLocal{}
	bookProto := &Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bookProtoSecond := &Book{}

	fmt.Printf("Объект book: %v\n", book)
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)

	res := book.MarshalJSON()
	bookSecond.UnmarshalJSON(res)

	fmt.Printf("book в JSON: %v\n", string(res))
	fmt.Printf("Объект bookSecond: %v\n", bookSecond)

	fmt.Printf("-----------------------------------------------\n")

	fmt.Printf("Объект bookProto: %v\n", bookProto)
	fmt.Printf("Объект bookProtoSecond: %v\n", bookProtoSecond)

	protoRes := bookProto.MarshalPROTO()
	bookProtoSecond.UnmarshalPROTO(protoRes)

	fmt.Printf("Объект protoRes: %v\n", protoRes)
	fmt.Printf("Объект bookProtoSecond: %v\n", bookProtoSecond)
}

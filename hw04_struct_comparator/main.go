package main

import "fmt"

const (
	year = "year"
	size = "size"
	rate = "rate"
)

type Book struct {
	id     int64
	title  string
	author string
	year   int
	size   int
	rate   float32
}

// getters

func (b *Book) BookID() int64 {
	return b.id
}

func (b *Book) BookTitle() string {
	return b.title
}

func (b *Book) BookAuthor() string {
	return b.author
}

func (b *Book) BookYear() int {
	return b.year
}

func (b *Book) BookSize() int {
	return b.size
}

func (b *Book) BookRate() float32 {
	return b.rate
}

// setters

func (b *Book) SetBookID(id int64) {
	b.id = id
}

func (b *Book) SetBookTitle(title string) {
	b.title = title
}

func (b *Book) SetBookAuthor(author string) {
	b.author = author
}

func (b *Book) SetBookYear(year int) {
	b.year = year
}

func (b *Book) SetBookSize(size int) {
	b.size = size
}

func (b *Book) SetBookRate(rate float32) {
	b.rate = rate
}

func CompareBooks(b1 *Book, b2 *Book, compareField string) bool {
	switch compareField {
	case year:
		return b1.BookYear() > b2.BookYear()
	case size:
		return b1.BookSize() > b2.BookSize()
	case rate:
		return b1.BookRate() > b2.BookRate()
	default:
		return false
	}
}

func main() {
	bk := Book{}
	bk2 := Book{}

	bk.SetBookID(1)
	bk.SetBookTitle("Книга 1")
	bk.SetBookAuthor("Какой-то автор")
	bk.SetBookYear(1996)
	bk.SetBookSize(56)
	bk.SetBookRate(7.8)

	bk2.SetBookID(1)
	bk2.SetBookTitle("Книга 2")
	bk2.SetBookAuthor("Какой-то автор")
	bk2.SetBookYear(1997)
	bk2.SetBookSize(51)
	bk2.SetBookRate(7.3)

	// fmt.Println(bk.BookID())
	// fmt.Println(bk.BookTitle())
	// fmt.Println(bk.BookAuthor())
	// fmt.Println(bk.BookYear())
	// fmt.Println(bk.BookSize())
	// fmt.Println(bk.BookRate())

	fmt.Println(CompareBooks(&bk, &bk2, year))
	fmt.Println(CompareBooks(&bk, &bk2, size))
	fmt.Println(CompareBooks(&bk, &bk2, rate))
}

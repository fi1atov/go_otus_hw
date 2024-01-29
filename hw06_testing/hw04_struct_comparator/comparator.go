package comparator

const (
	year string = "year"
	size string = "size"
	rate string = "rate"
	none string = "none"
)

type Book struct {
	isbn   int64
	title  string
	author string
	year   int
	size   int
	rate   float32
}

type Comparator struct {
	fieldCompare string
}

// getters

func (b *Book) ISBN() int64 {
	return b.isbn
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

// setters

func (b *Book) SetISBN(isbn int64) {
	b.isbn = isbn
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func NewComparator(fieldCompare string) *Comparator {
	comparator := Comparator{}
	comparator.fieldCompare = fieldCompare
	return &comparator
}

func (c *Comparator) Compare(bookOne, bookTwo *Book) bool {
	switch c.fieldCompare {
	case year:
		return bookOne.Year() > bookTwo.Year()
	case size:
		return bookOne.Size() > bookTwo.Size()
	case rate:
		return bookOne.Rate() > bookTwo.Rate()
	default:
		return false
	}
}

package main

import (
	"testing"

	comparator "github.com/fi1atov/go_otus_hw/hw06_testing/hw04_struct_comparator"
	"github.com/stretchr/testify/assert"
)

const (
	year string = "year"
	size string = "size"
	rate string = "rate"
	none string = "none"
)

func getBooks() (bk1, bk2 comparator.Book) {
	bk1 = comparator.Book{}
	bk2 = comparator.Book{}

	bk1.SetISBN(1)
	bk1.SetTitle("Книга 1")
	bk1.SetAuthor("Автор 1")
	bk1.SetYear(1998)
	bk1.SetSize(56)
	bk1.SetRate(7.8)

	bk2.SetISBN(2)
	bk2.SetTitle("Книга 2")
	bk2.SetAuthor("Автор 2")
	bk2.SetYear(1997)
	bk2.SetSize(51)
	bk2.SetRate(7.3)

	return
}

func TestStructComparator(t *testing.T) {
	t.Parallel()
	bk1, bk2 := getBooks()

	c := comparator.NewComparator(none)
	assert.False(t, c.Compare(&bk1, &bk2))
}

func TestStructComparatorTdt(t *testing.T) {
	t.Parallel()
	bk1, bk2 := getBooks()

	testCases := []struct {
		desc     string
		param    string
		bk1, bk2 comparator.Book
	}{
		{
			desc:  "yearTest",
			param: year,
		},
		{
			desc:  "sizeTest",
			param: size,
		},
		{
			desc:  "rateTest",
			param: rate,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := comparator.NewComparator(tC.param)
			assert.True(t, c.Compare(&bk1, &bk2))
		})
	}
}

package tests

import (
	"reflect"
	"testing"

	jsondata "github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app"
	chessboard "github.com/fi1atov/go_otus_hw/hw06_testing/hw03_chessboard"
	comparator "github.com/fi1atov/go_otus_hw/hw06_testing/hw04_struct_comparator"
	shapes "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes"
	object "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes/objects"
)

func TestFixApp(t *testing.T) {
	res, _ := jsondata.GetJSONData()

	if reflect.TypeOf(res).Kind() != reflect.Slice {
		t.Errorf("Oshibochka")
	}
}

func TestChessboard(t *testing.T) {
	res := chessboard.GetChessBoard(10)

	if reflect.TypeOf(res).Kind() != reflect.String {
		t.Errorf("Oshibochka")
	}
}

func TestStructComparator(t *testing.T) {
	const (
		year string = "year"
		size string = "size"
		rate string = "rate"
		none string = "none"
	)

	bk := comparator.Book{}
	bk2 := comparator.Book{}

	bk.SetISBN(1)
	bk.SetTitle("Книга 1")
	bk.SetAuthor("Автор 1")
	bk.SetYear(1996)
	bk.SetSize(56)
	bk.SetRate(7.8)

	bk2.SetISBN(2)
	bk2.SetTitle("Книга 2")
	bk2.SetAuthor("Автор 2")
	bk2.SetYear(1997)
	bk2.SetSize(51)
	bk2.SetRate(7.3)

	c := comparator.NewComparator(year)
	if c.Compare(&bk, &bk2) == true {
		t.Errorf("Oshibochka")
	}
	c = comparator.NewComparator(size)
	if c.Compare(&bk, &bk2) == false {
		t.Errorf("Oshibochka")
	}
	c = comparator.NewComparator(rate)
	if c.Compare(&bk, &bk2) == false {
		t.Errorf("Oshibochka")
	}
	c = comparator.NewComparator(none)
	if c.Compare(&bk, &bk2) == true {
		t.Errorf("Oshibochka")
	}
}

func TestShapes(t *testing.T) {
	circle := object.Circle{R: 5}
	rectangle := object.Rectangle{X: 10, Y: 5}
	triangle := object.Triangle{X: 8, H: 6}
	hello := "Hello"

	cRes, _ := shapes.CalculateArea(&circle)
	rRes, _ := shapes.CalculateArea(&rectangle)
	tRes, _ := shapes.CalculateArea(&triangle)
	noneRes, _ := shapes.CalculateArea(&hello)

	if cRes != 78.53981633974483 {
		t.Errorf("Oshibochka")
	}
	if rRes != 50 {
		t.Errorf("Oshibochka")
	}
	if tRes != 24 {
		t.Errorf("Oshibochka")
	}
	if noneRes != 0 {
		t.Errorf("Oshibochka")
	}
}

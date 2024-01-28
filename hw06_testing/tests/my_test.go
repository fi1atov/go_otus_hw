package tests

import (
	"reflect"
	"testing"

	jsondata "github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app"
	chessboard "github.com/fi1atov/go_otus_hw/hw06_testing/hw03_chessboard"
	comparator "github.com/fi1atov/go_otus_hw/hw06_testing/hw04_struct_comparator"
	shapes "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes"
	object "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes/objects"
	"github.com/stretchr/testify/assert"
)

func TestFixApp(t *testing.T) {
	res, _ := jsondata.GetJSONData()
	assert.Equal(t, reflect.TypeOf(res).Kind(), reflect.Slice)
}

func TestChessboard(t *testing.T) {
	res := chessboard.GetChessBoard(10)
	assert.Equal(t, reflect.TypeOf(res).Kind(), reflect.String)
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
	assert.False(t, c.Compare(&bk, &bk2))

	c = comparator.NewComparator(size)
	assert.True(t, c.Compare(&bk, &bk2))

	c = comparator.NewComparator(rate)
	assert.True(t, c.Compare(&bk, &bk2))

	c = comparator.NewComparator(none)
	assert.False(t, c.Compare(&bk, &bk2))
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

	assert.Equal(t, cRes, 78.53981633974483)
	assert.Equal(t, rRes, float64(50))
	assert.Equal(t, tRes, float64(24))
	assert.Equal(t, noneRes, float64(0))
}

func TestShapesTdt(t *testing.T) {
	testCases := []struct {
		desc string
		obj  any
		want float64
	}{
		{
			desc: "testCircle",
			obj:  object.Circle{R: 5},
			want: 78.53981633974483,
		},
		{
			desc: "testRectangle",
			obj:  object.Rectangle{X: 10, Y: 5},
			want: 50,
		},
		{
			desc: "testTriangle",
			obj:  object.Triangle{X: 8, H: 6},
			want: 24,
		},
		{
			desc: "testNone",
			obj:  "Hello",
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := shapes.CalculateArea(tC.obj)
			assert.Equal(t, tC.want, got)
		})
	}
}

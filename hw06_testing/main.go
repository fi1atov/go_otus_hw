package main

import (
	"fmt"
	"reflect"

	jsondata "github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app"
	chessboard "github.com/fi1atov/go_otus_hw/hw06_testing/hw03_chessboard"
	comparator "github.com/fi1atov/go_otus_hw/hw06_testing/hw04_struct_comparator"
	shapes "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes"
	object "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes/objects"
)

const (
	year string = "year"
	size string = "size"
	rate string = "rate"
	none string = "none"
)

func main() {
	var size2 int
	fmt.Scanf("%d", &size2)
	res := chessboard.GetChessBoard(size2)
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))

	res2, _ := jsondata.GetJSONData()
	fmt.Println(res2)

	bk := comparator.Book{}
	bk2 := comparator.Book{}

	bk.SetISBN(1)
	bk.SetTitle("Книга 1")
	bk.SetAuthor("Какой-то автор")
	bk.SetYear(1996)
	bk.SetSize(56)
	bk.SetRate(7.8)

	bk2.SetISBN(2)
	bk2.SetTitle("Книга 2")
	bk2.SetAuthor("Какой-то автор")
	bk2.SetYear(1997)
	bk2.SetSize(51)
	bk2.SetRate(7.3)

	fmt.Println(bk.ISBN())
	fmt.Println(bk.Title())
	fmt.Println(bk.Author())
	fmt.Println(bk.Year())
	fmt.Println(bk.Size())
	fmt.Println(bk.Rate())

	c := comparator.NewComparator(year)
	fmt.Println(c.Compare(&bk, &bk2))
	c = comparator.NewComparator(size)
	fmt.Println(c.Compare(&bk, &bk2))
	c = comparator.NewComparator(rate)
	fmt.Println(c.Compare(&bk, &bk2))
	c = comparator.NewComparator(none)
	fmt.Println(c.Compare(&bk, &bk2))

	circle := object.Circle{R: 5}
	rectangle := object.Rectangle{X: 10, Y: 5}
	triangle := object.Triangle{X: 8, H: 6}
	hello := "Hello"

	cRes, _ := shapes.CalculateArea(&circle)
	rRes, _ := shapes.CalculateArea(&rectangle)
	tRes, _ := shapes.CalculateArea(&triangle)
	_, err := shapes.CalculateArea(&hello)

	fmt.Printf("Круг: радиус %d Площадь: %v\n", int(circle.R), cRes)
	fmt.Printf("Прямоугольник: ширина %d, высота %d Площадь: %d\n", int(rectangle.X), int(rectangle.Y), int(rRes))
	fmt.Printf("Треугольник: основание %d, высота %d Площадь: %d\n", int(triangle.X), int(triangle.H), int(tRes))
	fmt.Println(err)
}

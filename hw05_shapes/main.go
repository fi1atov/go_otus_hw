package main

import (
	"errors"
	"fmt"

	"github.com/fi1atov/go_otus_hw/hw05_shapes/objects"
)

type Shape interface {
	Area() float64
}

func calculateArea(s any) (float64, error) {
	if myShape, ok := s.(Shape); ok {
		return myShape.Area(), nil
	}
	return 0, errors.New("ошибка: переданный объект не является фигурой")
}

func main() {
	c := objects.Circle{R: 5}
	r := objects.Rectangle{X: 10, Y: 5}
	t := objects.Triangle{X: 8, H: 6}
	s := "Hello"

	cRes, _ := calculateArea(&c)
	rRes, _ := calculateArea(&r)
	tRes, _ := calculateArea(&t)
	_, err := calculateArea(&s)

	fmt.Printf("Круг: радиус %d Площадь: %v\n", int(c.R), cRes)
	fmt.Printf("Прямоугольник: ширина %d, высота %d Площадь: %d\n", int(r.X), int(r.Y), int(rRes))
	fmt.Printf("Треугольник: основание %d, высота %d Площадь: %d\n", int(t.X), int(t.H), int(tRes))
	fmt.Println(err)
}

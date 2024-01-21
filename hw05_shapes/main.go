package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (r *Rectangle) Area() float64 {
	return r.X * r.Y
}

func (t *Triangle) Area() float64 {
	return t.X * t.H / 2
}

func calculateArea(s any) (float64, error) {
	if myShape, ok := s.(Shape); ok {
		return myShape.Area(), nil
	}
	return 0, errors.New("ошибка: переданный объект не является фигурой")
}

func main() {
	c := Circle{5}
	r := Rectangle{10, 5}
	t := Triangle{8, 6}
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

package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

type Rectangle struct {
	x, y float64
}

type Triangle struct {
	x, h float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	return r.x * r.y
}

func (t *Triangle) area() float64 {
	return t.x * t.h / 2
}

func calculateArea(s any) (float64, error) {
	if myShape, ok := s.(Shape); ok {
		return myShape.area(), nil
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

	fmt.Printf("Круг: радиус %d Площадь: %v\n", int(c.r), cRes)
	fmt.Printf("Прямоугольник: ширина %d, высота %d Площадь: %d\n", int(r.x), int(r.y), int(rRes))
	fmt.Printf("Треугольник: основание %d, высота %d Площадь: %d\n", int(t.x), int(t.h), int(tRes))
	fmt.Println(err)
}

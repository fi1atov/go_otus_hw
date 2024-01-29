package objects

import "math"

type Circle struct {
	R float64
}

type Rectangle struct {
	X, Y float64
}

type Triangle struct {
	X, H float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (r Rectangle) Area() float64 {
	return r.X * r.Y
}

func (t Triangle) Area() float64 {
	return t.X * t.H / 2
}

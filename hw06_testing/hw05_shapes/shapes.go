package shapes

import "errors"

type Shape interface {
	Area() float64
}

func CalculateArea(s any) (float64, error) {
	if myShape, ok := s.(Shape); ok {
		return myShape.Area(), nil
	}
	return 0, errors.New("ошибка: переданный объект не является фигурой")
}

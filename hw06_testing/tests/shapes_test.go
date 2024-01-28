package tests

import (
	"testing"

	shapes "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes"
	object "github.com/fi1atov/go_otus_hw/hw06_testing/hw05_shapes/objects"
	"github.com/stretchr/testify/assert"
)

func TestShapes(t *testing.T) {
	t.Parallel()
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

// Переписал TestShapes с использованием tdt.
func TestShapesTdt(t *testing.T) {
	t.Parallel()
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

package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {

	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("exceeded the number of shaped")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {

	if i >= len(b.shapes) {

		return nil, fmt.Errorf("exceeded")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {

		return nil, fmt.Errorf("exceeded")
	}

	output := b.shapes[i]
	b.shapes = removeIndex(b.shapes, i)
	return output, nil

}

func removeIndex(s []Shape, index int) []Shape {
	return append(s[:index], s[index+1:]...)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	if i >= len(b.shapes) {

		return nil, fmt.Errorf("exceeded")
	}

	output := b.shapes[i]
	b.shapes[i] = shape

	return output, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {

	var output float64 = 0
	for _, b := range b.shapes {

		output += b.CalcPerimeter()
	}

	return output

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var output float64 = 0
	for _, b := range b.shapes {

		output += b.CalcArea()
	}

	return output

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	var output = make([]Shape, 0)
	for _, box := range b.shapes {

		_, ok := box.(Circle)
		if !ok {
			output = append(output, box)
		}

	}

	if len(output) == len(b.shapes) {

		return fmt.Errorf("no circle exist")
	} else {
		b.shapes = output
	}

	return nil

}

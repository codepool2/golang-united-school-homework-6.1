package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (circle Circle) CalcPerimeter() float64 {
	return circle.Radius * 2 * math.Pi
}

func (circle Circle) CalcArea() float64 {

	return math.Pi * math.Pow(circle.Radius, 2)

}

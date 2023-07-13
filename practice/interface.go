package practice

import "fmt"

type Shape interface {
	GetArea() float64
}

type Square struct {
	SideLength float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (s Square) GetArea() float64 {
	area := s.SideLength * s.SideLength
	return area
}

func (t Triangle) GetArea() float64 {
	area := 0.5 * t.Base * t.Height
	return area
}

func PrintArea(s Shape) {
	fmt.Println(s.GetArea())
}

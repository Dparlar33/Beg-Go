package main

import "fmt"

type Shaper interface {
	Area() float64
}

func Calculator(s Shaper) float64 {
	return s.Area()
}

type Rect struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}
type Triangle struct {
	w float64
	h float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (t Triangle) Area() float64 {
	return (t.h * t.w) / 2
}
func main() {
	r := Rect{width: 5.0, height: 4.0}
	c := Circle{radius: 3.0}
	t := Triangle{3, 5}
	fmt.Printf("rect area is %.2f\n\n", Calculator(r))
	fmt.Printf("circle area is %.2f\n\n", Calculator(c))
	fmt.Printf("triangle area is %.2f\n\n", Calculator(t))
}

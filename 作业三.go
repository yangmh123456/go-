package main

import "fmt"
import "math"

type Shape interface {
	area()
}

type Circle struct {
	R float64
	//radius float64
}
type Rectangle struct {
	length, width float64
}

func (m1 Circle) area() {
	aea := math.Pi * math.Pow(m1.R, 2)
	//var mian float64 = 3.14 * m1.R * m1.R
	fmt.Println(aea)
}
func (m2 Rectangle) area() {
	var mian float64 = m2.width * m2.length
	fmt.Println(mian)
}
func Play(s Shape) {
	s.area()
}
func main() {
	var m2 Circle = Circle{R: 12.0}
	var m1 Rectangle = Rectangle{length: 12.0, width: 23.0}
	Play(m1) //这是music1的水声
	Play(m2) //这是music2的水声
}

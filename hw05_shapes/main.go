package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type rectangle struct {
	length, width float64
}

func (r *rectangle) area() float64 {
	return r.length * r.width
}

type triangle struct {
	length, base float64
}

func (t *triangle) area() float64 {
	return t.base * t.length / 2
}

func calculateArea(s any) float64 {
	shape := s.(Shape)

	return shape.area()
}

func main() {
	c := circle{radius: 2}
	a := calculateArea(c)
	fmt.Println(a)
}

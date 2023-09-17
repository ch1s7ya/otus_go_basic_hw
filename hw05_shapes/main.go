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
	shape, ok := s.(Shape)

	if !ok {
		fmt.Println("Object isn't Shape")
	}

	return shape.area()
}

func main() {
	c := circle{radius: 5}
	a := calculateArea(&c)
	fmt.Printf("Круг: радиус %.2f\nПлощадь: %.2f\n", c.radius, a)

	r := rectangle{10, 5}
	a = calculateArea(&r)
	fmt.Printf("Прямоугольник: высота %.2f, ширина: %.2f\nПлощадь: %.2f\n", r.length, r.width, a)

	t := triangle{8, 6}
	a = calculateArea(&t)
	fmt.Printf("Треугольник: высота %.2f, основание: %.2f\nПлощадь: %.2f\n", t.length, t.base, a)
}

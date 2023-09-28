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

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)

	if !ok {
		return 0, fmt.Errorf("object isn't Shape")
	}

	return shape.area(), nil
}

func main() {
	c := circle{radius: 5}
	a, err := calculateArea(&c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Круг: радиус %.2f\nПлощадь: %.2f\n", c.radius, a)

	r := rectangle{10, 5}
	a, err = calculateArea(&r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Прямоугольник: высота %.2f, ширина: %.2f\nПлощадь: %.2f\n", r.length, r.width, a)

	t := triangle{8, 6}
	a, err = calculateArea(&t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Треугольник: высота %.2f, основание: %.2f\nПлощадь: %.2f\n", t.length, t.base, a)

	someObject := struct {
		a float64
		b float64
	}{
		a: 2,
		b: 5,
	}
	_, err = calculateArea(&someObject)
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"math"
)

type DigitCounter interface {
	CountOddEven() (int, int)
}

type Shape interface {
	Area() float64
}

type DigitString string

func (ds DigitString) CountOddEven() (odds int, evens int) {
	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return
}

type Circle struct {
	radius float64
	name   string
}

type Square struct {
	length float64
	name   string
}

type Triangle struct {
	base   float64
	height float64
	name   string
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) String() string {
	return fmt.Sprintf(
		"Area is %v Circumference os %v", c.Area(), c.Circumference())
}

func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}

func calculateArea(listOfShapes []Shape) {
	for _, shape := range listOfShapes {
		fmt.Println("Area of shape is", shape.Area())
	}
}

func main() {
	var d DigitCounter
	s := DigitString("123456789")
	d = s
	c1 := Circle{radius: 5, name: "c1"}
	s1 := Square{length: 6, name: "s1"}
	t1 := Triangle{base: 6, height: 8, name: "t1"}
	var v interface{} = c1

	w, ok := v.(Shape)

	fmt.Println(s.CountOddEven())
	fmt.Println(d.CountOddEven())

	fmt.Println("The area of", c1.name, "is", c1.Area())
	fmt.Println("The area of", s1.name, "is", s1.Area())

	shapes := []Shape{c1, s1, t1}
	calculateArea(shapes)

	fmt.Println(c1)

	if ok {
		fmt.Println(w.Area())
	}
}

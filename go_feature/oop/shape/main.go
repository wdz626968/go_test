package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{
		radius: 5,
	}
	rectangle := Rectangle{
		height: 10.3,
		width:  5.3,
	}

	area := circle.Area()
	fmt.Println("圆的面积是", area)
	perimeter := circle.Perimeter()
	fmt.Println("圆的周长是", perimeter)
	f := rectangle.Area()
	retanglePremeter := rectangle.Perimeter()
	fmt.Println("长方形面积是", f)
	fmt.Println("长方形周长是", retanglePremeter)

}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) Perimeter() float64 {
	return 2 * (c.radius * math.Pi)
}

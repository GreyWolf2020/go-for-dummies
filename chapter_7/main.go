package main

import (
	"fmt"
	"math"
)

type point struct {
	x float32
	y float32
	z float32
}

func main() {
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2
	pt2 := point{x: 5.6, y: 3.8, z: 6.9}
	pt3 := point{
		x: 5.6,
		y: 3.8,
	}
	pt4 := newPoint(7.8, 9.1, 2.3)
	pt5 := pt4
	pt6 := *pt4
	pt7 := pt6
	pt8 := &pt7

	fmt.Println(pt1.x)
	fmt.Println(pt1.y)
	fmt.Println(pt1.z)

	fmt.Println(pt2)
	fmt.Println(pt3)

	fmt.Println(pt4)

	pt4.x = 0
	fmt.Println(pt4)
	fmt.Println(pt5)
	fmt.Println(pt6)
	fmt.Println(pt7)
	fmt.Println(pt8)

	fmt.Println(pt6.length())
	pt6.move(0.1, 0.1, 0.1)
	fmt.Println(pt6)

}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func (p point) length() float64 {
	return math.Sqrt(
		math.Pow(float64(p.x), 2) +
			math.Pow(float64(p.y), 2) +
			math.Pow(float64(p.z), 2))
}

func (p *point) move(deltx, deltay, deltaz float32) {
	p.x += deltx
	p.y += deltay
	p.z += deltaz
}

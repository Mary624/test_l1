package solutions

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Example24() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	fmt.Println(Dist(p1, p2))
}

func Dist(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x)-float64(p2.x), 2) + math.Pow(float64(p1.y)-float64(p2.y), 2))
}

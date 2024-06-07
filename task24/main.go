package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	if count, err := fmt.Scan(&x1, &y1, &x2, &y2); count != 4 || err != nil {
		log.Fatalf("Some errors in input: count %d and err is %s", count, err.Error())
	}
	a := NewPoint(x1, y1)
	b := NewPoint(x2, y2)
	fmt.Print(a.GetDistance(b))

}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetDistance(point *Point) float64 {
	return math.Sqrt(math.Pow(p.x - point.x, 2) + math.Pow(p.y - point.y, 2))
}

type Point struct {
	x, y float64
}
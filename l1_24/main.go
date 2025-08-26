package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct {
	x, y float64
}

func (p *Point) NewPoint(x, y float64) {
	p.x, p.y = x, y
}

func (p *Point) Distance(dest Point) float64 {
	return math.Sqrt(math.Pow(dest.x-p.x, 2) + math.Pow(dest.y-p.y, 2))
}

func (p *Point) PrintPoints() {
	fmt.Println("Точка имеет координаты: ", p.x, p.y)
}

func main() {

	var point1, point2 Point

	point1.NewPoint(rand.Float64()+float64(rand.Intn(100)), rand.Float64()+float64(rand.Intn(100)))
	point2.NewPoint(rand.Float64()+float64(rand.Intn(100)), rand.Float64()+float64(rand.Intn(100)))

	fmt.Print("1: ")
	point1.PrintPoints()
	fmt.Print("2: ")
	point2.PrintPoints()

	fmt.Println("Расстояние между Point1 and Point2 =", point1.Distance(point2))

}

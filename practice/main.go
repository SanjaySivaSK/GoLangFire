package main

import (
	"fmt"
	"math"
)

type shape interface {
	area()
	// perimeter()
}

type circle struct {
	radius int
}

func main() {
	c := circle{2}
	areas(c)
   


}


func (c circle) area() {
	fmt.Println(math.Pi*float64(c.radius)*float64(c.radius))
}

func areas(s shape){
	s.area()
}
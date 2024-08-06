package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	x float32
	y float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (r rectangle) area() float32 {
	return r.x * r.y
}

func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

func LearnInterface() {
	var rec = rectangle{3, 4}
	var c = circle{5}
	var s shape
	s = circle{4}

	fmt.Println(s.area())
	fmt.Println(rec.area())
	fmt.Println(c.area())
}

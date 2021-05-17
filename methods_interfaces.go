package main

import (
	"fmt"
	"math"
)

type vertex2 struct {
	X, Y float64
}

func (v vertex2) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func tour35() {
	v := vertex2{3, 4}
	fmt.Println(v.abs())
}

func scale(v vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func tour36() {
	v := vertex2{3, 4}
	scale(v, 10)
	fmt.Println(v.abs())
}

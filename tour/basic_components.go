package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x, y int) (result int) {
	result = x + y
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func tour1() {
	fmt.Println("Hello, I am Hosokawa.")
	fmt.Println("the time is", time.Now())
	fmt.Println("Random number:", rand.Intn(10))
	fmt.Println("Sqrt number:", math.Sqrt(7))
}

func tour2() {
	var c, python, java bool
	var hskw = "Hosokawa"
	akr := "akira"

	fmt.Println(add(42, 13))
	fmt.Println(swap("Hosokawa", "Akira"))
	fmt.Println(c, python, java)
	fmt.Println(hskw, akr)
}

func tour3() {
	var (
		toBe   bool       = false
		maxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value:%v\n", toBe, toBe)
	fmt.Printf("Type: %T Value:%v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value:%v\n", z, z)

	var i int = 42
	var f float64 = float64(i)
	var s string
	var b bool

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	inferred := i
	fmt.Printf("Type: %T Value:%v\n", inferred, inferred)

	const World = "世界"
	fmt.Println("Ohayo", World)
}

func needInt(x int) int { return x*10 + 1 }
func tour4() {
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		small = big >> 99
	)

	fmt.Println(needInt(small))
	// fmt.Println(needInt(big))
}

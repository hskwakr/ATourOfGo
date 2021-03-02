package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	tour12()
}

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

func tour5() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func tour6() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

func sqrt1(x float64) float64 {
	z := 1.0
	for i := 0; i <= 10; i++ {
		amount := (z*z - x) / (2 * z)
		z -= amount

		if amount > 0 && amount < 1e-16 {
			break
		}
	}
	return z
}

func tour7() {
	fmt.Println("math.Sqrt:", math.Sqrt(2))
	fmt.Println("Sqrt     :", sqrt1(2))
}

func tour8() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

func tour9() {
	fmt.Println("When's Thursday?")

	today := time.Now().Weekday()

	switch time.Thursday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow!")
	default:
		fmt.Println("Too far away!")
	}
}

func tour10() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func tour11() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func tour12() {
	fmt.Println("Start")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("End")
}

package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

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

func tour7() {
	fmt.Println("math.Sqrt:", math.Sqrt(2))
	sqrt, _ := sqrt1(2)
	fmt.Println("Sqrt     :", sqrt)
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

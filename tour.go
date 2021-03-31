package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
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

func tour13() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(i)
}

type vertex struct {
	x int
	y int
}

func tour14() {
	v := vertex{1, 2}
	v.x = 4
	fmt.Println(v.x)
}

var (
	v1 = vertex{1, 2}
	v2 = vertex{x: 1}
	v3 = vertex{}
	vp = &vertex{}
)

func tour15() {
	fmt.Println(v1, v2, v3, vp)
}

func tour16() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
}

func tour17() {
	names := [4]string{
		"J",
		"P",
		"G",
		"R",
	}
	fmt.Println(names)

	a := names[0:2] // J, P
	b := names[1:3] // P, G
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func tour18() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}
	fmt.Println(s)
}

func tour19() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s) // 3, 5, 7

	s = s[:2]
	fmt.Println(s) // 3, 5

	s = s[1:]
	fmt.Println(s) // 5
}

func tour20() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)
	// len=0 cap=6 []

	// Extend its length
	s = s[:4]
	printSlice(s)
	// len=4 cap=6 [2, 3, 5, 7]

	// Drop its first two values
	s = s[2:]
	printSlice(s)
	// len=2 cap=4 [5, 7]

	// panic happen
	//s1 := []int{1, 3, 5}
	//s1 = s1[:5]
	//fmt.Println(s1)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type foo struct {
	int
	string
}

func tour21() {
	// var a int = 0;
	a := 0
	b := "test"
	c := foo{a, b}

	fmt.Println(a, b, c)
}

func tour22() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func tour23() {
	a := make([]int, 5)
	printSlice(a)
	// len=5 cap=5 [0,0,0,0,0]

	b := make([]int, 0, 5)
	printSlice(b)
	// len=0 cap=5 []
}

func tour24() {
	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// X _ X
	// O _ X
	// _ _ O
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func tour25() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)
	// len=1 cap=1 [0]

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)
	// len=2 cap=2 [0, 1]

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
	// len=5 cap=6 [0, 1, 2, 3, 4]
}

func tour26() {
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func tour27() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

func grayscale(dx, dy int) [][]uint8 {
	s1 := make([][]uint8, dy)
	s2 := make([]uint8, dx)

	for i := range s1 {
		s1[i] = s2
	}
	for i, v := range s1 {
		for j := range v {
			s1[i][j] = uint8(i * j)
		}
	}

	return s1
}

func printS(s [][]uint8) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func tour28() {
	pic.Show(grayscale)
}

type vertex1 struct {
	Lat, Long float64
}

var m = map[string]vertex1{
	"Bell":   vertex1{40.1234, -74.1234},
	"Google": {40.1234, -74.1234},
}

func tour29() {
	fmt.Println(m)
}

func tour30() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)

	for i, str := range words {
		count := 1
		for j := range words {
			if i == j {
				continue
			} else if words[i] == words[j] {
				count++
			}
		}
		result[str] = count
	}

	return result
}

func tour31() {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func tour32() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func tour33() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	previous1 := 0
	previous2 := 0
	result := 0

	return func() int {
		if previous1 == 0 && previous2 == 0 {
			previous2 = 1
			result = 0
		} else if previous1 == 0 && previous2 == 1 {
			previous1 = 1
			result = 1
		} else if previous1 == 1 && previous2 == 1 {
			previous2 = 2
		} else {
			result = previous1 + previous2
			previous1 = previous2
			previous2 = result
		}

		return result
	}
}

func tour34() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

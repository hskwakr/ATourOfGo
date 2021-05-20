package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

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
	x := 0
	y := 1
	result := 0

	return func() int {
		result = x
		x, y = y, x+y

		return result
	}
}

func tour34() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

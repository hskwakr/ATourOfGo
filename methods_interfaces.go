package main

import (
	"fmt"
	"math"
	"time"
)

type vertex2 struct {
	X, Y float64
}

func (v vertex2) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func absFunc(v vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func tour35() {
	v := vertex2{3, 4}
	fmt.Println(v.abs())
}

func (v *vertex2) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func tour36() {
	v := vertex2{3, 4}
	v.scale(2)
	scaleFunc(&v, 10)

	p := &vertex2{4, 3}
	p.scale(3)
	scaleFunc(p, 8)

	fmt.Println(v, p)
}

func tour37() {
	v := vertex2{3, 4}
	fmt.Println(v.abs())
	fmt.Println(absFunc(v))

	p := &vertex2{4, 3}
	fmt.Println(p.abs())
	fmt.Println(absFunc(*p))
}

type abser interface {
	abs() float64
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func tour38() {
	var a abser
	f := myFloat(-math.Sqrt2)
	v := vertex2{3, 4}

	a = f
	fmt.Println(a.abs())

	a = v
	fmt.Println(a.abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func tour39() {
	var i I = T{"Hello"}
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func tour40() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func tour41() {
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "hello"
	describe1(i)
}

func tour42() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func tour43() {
	do(21)
	do("hello")
	do(true)
}

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func tour44() {
	a := person{"Arthur", 42}
	z := person{"Zaphod", 9001}
	fmt.Println(a, z)
}

type ipAddr [4]byte

func (ip ipAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func tour45() {
	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type myError struct {
	when time.Time
	what string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &myError{
		time.Now(),
		"it didn't work",
	}
}

func tour46() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	v := fmt.Sprint(float64(*e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", v)
}

func sqrt1(x float64) (float64, error) {
	z := 1.0
	if x < 0.0 {
		e := ErrNegativeSqrt(x)
		return x, &e
	}

	for i := 0; i <= 10; i++ {
		amount := (z*z - x) / (2 * z)
		z -= amount

		if amount > 0 && amount < 1e-16 {
			break
		}
	}
	return z, nil
}

func tour47() {
	fmt.Println(sqrt1(2))
	fmt.Println(sqrt1(-2))
}

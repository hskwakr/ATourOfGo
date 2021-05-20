package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func tour52() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func tour53() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func tour54() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func tour55() {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func tour56() {
	c := make(chan int, 20)
	quit := make(chan int)
	go func() {
		for i := 0; i < cap(c); i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func tour57() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		ch <- t.Value
		walker(t.Left)
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	r := true
	sumX, sumY := 0, 0
	for {
		x, ok1 := <-c1
		y, ok2 := <-c2

		sumX += x
		sumY += y

		if ok1 != ok2 {
			r = false
			break
		}
		if !ok1 && !ok2 {
			break
		}
	}

	if sumX != sumY {
		r = false
	}
	return r
}

func tour58() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(3)))
}

type safeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *safeCounter) inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *safeCounter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func tour59() {
	c := safeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//s := []int{7, 2, 8, -9, 4, 0}

	// go say("world")
	// say("hello")
}

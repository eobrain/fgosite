package main

import "fmt"

func sum(a, c) {
	// send sum to c
	c <- +(...a)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	{
		x, y := <-c, <-c // receive from c

		fmt.Println(x, y, x+y)
	}
}

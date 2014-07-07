package main

import "fmt"

func sum(a, c) {
	// send sum to c
	c <- +(...a)
}

func main() {
	a    := [7, 2, 8, -9, 4, 0]
	c    := make(chan)
	half := count(a)/2

	go sum(half  drop  a, c)
	go sum(half  take  a, c)
	{
		// receive from c
		[x, y] := [<-c, <-c]
		fmt.Println(x, y, x + y)
	}
}

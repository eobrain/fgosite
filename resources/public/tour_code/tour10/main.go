package main

import "fmt"

func split(sum int) { //(x, y int) {
	const(
		x = int(sum * 4 / 9)
		y = sum - x
	)
	[x, y]
}

func main() {
    fmt.Println(split(17))
}

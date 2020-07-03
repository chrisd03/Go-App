package main

import "fmt"

func main() {
	x := plusTwo()
	fmt.Printf("%d\n", x(2))
}

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

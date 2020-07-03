package main

import "fmt"

func main() {
	y := plusX(5)
	fmt.Printf("%d\n", y(2))
}

func plusX(y int) func(int) int {
	return func(x int) int {
		return y + x
	}
}

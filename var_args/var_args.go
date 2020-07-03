package main

import "fmt"

func main() {
	ints(1, 4, 67)
}

func ints(arg ...int) {
	for _, x := range arg {
		fmt.Printf("%d\n", x)
	}
}

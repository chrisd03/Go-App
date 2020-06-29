package main

import "fmt"

func main() {
	var x [10]int
	for i := 0; i < 10; i++ {
		x[i]++
	}
	for _, v := range x {
		fmt.Printf("%d\n", v)
	}
}

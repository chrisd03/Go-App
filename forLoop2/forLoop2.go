package main

import "fmt"

func main() {
	var x [10]int
	for i := 0; i < 10; i++ {
		x[i]++
	}
	for _, v := range x {
		print(v)
	}
}

func print(y int) {
	fmt.Printf("%d\n", y)
}

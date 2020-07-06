package main

import "fmt"

func main() {
	list := []int{1, 6, 3, 8, 1}
	newFunc := func(x int) int {
		return x * 10
	}
	fmt.Printf("%v", (mapTest(newFunc, list)))
}

func mapTest(f func(int) int, list []int) []int {
	newList := make([]int, len(list))
	for x, y := range list {
		newList[x] = f(y)
	}
	return newList
}

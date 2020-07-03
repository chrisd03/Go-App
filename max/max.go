package main

import "fmt"

func main() {
	x := []int{1, 5, 8, 3, 6, 9, 14, 11, 21, 65, 3, 2}
	fmt.Print(max(x))
}

func max(x []int) (max int) {
	max = 0
	for i, _ := range x {
		for k := i + 1; k < len(x); k++ {
			if x[k] > x[i] {
				max = x[k]
			}
		}
	}
	return
}

package main

import "fmt"

func main() {
	arr := [...]float64{1.34, 1.54, 1.76}
	slice := arr[:]
	fmt.Print(mean(slice))
}

func mean(x []float64) (average float64) {
	avg := 0.0
	for _, i := range x {
		avg += i
	}
	average = avg / float64(len(x))
	return
}

package main
import "fmt"

func main() {
	arr := [...]float64{1.532, 2.76, 3.35, 4.92, 5.32}
	slice := arr[:]
	average := 0.0
	for _, x := range slice{
		average += x
	}
	fmt.Print(average / 5)
}
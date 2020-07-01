package main

import "fmt"

func main() {
	x := []int{4, 65, 10, 24, 16, 78, 66, 34, 105, -43, -786, 1086}
	fmt.Print("Unsorted: \n", x)

	sort(x)
	fmt.Print("\nSorted: \n", x)

}

func sort(i []int) {
	for j, _ := range i {
		for k := j + 1; k < len(i); k++ {
			if i[k] < i[j] {
				i[j], i[k] = i[k], i[j]
			}
		}
	}
}

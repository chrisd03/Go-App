package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s stack
	s.push(25)
	s.push(14)
	s.push(64)
	s.push(75)
	s.push(65)
	s.push(2564)
	s.push(675)
	s.push(7354)
	s.push(252)
	fmt.Printf("%v", s.String())
}
func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}
func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

type stack struct {
	i    int
	data [10]int
}

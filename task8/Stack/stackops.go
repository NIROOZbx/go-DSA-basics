package main

import "fmt"

type Stack struct {
	items []int
}

var top int

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
	top++

}

func (s *Stack) Pop() []int {

	top--

	return s.items[:top]
}

func (s *Stack) Peek() int {
	return s.items[top-1]
}

func main() {

	var st Stack

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)

	fmt.Println(st.Pop())
	fmt.Println(st.Peek())

}

package main

import "fmt"

type Stack struct {
	items []int
}

var top = -1

func (s *Stack) Push(val int) {
	top++
	s.items = append(s.items, val)

}

func (s *Stack) Pop() []int {

	top--

	return s.items[:top+1]
}

func (s *Stack) Peek() int {
	return s.items[top]
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

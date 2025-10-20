package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(val int) {
	newNode := &Node{Data: val, Next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() {
	s.top = s.top.Next
	s.size--

}

func (s *Stack) Peek() int{
	return s.top.Data
}

func main() {

	var st Stack

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	st.Push(50)
	st.Push(60)
	st.Pop()

	fmt.Println(st.top)
	temp := st.top
	for temp != nil {

		fmt.Println(temp.Data)
		temp = temp.Next
	}

	fmt.Println("Top elemeent",st.Peek())

}

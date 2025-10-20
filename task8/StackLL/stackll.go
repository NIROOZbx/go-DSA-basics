package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	top *Node
	size int
}

func (s *Stack) Push(val int) {
	newNode := &Node{Data: val, Next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() {
	s.top=s.top.Next
	s.size--


}

func main() {

	var st Stack

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Pop()

	temp:=st.top
  for temp!=nil{

	fmt.Println(temp.Data)
	temp=temp.Next
  }



	

	

}

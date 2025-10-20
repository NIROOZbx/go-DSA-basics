package main

import "fmt"

type Stack struct {
	items []int
	top   int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
	s.top++
}

func (s *Stack) Pop() {
	s.top--
	s.items = s.items[:s.top]
}

func main() {

	var data int
	var input string
	var st Stack

	for {

		fmt.Println("Choose the operations")

		fmt.Println("Push")
		fmt.Println("Pop")
		fmt.Println("View")
		fmt.Println("Exit")

		fmt.Scanln(&input)

		if input == "Push" {
			fmt.Println("Enter the value to add to stack")
			fmt.Scanln(&data)
			st.Push(data)
		} else if input == "Pop" {
			st.Pop()
		} else if input == "View" {
			fmt.Println(st.items)
		} else {
			return
		}

	}

	// st.Push(10)
	// st.Push(20)
	// st.Push(30)

	// st.Pop()

}

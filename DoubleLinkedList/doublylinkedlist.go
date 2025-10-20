package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Back *Node
}

func main() {

	arr := []int{1, 2, 3, 4, 5}

	head := &Node{Data: arr[0]}

	current := head

	for i := 1; i < len(arr); i++ {
		newNode := &Node{Data: arr[i], Back: current}
		current.Next = newNode
		current = newNode
	}

	temp := head

	for temp != nil {
		fmt.Println("Temp data",temp.Data,"Temp Back",temp.Back,"Temp next",temp.Next)
		temp=temp.Next
	}


}
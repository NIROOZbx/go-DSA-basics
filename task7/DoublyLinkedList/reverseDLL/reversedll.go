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

		current.Next = &Node{Data: arr[i], Back: current}

		current = current.Next
	}

	head = reverse(head)

	temp := head

	for temp != nil {
		fmt.Println(temp.Data)
		temp=temp.Next
	}

}

func reverse(head *Node) *Node {

	temp := head

	var prev *Node

	for temp != nil {

		prev=temp.Back
		temp.Back=temp.Next
		temp.Next= prev

		temp = temp.Back
		
	}
	

	return prev.Back

}
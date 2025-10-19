package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {

	node1 := &Node{Data: 94}
	node2 := &Node{Data: 45}
	node3 := &Node{Data: 3}
	node4 := &Node{Data: 1}

	node1.Next=node2
	node2.Next=node3
	node3.Next=node4


	current:=node1

	var arr []int

	for current!=nil{
		arr=append(arr, current.Data)
		fmt.Println(current.Data,current.Next)
		current=current.Next

	}

fmt.Println(arr)
}

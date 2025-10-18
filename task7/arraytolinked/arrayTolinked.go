package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {

	var arr = []int{3, 4, 5, 6, 7}

	head := &Node{Data: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {

		newNode := &Node{Data: arr[i]}
		current.Next = newNode
		fmt.Println(newNode.Data, newNode.Next)

	}

}

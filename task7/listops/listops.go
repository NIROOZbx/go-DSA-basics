package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

//head deletion

// func DeleteHead(h *Node) *Node {

// 	newHead := h.Next

// 	return newHead

// }

func main() {

	var arr = []int{5, 6, 7, 8, 9}

	head := &Node{Data: arr[0]}

	listLen := 0

	current := head

	for i := 1; i < len(arr); i++ {

		newNode := &Node{Data: arr[i]}

		current.Next = newNode

		current = newNode

	}

	// head = DeleteHead(head)

	// iterator := head
	// for iterator != nil {

	// 	fmt.Println(iterator.Data)

	// 	iterator = iterator.Next
	// }

	//tail deletion

	// temp := head

	// for temp.Next.Next != nil {
	// 	temp = temp.Next

	// 	fmt.Println(temp)
	// }
	// temp.Next = nil

	newHead := insertNode(head, 100)

	newTemp := newHead

	for newTemp.Next != nil {
		newTemp = newTemp.Next
	}

	newTemp.Next = insertTail(50)

	
	
	k := 4

	temp := newHead

	var prev *Node

	for temp != nil {

		listLen++

		if k == 1 {
			newHead = newHead.Next
			break
		}

		if listLen == k {

			prev.Next = prev.Next.Next
			break
		}
		prev = temp
		temp = temp.Next

	}

	x:=3
	cn:=0

	newCurrent:=newHead

	
	for newCurrent!=nil{
		cn++


		if cn==x-1{
			newCurrent.Next=insertAtK(newCurrent.Next, 40)
			break
		}

		newCurrent=newCurrent.Next

	}

	iterator := newHead

	for iterator != nil {

		fmt.Println(iterator.Data)

		iterator = iterator.Next
	}

}

func insertNode(h *Node, val int) *Node {
	return &Node{Data: val, Next: h}
}

func insertTail( val int) *Node {
	return &Node{Data: val}
}
func insertAtK(n *Node,val int)*Node{
	return &Node{Data:val,Next: n }
}
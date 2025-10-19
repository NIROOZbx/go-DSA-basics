package main

import (
	"fmt"

)

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

	head = insertFront(30, head)
	head = insertFront(40, head)
	head = insertFront(50, head)

	head = DeleteTail(head)

	head = DeleteHead(head)

	head = DeleteK(head, 2)

	head=insertBack(head,6)

	head=insertAtK(head, 3,7)

	temp := head

	for temp != nil {
		fmt.Println("Current data", temp.Data, "Back node", temp.Back)
		temp = temp.Next
	}
}

func insertFront(val int, head *Node) *Node {
	newNode := &Node{Data: val, Next: head}

	if head != nil {
		head.Back = newNode
	}

	return newNode
}
func insertBack(head *Node ,val int)*Node{


	temp:=head

	for temp.Next!=nil{
		temp=temp.Next
	}



	temp.Next=&Node{Data: val,Back: temp}


	return  head

}

func DeleteHead(head *Node) *Node {

	if head == nil {
		return nil
	}

	prev := head
	head = head.Next

	head.Back = nil
	prev.Next = nil

	return head

}

func DeleteTail(head *Node) *Node {

	temp := head

	for temp.Next != nil {

		temp = temp.Next

	}
	temp.Back.Next = nil
	temp.Back = nil

	return head
}

func DeleteK(head *Node, k int) *Node {

	var counter = 0

	temp := head

	for temp != nil {

		counter++

		if counter == k {

			temp.Back.Next = temp.Next
			temp.Next.Back = temp.Back

			temp.Next=nil
			temp.Back=nil
			break
		}

		temp = temp.Next
	}

	return head

}

func insertAtK(head *Node , k int,val int) *Node{

	var counter=0

	temp:=head

	for temp!=nil{
		counter++

		if counter==(k-1){
			newNode:=&Node{Data: val,Next:temp.Next,Back:temp}
			temp.Next.Back=newNode
			temp.Next=newNode
			
			break
		}

		temp=temp.Next
	}


	return head

}

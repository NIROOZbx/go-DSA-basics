package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	len := 0

	temp := head

	res := 0

	for temp != nil {
		len++
		temp = temp.Next
	}

	res = len / 2

	fmt.Println(res)

	temp = head
	for i := 0; i < res; i++ {
		temp = temp.Next
	}
	return temp

}

func main() {

	arr := []int{1,2,3,4,5}

	head := &ListNode{Val: arr[0]}

	current := head

	for i := 1; i < len(arr); i++ {

		current.Next = &ListNode{Val: arr[i]}

		current = current.Next
	}

	fmt.Println(middleNode(head))
}

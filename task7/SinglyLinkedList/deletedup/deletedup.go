package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	current := head

	for current != nil && current.Next != nil {

		if current.Next.Val == current.Val {

			current.Next = current.Next.Next

		} else {

			current = current.Next
		}

	}

	return head

}

func main() {

	nums := []int{1, 1, 1}

	head := &ListNode{Val: nums[0]}

	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	data := deleteDuplicates(head)

	for data != nil {

		fmt.Println(data.Val)
		data = data.Next
	}

}

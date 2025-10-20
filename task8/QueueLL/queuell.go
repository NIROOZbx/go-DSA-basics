package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Start       *Node
	End         *Node
	currentSize int
}

func (q *Queue) Enqueue(val int) {

	newNode := &Node{Data: val}

	if q.currentSize == 0 {
		q.Start = newNode
		q.End = newNode
	}else{
		q.End.Next=newNode
		q.End=newNode

	}

	q.currentSize++

}

func (q *Queue) Dequeue(){

	q.Start = q.Start.Next
	q.currentSize--
}

func main() {

	var q Queue

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Dequeue()

	temp := q.Start

	
	for temp != nil {

		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

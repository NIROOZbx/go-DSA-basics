package main

import "fmt"

type Queue struct {
	items       [5]int
	currentSize int
	Start       int
	End         int
}

func (q *Queue) Enqueue(val int) {

	if q.currentSize == len(q.items) {
		fmt.Println("Queue limit exceeded")
		return
	}


	q.items[q.End] = val
	q.End = (q.End + 1) % len(q.items)
	q.currentSize++

}

func (q *Queue) Dequeue() {

	if q.currentSize == 0 {
		return 
	}

	q.Start=(q.Start+1)%len(q.items)
	q.currentSize--
	
}

func (q *Queue) Peek() int {
	return q.items[q.Start]
}

func main() {
	var Q Queue

	Q.Enqueue(10)
	Q.Enqueue(20)
	Q.Enqueue(30)
	Q.Enqueue(40)
	Q.Enqueue(50)
	Q.Dequeue()
	Q.Enqueue(60)
	

	fmt.Println("Item at start position", Q.Peek())
	fmt.Println("Item at end position", Q.items[Q.End])
	fmt.Println(Q.items)
}

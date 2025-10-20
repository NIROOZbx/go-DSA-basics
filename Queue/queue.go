package main

import "fmt"

type Queue struct {
	items [5]int
	size  int
	start int
	end   int
}

func (q *Queue) Enqueue(val int) {

	if len(q.items)==q.size{
		fmt.Println("Queue limit exceeded")
		return
	}

	q.items[q.end]=val
	q.end=(q.end+1)%len(q.items)
	q.size++
}


func (q *Queue) Dequeue(){
	
	q.start=(q.start+1)%len(q.items)
	q.size--
}
func main() {

	var q Queue

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Dequeue()
	q.Enqueue(60)
	q.Enqueue(70)

	fmt.Println(q.items)

}
package main
type MyQueue struct {
     items []int
    size  int
    
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {

    this.items=append(this.items,x)
    this.size++
    
}


func (this *MyQueue) Pop() int {
    this.size--

    p := this.items[0]
    this.items = this.items[1:]
    return p
    
}


func (this *MyQueue) Peek() int {
    return this.items[0]
    
}


func (this *MyQueue) Empty() bool {
    return this.size==0
    
}

func main(){
	
}
/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
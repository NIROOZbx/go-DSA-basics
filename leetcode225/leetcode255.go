package main
type MyStack struct {

    items []int

    
}


func Constructor() MyStack {
    return MyStack{}
}


func (this *MyStack) Push(x int)  {

    this.items=append(this.items,x)
    for i := 0; i < len(this.items)-1; i++ {
        this.items = append(this.items, this.items[0])
        this.items = this.items[1:]
    }


}


func (this *MyStack) Pop() int {
    val := this.items[0]
    this.items = this.items[1:]
    return val
  
    
}


func (this *MyStack) Top() int {

    return this.items[0]
    
}


func (this *MyStack) Empty() bool {

   return len(this.items) == 0
    
}


func main(){
	
}
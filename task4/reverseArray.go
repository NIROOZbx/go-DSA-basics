package main

import "fmt"

func main() {
	
	var arr = []int{1,2,3,4,5,6,7,8,9,10,11}
	
	// for i := len(arr) - 1; i >=0; i-- {
		// 	fmt.Print(arr[i])
		// }
		
		temp:=len(arr)-1
		i:=0
	for i<temp{
		arr[i],arr[temp]=arr[temp],arr[i]
		i++
		temp--
	}
	fmt.Print(arr)
}
package main

import (
	"fmt"
)

func main() {
	arr := []int{}

	arr = append(arr, 32, 54, 6, 7,9) //inserting at end of the array
	index := 2
	searchEle:=9
	deleteIndex:=3
	
	 arr=append(arr, 0)
     fmt.Println(arr)
	 copy(arr[index+1:],arr[index:])//inserting at a custom position

	 arr[index]=56

	 arr=append(arr[:deleteIndex],arr[deleteIndex+1:]...)


	for i:=0;i<len(arr);i++{
		if arr[i]==searchEle{
			fmt.Println("Found element",arr[i],"At the index:",i)//searching element
		}
	}
	fmt.Println(arr)

}
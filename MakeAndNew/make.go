package main

import "fmt"

func main() {
	newSlice := make([]*int, 0, 5)
	for i := 0; i < 4; i++ {
		newSlice = append(newSlice, &i)
	}
	for _, val := range newSlice {
		fmt.Println(*val)
	}
	x := new([]int)
	x1:=new(int)
	fmt.Println(*x1)
	*x = append(*x, 55, 3)
	fmt.Println(*x)

}

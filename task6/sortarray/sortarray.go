package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 0
	newMap := map[int]bool{}


	for _,val:=range nums {
		fmt.Print(val)

		if !newMap[val]{
			newMap[val]=true
			nums[k]=val
			k++
		}

	}
	fmt.Println(nums)
	fmt.Println(newMap)
	return k

}

func main() {
	// fmt.Println(removeDuplicates( []int{1,1,2}))
	fmt.Println(removeDuplicates( []int{0,0,1,1,1,2,2,3,3,4}))
}
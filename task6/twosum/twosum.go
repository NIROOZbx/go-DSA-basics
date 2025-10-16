package main

import "fmt"

func twoSum(nums []int, target int) []int{

	lowNum := []int{}
	newMap:=map[int]int{}
	for i := 0; i < len(nums); i++ {

		_,ok:=newMap[target-nums[i]]

		if ok{
			lowNum=append(lowNum, newMap[target-nums[i]],i)
		}else{
			newMap[nums[i]] = i


		}
	}
	return lowNum
	
}

func main() {
	fmt.Println(twoSum([]int{3,2,4},6))
	fmt.Println(twoSum([]int{2,7,11,15},9))
	fmt.Println(twoSum([]int{3,3},6))

}
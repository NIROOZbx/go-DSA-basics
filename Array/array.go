package main

import "fmt"

func findMax() int {
	var nums = [7]int{3, 2, 9, 24, 1, 4, 5}

	max := nums[0]

	for i := 0; i < len(nums); i++ {

		if max < nums[i] {
			max = nums[i]
		}
	}

	return max


func main() {

	fmt.Println(findMax())

}

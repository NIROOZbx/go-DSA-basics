package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	copy(nums1[m:], nums2[:n])
	left := 0
	right := 1

	for left < len(nums1)-1 {

		if nums1[left] > nums1[right] {
			nums1[left], nums1[right] = nums1[right], nums1[left]

		}
		right++

		if right > len(nums1)-1 {
			left++
			right = left
		}

	}

}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

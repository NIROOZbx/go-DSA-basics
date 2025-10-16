package main
import "fmt"


func merge(nums1 []int, m int, nums2 []int, n int)  {

	copy(nums1[m:], nums2[:n])
	fmt.Println(nums1)

	for i:=0;i<len(nums1);i++{
		for j:=i+1;j<len(nums1);j++{
			if nums1[j]<nums1[i]{
				nums1[i],nums1[j]=nums1[j],nums1[i]
			}
		}
	}
	
    
}

func main(){
 merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6},3)
}
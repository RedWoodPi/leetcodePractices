package __Two_Sum

import "fmt"

func main()  {
	nums := []int{1,3,5,6,8}
	tar := 9
	fmt.Println(twoSum(nums, tar))
}
func twoSum(nums []int, target int)  []int {
	re := []int{}
	for i:=0;i<len(nums);i++ {
		for j:=i+1; j<len(nums); j++ {
			if nums[i]+nums[j] == target {
				re = append(re, i)
				re = append(re, j)
			}
		}
	}
	return  re
}

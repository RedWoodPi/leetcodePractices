package main

import (
	"fmt"
)

func main()  {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(solution(nums))
}

func solution(nums []int) int{
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i:=1;i<len(nums);i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j+1
}
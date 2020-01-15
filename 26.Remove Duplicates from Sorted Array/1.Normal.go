package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	for i:=0;i<len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[0:i],nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

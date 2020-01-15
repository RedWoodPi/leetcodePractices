package main

import "fmt"

func main()  {
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(search1(nums,0))
}

func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := right/2

	for left <= right {
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = left + (right - left)/2
	}

	return -1
}

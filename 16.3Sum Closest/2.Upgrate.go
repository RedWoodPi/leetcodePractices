package main

import "fmt"

func main()  {
	arr := []int{-1,2,1,-4}
	fmt.Println(threeSumClosest2(arr,1))
}

func threeSumClosest2(nums []int, target int) int {
	nums = shellSort(nums)
	closest := nums[0]+nums[1]+nums[len(nums)-1]
	for i:=0;i<len(nums);i++ {
		left := i+1
		right := len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(target - sum) < abs(target - closest) {
				closest = sum
			}
			if sum > target {
				right--
			} else if target > sum {
				left++
			} else {
				return closest
			}
		}
	}

	return closest
}

func abs(num int)  int{
	if num < 0 {
		return -num
	} else {
		return num
	}
}

//希尔排序
func shellSort(nums []int) []int {
	gap := len(nums)/2
	for gap > 0 {
		for i:=gap;i<len(nums);i++ {
			temp := nums[i]
			index := i - gap
			for index >= 0 && nums[index] > temp {
				nums[index+gap] = nums[index]
				index = index-gap
			}
			nums[index+gap] = temp
		}
		gap = gap/2
	}
	return nums
}
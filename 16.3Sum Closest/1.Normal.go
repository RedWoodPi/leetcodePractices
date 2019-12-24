//package _6__3Sum_Closest
package main

import "fmt"

func main()  {
	arr := []int{0,2,1,-3}
	fmt.Println(threeSumClosest1(arr,1))
}

func threeSumClosest1(nums []int, target int) int {
	nums = ShellSort(nums)
	closest := Abs(nums[0] + nums[1] + nums[2] - target)
	closestIndex := []int{0,1,2}
	for i:=1;i<len(nums);i++ {
		left := i-1
		for left >=0 {
			right := i+1
			for right < len(nums) {
				if Abs(nums[left]+nums[i]+nums[right] -target) < closest {
					closest = Abs(nums[left]+nums[i]+nums[right] -target)
					closestIndex = []int{left,i,right}
				}
				right++
			}
			left--
		}

	}
	return nums[closestIndex[0]]+nums[closestIndex[1]]+nums[closestIndex[2]]
}

func Abs(num int)  int{
	if num < 0 {
		return -num
	} else {
		return num
	}
}

//希尔排序
func ShellSort(nums []int) []int {
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
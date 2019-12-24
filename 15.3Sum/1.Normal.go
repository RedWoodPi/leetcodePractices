package leetcodePractices
//package main

import (
	"fmt"
)

func main()  {
	arr := []int{-1,0,1,0}
	fmt.Println(BubulleSort([]int{-1,2,1,-4}))
	fmt.Println(threeSum(arr))
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	//排序，算法可选
	nums = BubulleSort(nums)
	for i:=1;i<len(nums);i++ {
		//以nums[i]为中心向两边移动指针
		left := i-1
		right := i+1
		for left >= 0 && right < len(nums) {
			if nums[left] + nums[i] + nums[right] > 0 {
				left--
			} else if nums[left] + nums[i] + nums[right] < 0 {
				right++
			} else {
				hasElm := false
				//去重
				for j:=0;j<len(result);j++ {
					if result[j][0] == nums[left] && result[j][1] == nums[i]{
						hasElm = true
					}
				}
				if !hasElm {
					result = append(result,[]int{nums[left],nums[i],nums[right]})
				}

				left--
				right++
			}
		}
	}
	return result
}

//冒泡排序
func BubulleSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	temp := 0
	for i:=0;i<len(nums);i++ {
		for j:=1;j<len(nums)-i;j++ {
			if nums[j] < nums[j-1] {
				temp = nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}

	}
	return nums
}

//选择排序
func SelectionSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	arr := []int{}
	var temp int
	var tempKey int = 0

	for len(nums) != 0 {
		temp = nums[0]
		tempKey = 0
		for i:=1;i<len(nums);i++ {
			if nums[i] < temp {
				temp = nums[i]
				tempKey = i
			}
		}
		arr = append(arr,temp)
		if tempKey == len(nums) - 1 {
			nums = nums[:tempKey]
		} else {
			nums = append(nums[:tempKey],nums[(tempKey+1):]...)
		}

	}
	return arr
}

//插入排序
func InsertionSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i:=0;i<len(nums)-1;i++ {
		for j:=i+1;j>0;j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}
	return nums
}

//希尔排序
func shellSort(nums []int) []int {
	fmt.Println(nums)
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

//归并排序
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums)/2
	left := nums[0:mid]
	right := nums[mid:]

	return merge(mergeSort(left),mergeSort(right))
}
func merge(left []int,right []int) []int {
	lenth := len(left) + len(right)
	result := []int{}
	i,j:=0,0;
	for index:=0;index < lenth;index++ {
		if i >= len(left) {
			result = append(result, right[j])
			j++
		} else if j >= len(right) {
			result = append(result, left[i])
			i++
		} else  if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	return result
}

//快速排序
func quickSort(nums []int,start int, end int) []int {
	if start >= end {
		return nil
	}
	mark := start
	pivot := nums[start]
	pointer := start+1

	for pointer < len(nums) {
		if nums[pointer] < pivot{
			mark++
			nums[mark], nums[pointer] = nums[pointer],nums[mark]
		}
		pointer++
	}
	nums[start],nums[mark] = nums[mark],nums[start]


	quickSort(nums,start,mark-1)
	quickSort(nums,mark+1,end)

	return nums
}

//堆排序
func heapSort(nums []int) []int {
	//初始化堆长度
	nLen := len(nums)
	//构造大顶堆
	for i:=nLen/2-1;i>=0;i-- {
		adjustHead(nums,i,nLen)
	}
	for nLen > 0 {
		nums[0],nums[nLen-1] = nums[nLen-1],nums[0]
		nLen--
		adjustHead(nums,0,nLen)
	}
	return nums
}
func adjustHead(nums []int,i int,nLen int)  {
	maxIndex := i
	if i*2+1<nLen && nums[i*2+1] > nums[maxIndex] {
		maxIndex = i*2+1
	}
	if i*2+2<nLen && nums[i*2+2] > nums[maxIndex]{
		maxIndex = i*2+2
	}
	if maxIndex != i {
		nums[i],nums[maxIndex] = nums[maxIndex],nums[i]
		adjustHead(nums,maxIndex,nLen)
	}
}
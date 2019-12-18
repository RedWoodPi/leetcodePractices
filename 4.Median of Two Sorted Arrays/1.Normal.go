package __Median_of_Two_Sorted_Arrays

import (
	"fmt"
)

func main()  {
	nums1 := []int{1,3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays1(nums1,nums2))
}

func findMedianSortedArrays1(nums1 []int, nums2 []int)  float64{
	m,n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays1(nums2,nums1)
	}
	iMin, iMax, laftLen := 0, m, (m+n+1)/2

	for iMin <= iMax {
		i := iMin + (iMax - iMin)/2
		j := laftLen - i

		if i > iMin && nums1[i-1] > nums2[j]{
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i]{
			iMin = i + 1
		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums1[i-1],nums2[j-1])
			}

			if (m+n)%2 != 0{
				return float64(leftMax)
			}

			rightMin := 0
			if i == m {
				rightMin = nums2[j]
			} else if j == n{
				rightMin = nums1[i]
			} else{
				rightMin = min(nums1[i],nums2[j])
			}
			return float64(leftMax+rightMin)/2.0
		}
	}

	return 0.0
}

func max(a,b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}

func min(a,b int )int{
	if a > b{
		return b
	}else{
		return a
	}
}

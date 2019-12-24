//package __Container_With_Most_Water
package main

import "fmt"

func main()  {
	arr := []int{1,1}
	fmt.Println(maxArea(arr))
}

func maxArea(height []int) int {
	temp := 0
	area := 0
	for k,_ := range height {
		for i:=k+1;i<len(height);i++{
			if height[k] > height[i] {
				area = height[i] * (i - k)
			} else {
				area = height[k] * (i - k)
			}
			if area > temp{
				temp = area
			}
		}

	}
	return temp
}
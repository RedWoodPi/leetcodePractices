package main

import (
	"fmt"
	"math"
)

func main()  {
	arr := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(doublePointer(arr))
}

func doublePointer(height []int) int {
	temp := 0
	area := 0
	lp := 0
	rp := len(height) - 1
	for lp < rp {
		area = int(math.Min(float64(height[lp]),float64(height[rp]))) * (rp - lp)
		if height[lp] > height[rp] {
			rp--
		} else {
			lp++
		}
		if area > temp {
			temp = area
		}
	}
	return temp
}
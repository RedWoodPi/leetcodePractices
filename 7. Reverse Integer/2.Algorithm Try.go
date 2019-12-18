package ___Reverse_Interger

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(reverse2(-123))
}

func reverse2(x int) int {
	temp := []int{}
	for x != 0 {
		temp = append(temp,x%10)
		x = x/10
	}
	tempLen := len(temp)
	y := 0
	for j:=0;j<tempLen;j++ {
		y = y + temp[j] * int(math.Pow10(tempLen - j - 1))
	}
	if y < int(math.Pow(-2,31)) || y > int(math.Pow(2,31)) -1 {
		y = 0
	}
	return y
}
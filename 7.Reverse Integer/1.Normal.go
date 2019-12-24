package ___Reverse_Interger

import (
	"fmt"
	"strconv"
	"math"
)

func main()  {
	fmt.Println(reverse1(124524545))
}
func reverse1(x int) int {
	//这个不是算法，只能说是解决方案
	signed := ""
	if x < 0 {
		x = -x
		signed = "-"
	}

	y := strconv.Itoa(x)[:]
	yLenth := len(y)
	temp := ""
	for i:=0;i<yLenth;i++ {
		temp = temp + string(y[yLenth-i-1])
	}
	temp = signed + temp
	reverseNumber,err := strconv.Atoi(temp)
	if err != nil {
		reverseNumber = 0
	}

	if reverseNumber < int(math.Pow(-2,31)) || reverseNumber > int(math.Pow(2,31)) -1 {
		reverseNumber = 0
	}

	return reverseNumber
}

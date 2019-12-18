package ___Palindrome_Number

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return  false
	}
	y := x
	temp := []int{}
	for x != 0 {
		temp = append(temp,x%10)
		x = x/10
	}
	t := 0
	for i:=0;i<len(temp);i++ {
		t = t + temp[i] * int(math.Pow10(len(temp)-i-1))
	}
	if y == t {
		return  true
	} else {
		return false
	}
}

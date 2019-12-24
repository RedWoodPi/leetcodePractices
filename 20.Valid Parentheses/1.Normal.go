package main

import "fmt"

func main()  {
	str3 := "[()](())"
	fmt.Println(isValid(str3))
}

func isValid(s string) bool {
	stackArr := []uint8{}
	for i:=0;i<len(s);i++ {
		if len(stackArr) > 0 {
			if s[i] == stackArr[len(stackArr)-1]+1 || s[i] == stackArr[len(stackArr)-1]+2{
				stackArr = stackArr[0:len(stackArr)-1]
			} else {
				stackArr = append(stackArr,s[i])
			}

		} else {
			stackArr = append(stackArr,s[i])
		}
	}
	fmt.Println(stackArr)
	if len(stackArr) == 0 {
		return true
	}
	return false
}
package main

import (
	"fmt"
	"strings"
)

func main()  {
	num1 := "12"
	num2 := "63"
	fmt.Println(multiply(num1,num2))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := strings.Repeat("0",len(num1)+len(num2))
	arr := strings.Split(res,"")
	for i:=len(num2)-1;i>=0;i--{
		for j:=len(num1)-1;j>=0;j--{
			temp := arr[i+j+1][0]-'0' + (num1[j]-'0')*(num2[i]-'0')
			arr[i+j+1] = string(temp%10 + '0')
			arr[i+j] = string(arr[i+j][0]-'0' + temp/10 + '0')
		}
	}
	for k,_ := range arr {
		if arr[k] != "0" {
			arr = arr[k:]
			break
		}
	}
	str := ""
	for _,v := range arr {
		str = str + v
	}
	return str
}

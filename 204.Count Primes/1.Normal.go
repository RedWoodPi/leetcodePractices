package main

import "fmt"

func main()  {
	fmt.Println(countPrimes(499979))
}

func countPrimes(n int) int {
	count := 0
	for i:=2;i<n;i++ {
		temp := true
		for j:=2;j<i;j++{
			if i%j == 0 {
				temp = false
				break
			}
		}
		if temp {
			count++
		}
	}

	return count
}

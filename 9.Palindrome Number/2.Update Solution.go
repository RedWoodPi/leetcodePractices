package __Palindrome_Number

import "fmt"

func main()  {
	fmt.Println(isPalindrome2(121))
}
func isPalindrome2(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0){
		return  false
	}
	temp := 0

	for x > temp {
		temp = temp*10 + x%10
		x = x/10
	}
	return x == temp || x == temp/10
}
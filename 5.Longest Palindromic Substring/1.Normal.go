package __Longest_Palindromic_Substring

import "fmt"

func main()  {
	s := "a"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	bts := []byte(s)
	strlen := len(s)

	temp1 := []byte{}
	if strlen != 0 {
		temp1 = []byte{bts[0]}
	}
	temp2 := []byte{}
	for k,_ := range bts {
		temp2 = []byte{bts[k]}
		for i:=1;k+i<strlen;i++ {
			if k-i >= 0 {
				if bts[k-i] == bts[k+i] {
					temp2 = bts[k-i:k+i+1]
					if len(temp1) < len(temp2) {
						temp1 = temp2
					}
				} else {
					break
				}
			}
		}

		if k+1 < strlen && bts[k] == bts[k+1] {
			temp2 = bts[k:k+2]
			if len(temp1) < len(temp2) {
				temp1 = temp2
			}
			for i:=1;k+i+1<strlen;i++ {
				if k-i >= 0 {
					if bts[k-i] == bts[k+i+1] {
						temp2 = bts[k-i:k+i+2]
						if len(temp1) < len(temp2) {
							temp1 = temp2
						}
					} else {
						break
					}
				}
			}

		}

	}
	return string(temp1)
}

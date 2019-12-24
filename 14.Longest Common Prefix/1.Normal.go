//package __Longest_Common_Prefix
package main

import (
	"fmt"
)

func main()  {
	input := []string{"a"}
	fmt.Println(longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//if len(strs) == 1 {
	//	return strs[0]
	//}
	minLen := len(strs[0])
	longest := ""
	for i:=0;i<minLen+1;i++ {
		longest = strs[0][0:i]
		for j:=0;j<len(strs);j++ {
			if len(strs[j]) == 0 {
				return ""
			}
			if len(longest) > len(strs[j]) || longest != strs[j][0:i] {
				return longest[0:i-1]
			}
		}
	}

	return longest
}
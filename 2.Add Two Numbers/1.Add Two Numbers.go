package __Add_Two_Numbers

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	a := ListNode{
		9,
		&ListNode{
			8,
			nil,
		},
	}

	b := ListNode{
		1,
		nil,
	}

	test := addTwoNumbers(&a, &b)

	for test != nil{
		fmt.Println(test.Val)
		test = test.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{0,nil}
	curr := &dummyHead
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1!=nil {
			x = l1.Val
		}

		if l2!=nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = sum/10
		curr.Next = &ListNode{sum%10,nil}
		curr = curr.Next

		if l1!=nil {
			l1 = l1.Next
		}
		if l2!=nil {
			l2 = l2.Next
		}

	}

	if carry > 0{
		curr.Next = &ListNode{carry,nil}
	}

	return dummyHead.Next
}

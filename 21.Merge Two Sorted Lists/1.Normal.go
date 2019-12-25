package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func main()  {
	l1 := ListNode{1,&ListNode{2,&ListNode{3,nil}}}
	l2 := ListNode{1,&ListNode{3,&ListNode{4,nil}}}
	fmt.Println(mergeTwoLists2(&l1,&l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1,l2.Next)
		return l2
	}
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := new(ListNode)
	current := preHead
	for l1!=nil && l2!=nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1==nil{
		current.Next=l2
	} else {
		current.Next=l1
	}

	return preHead.Next

}
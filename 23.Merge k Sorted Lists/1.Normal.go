package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func main()  {
	listNode1 := ListNode{1,&ListNode{4,&ListNode{5,nil}}}
	listNode2 := ListNode{1,&ListNode{3,&ListNode{4,nil}}}
	listNode3 := ListNode{2,&ListNode{6,nil}}
	lists := []*ListNode{
		nil,
		&listNode1,
		&listNode2,
		nil,
		&listNode3,
	}

	fmt.Println(mergeKLists(lists))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	preHead := lists[0]
	for i:=1;i<len(lists);i++ {
		current := preHead
		l2 := lists[i]
		if current == nil {
			preHead = l2
			continue
		}
		l1 := new(ListNode)
		*l1 = *current
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				current.Next = l1
				l1 = l1.Next
			} else {
				current.Next = l2
				l2 = l2.Next
			}
			current = current.Next
		}
		if l1 == nil {
			current.Next = l2
		}
		if l2 == nil {
			current.Next = l1
		}
		preHead = preHead.Next
	}
	return preHead

	//for preHead != nil {
	//	fmt.Println(preHead)
	//	preHead = preHead.Next
	//}
	//return nil
}
package main

import "fmt"

func main() {
	fmt.Println(mergeTwoLists(&ListNode{}, &ListNode{}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	result := &ListNode{}
	head := result

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			result.Next = list2
			list2 = list2.Next
		} else {
			result.Next = list1
			list1 = list1.Next
		}
		result = result.Next
	}

	if list1 == nil {
		result.Next = list2
	} else {
		result.Next = list1
	}

	return head.Next
}

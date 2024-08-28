package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var dummy ListNode
	current := &dummy

	for l1 != nil || l2 != nil || carry != 0 {
		var partialSum int

		if l1 != nil {
			partialSum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			partialSum += l2.Val
			l2 = l2.Next
		}

		partialSum += carry
		carry = partialSum / 10
		current.Next = &ListNode{Val: partialSum % 10}
		current = current.Next
	}

	return dummy.Next
}

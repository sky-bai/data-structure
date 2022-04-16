package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	low := head
	fast := head

	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}

	return low

}

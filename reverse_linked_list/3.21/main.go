package __21

type node struct {
	next  *node
	value int
}

func reverse_linked_list1(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	new_head := reverse_linked_list1(head.next)
	head.next.next = head
	head.next = nil
	return new_head
}

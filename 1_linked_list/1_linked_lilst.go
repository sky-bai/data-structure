package __linked_list

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) insertData(data interface{}) *ListNode {
	node := &ListNode{Val: data}

	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	return l.Head
}

func main() {
	arr := []int{}
	arr = append(arr, 1)
}

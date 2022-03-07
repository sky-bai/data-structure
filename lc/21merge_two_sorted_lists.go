package main

func main() {

}

type ListNode struct {
	value int
	next  *ListNode
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.value < l2.value {
		l1.next = Merge(l1.next, l2)
		return l1
	}
	l2.next = Merge(l1, l2.next)
	return l2
}

// Merge 理解递归 找到这个函数的输入输出是什么
// 其实这里的merge 输入是 两个链表 输出的是两个链表两个值里面最小的那一个。 输出就是 我要确定下一个节点。

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// 将该链表里面的值全部打印出来

func getAllValue(p *Node) {
	// if it is the nil,it is the tail node.
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	var head = new(Node)
	head.data = 1
	var second = new(Node)
	head.next = second
	second.data = 2
	getAllValue(head)

}

func reversed(p *Node) {
	for p != nil {

	}
}

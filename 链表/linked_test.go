package main

import (
	"fmt"
	"time"
)

// 1.定义一个节点的结构体
// 里面有两个属性

type LinkedNode struct {
	payload interface{}
	next    *LinkedNode // 指向下一个节点的指针
}

// 2.表明linkednode要做什么操作
type LinkedNoder interface {
	add(payload interface{})
}

func (head *LinkedNode) add(payload interface{}) {
	point := head // 拿到指向头节点的指针

	for point.next != nil {
		point = point.next // 拿到第二个节点的地址
	}
	newNode := LinkedNode{payload: payload, next: nil}
	point.next = &newNode
}

func main() {
	var a <-chan time.Time
	fmt.Print(a)
}

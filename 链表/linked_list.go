package main

type LinkNode struct {
	payload interface{}
	Next    *LinkNode
}

func (l *LinkNode) Add(payload interface{}) {
	print(l.payload)
}

func main() {
	l1 := LinkNode{payload: "bai"}
	l1.Add("")
}

// 创建链表 创建两个属性 获取链表的属性 链表最重要的的就是找到一个指针

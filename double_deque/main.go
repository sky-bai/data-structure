package double_deque

type (
	Deque struct {
		head *Node
		tail *Node
		size int
	}
	Node struct {
		value interface{}
		next  *Node
		prev  *Node
	}
)

func (d *Deque) New() *Deque {
	return &Deque{head: nil, tail: nil, size: 0}
}

// Peek 返回队头元素
func (d *Deque) Peek() interface{} {
	if d.size == 0 {
		return nil
	}

	return d.head.value
}

// Push 入队
func (d *Deque) Push(value interface{}) {
	tem := &Node{value: value, next: nil, prev: nil}
	if d.size == 0 {
		d.head = tem
		d.tail = d.head
	} else {
		tem.prev = d.tail
		d.tail.next = tem
		d.tail = tem
	}
	d.size++

	// 向队列尾部添加元素 1.连接队尾与新元素 2.更新队尾元素
}

// Pop 出队
func (d *Deque) Pop() interface{} {
	if d.size == 0 {
		return nil
	}

	tem := d.head
	if d.head.next == nil {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.prev.next = nil
		d.head.prev = nil
	}
	d.size--

	return tem.value
}

func main() {

}

package __19

import (
	"fmt"
	"testing"
)

func Test_Insert(T *testing.T) {
	heap1 := NewHeap(8)
	heap1.Insert(2)
	heap1.Insert(4)
	heap1.Insert(1)
	heap1.Insert(5)
	heap1.Insert(3)
	heap1.Insert(6)
	heap1.Insert(7)
	fmt.Println(heap1.arr)
}

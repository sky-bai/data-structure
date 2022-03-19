package __19

import (
	"fmt"
	"testing"
)

func Test_heapify(T *testing.T) {
	h2 := NewHeap(8)
	h2.Insert(2)
	h2.Insert(3)
	h2.Insert(4)
	h2.Insert(5)
	h2.Insert(6)
	h2.Insert(7)
	h2.Insert(8)

	for i := 0; i <= len(h2.arr); i++ {
		num := h2.GetTop()
		if num != -1 {
			fmt.Println(num)
		}

	}

}

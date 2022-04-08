package __21

import (
	"fmt"
	"testing"
)

func Test_reverse(T *testing.T) {
	n1 := new(node)
	n1.value = 1
	n2 := new(node)
	n2.value = 2

	n3 := new(node)
	n3.value = 3

	n4 := new(node)
	n4.value = 4
	n1.next = n2
	n2.next = n3
	n3.next = n4
	new_head := reverse_linked_list1(n1)
	fmt.Println(new_head.next.value)

}

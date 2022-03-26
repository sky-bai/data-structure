package main

func main() {

}

type heap struct {
	arr            []int
	currentDataNum int
}

func createHeap() {
	arrList := []int{1, 2, 4, 2, 0, 3, 0, 9}
	var h heap
	h.arr = append(h.arr, 0)
	for _, value := range arrList {
		h.insertHeap(value)
	}
}

// 创建堆的时候将每个元素放入最后

func (h *heap) insertHeap(n int) {
	h.arr = append(h.arr, n)
	length := len(h.arr)
	h.currentDataNum = length - 1
	h.heapify(h.currentDataNum)

}

func (h *heap) heapify(currentDataNum int) {
	if currentDataNum <= 1 {
		return
	}
	if currentDataNum == 2 {
		if h.arr[currentDataNum] > h.arr[currentDataNum-1] {
			h.swap(currentDataNum, currentDataNum-1)
		}
		return
	}
	i := currentDataNum
	if i/2 > 0 && h.arr[i] > h.arr[i/2] {
		h.swap(i, i/2)
		i = i / 2
	}
	return
}

func (h *heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]

}

package __19

type Heap struct {
	arr      []int
	capacity int
	used     int
}

func NewHeap(capacity int) *Heap {
	dataNum := make([]int, capacity+1)
	dataNum[0] = 0
	return &Heap{
		arr:      dataNum,
		capacity: capacity,
		used:     0,
	}
}

// Insert 插入的时候自下向上进行堆化
func (h *Heap) Insert(data int) {
	if h.used > h.capacity {
		return
	}

	h.used++
	point := h.used
	h.arr[point] = data
	for point/2 > 0 && h.arr[point] > h.arr[point/2] {
		h.arr[point], h.arr[point/2] = h.arr[point/2], h.arr[point]
		point /= 2
	}
}

// GetTop 获取堆顶元素
func (h *Heap) GetTop() int {
	if h.used == 0 {
		return -1
	}

	top := h.arr[1]
	h.arr[1] = h.arr[h.used]
	h.used--
	h.topToDownHeapify()
	return top

}

// topToDownHeapify
func (h *Heap) topToDownHeapify() {
	// 从上自下排序 是因为已经排好序 从下自上 最后一个根节点进行排序 是因为 全部都没有 排序
	i := 1
	maxPro := i
	length := h.used
	for {
		if 2*i <= length && h.arr[2*i] > h.arr[i] {
			maxPro = 2 * i
		}

		if 2*i+1 <= length && h.arr[2*i+1] > h.arr[maxPro] {
			maxPro = 2*i + 1
		}
		if i == maxPro {
			break
		}
		h.swap(i, maxPro)

		i = maxPro
	}

}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// CreateHeap 根据数组
func CreateHeap(a []int) []int {
	h := NewHeap(len(a))
	for i := 0; i < len(a); i++ {
		h.arr[i+1] = a[i]
	}

	point := len(h.arr)

	for point/2 > 0 && h.arr[point] > h.arr[point/2] {
		h.arr[point], h.arr[point/2] = h.arr[point/2], h.arr[point]
		point /= 2
	}
	return a
}

func heapify(a []int) {
	// 数组长度减一
	// 从上自下进行堆化 通过数组长度求解最后一个根节点指针
	length := len(a) - 1
	for i := length / 2; i >= 1; i++ {

		for {
			maxPro := i
			if 2*i >= length && a[2*i] > a[i] {
				maxPro = 2 * i // 如果左节点有 并且比根节点大 就记录左节点
			}
			if 2*i+1 >= length && a[2*i+1] > a[maxPro] {
				maxPro = 2*i + 1 // 如果右节点有 并且比上一个节点大 就记录右节点
			}
			if i == maxPro {
				break // 如果节点没有发生移动 表示该三个节点构成的树 已经排好序
			}
			a[maxPro], a[i] = a[i], a[maxPro] // 将最大的与根节点进行交换
			i = maxPro                        // 找到之前这个小树的最大节点 然后将指针移到该节点  只去查看已移动节点下面的子树
		}
	}

}

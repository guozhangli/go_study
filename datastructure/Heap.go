package TestProject

type Heap struct {
	Array []int
	Count int
	Type  int
}

func NewHeap(len int) *Heap {
	return &Heap{
		Array: make([]int, len),
		Count: len,
		Type:  0,
	}
}

func (heap *Heap) ParentIndex(index int) int {
	if index <= 0 || index >= heap.Count {
		return -1
	}
	return (index - 1) / 2
}

func (heap *Heap) LeftChildIndex(index int) int {
	left := index*2 + 1
	if left >= heap.Count {
		return -1
	}
	return left
}

func (heap *Heap) RightChildIndex(index int) int {
	right := index*2 + 2
	if right >= heap.Count {
		return -1
	}
	return right
}

func (heap *Heap) GetMaximum() int {
	if heap.Count == 0 {
		return -1
	}
	return heap.Array[0]
}

//堆化当前元素
func (heap *Heap) PercolateDown(index int) {
	l := heap.LeftChildIndex(index)
	r := heap.RightChildIndex(index)
	var max int
	if l != -1 && heap.Array[l] > heap.Array[index] {
		max = l
	} else {
		max = index
	}
	if r != -1 && heap.Array[r] > heap.Array[max] {
		max = r
	}
	if max != index {
		heap.Array[max], heap.Array[index] = heap.Array[index], heap.Array[max]
	} else {
		return
	}
	heap.PercolateDown(max)
}

func (heap *Heap) DeleteMax() int {
	if heap.Count == 0 {
		return -1
	}
	data := heap.Array[0]
	heap.Array[0] = heap.Array[heap.Count-1]
	heap.Count--
	heap.ParentIndex(0)
	return data
}

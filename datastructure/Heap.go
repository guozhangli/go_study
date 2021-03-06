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
	var max = index
	if index <= heap.Count/2 {
		if l != -1 && heap.Array[l] > heap.Array[max] {
			max = l
		}
		if r != -1 && heap.Array[r] > heap.Array[max] {
			max = r
		}
		if max != index {
			heap.Array[index], heap.Array[max] = heap.Array[max], heap.Array[index]
			heap.PercolateDown(max)
		}
	}
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

func (heap *Heap) Insert(data int) {
	heap.Count++
	i := heap.Count - 1
	heap.Array = append(heap.Array, 0)
	for i > 0 && data > heap.Array[(i-1)/2] {
		heap.Array[i] = heap.Array[(i-1)/2]
		i = (i - 1) / 2
	}
	heap.Array[i] = data
}

func (heap *Heap) DestoryHeap() {
	heap.Count = 0
	heap.Array = nil
}

func (heap *Heap) BuildHeap(arr []int) {
	if heap == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		heap.Array[i] = arr[i]
	}
	for i := (heap.Count - 1) / 2; i >= 0; i-- {
		heap.PercolateDown(i)
	}
}

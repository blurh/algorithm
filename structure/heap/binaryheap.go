package heap

// 最大堆在堆排序实现了, 这里实现最小堆

type binaryHeap struct {
    data   []int
    length int
}

func InitBinaryHeap() *binaryHeap {
    heap := new(binaryHeap)
    heap.data = []int{-1}
    heap.length = 0
    return heap
}

// 从第一个非叶子节点开始往前扫 sift down 即可
func Heapify(arr []int) *binaryHeap {
    heap := InitBinaryHeap()
    heap.data = append(heap.data, arr...)
    heap.length = len(arr)
    lastNotLeafIndex := parentIndex(len(arr))
    for i := lastNotLeafIndex; i > 0; i-- {
        heap.SiftDown(i)
    }
    return heap
}

func parentIndex(index int) int {
    return index / 2
}

func leftIndex(index int) int {
    return index * 2
}

func rightIndex(index int) int {
    return index*2 + 1
}

func (heap *binaryHeap) Len() int {
    return heap.length
}

func (heap *binaryHeap) Pop() int {
    if heap.Len() <= 0 {
        return -1
    }
    topOfHeap := heap.data[1]
    heap.data[1], heap.data[heap.Len()] = heap.data[heap.Len()], heap.data[1]
    heap.data = heap.data[0:heap.Len()]
    heap.length--
    heap.SiftDown(1)
    return topOfHeap
}

func (heap *binaryHeap) Push(value int) bool {
    if heap.Len() <= 0 {
        heap.data = append(heap.data, value)
    } else {
        heap.data = append(heap.data[:1], append([]int{value}, heap.data[1:]...)...)
    }
    heap.length++
    heap.SiftDown(1)
    return true
}

// 上浮跟父节点比较
func (heap *binaryHeap) SiftUp(index int) bool {
    for heap.data[index] < heap.data[parentIndex(index)] && index != 1 {
        heap.data[index], heap.data[parentIndex(index)] = heap.data[parentIndex(index)], heap.data[index]
        index = parentIndex(index)
    }
    return true
}

// 下沉跟子节点比较
func (heap *binaryHeap) SiftDown(index int) bool {
    // 左子节点已经超出长度, 说明没有左子节点, 即自身已经是叶子节点了无法下沉, 返回 true 即可
    if leftIndex(index) > heap.Len() {
        return true
    }
    for leftIndex(index) <= heap.Len() {
        if rightIndex(index) > heap.Len() {
            if heap.data[leftIndex(index)] < heap.data[index] {
                heap.data[leftIndex(index)], heap.data[index] = heap.data[index], heap.data[leftIndex(index)]
                index = leftIndex(index)
            } else {
                break
            }
        } else if rightIndex(index) <= heap.Len() {
            if heap.data[rightIndex(index)] < heap.data[index] || heap.data[leftIndex(index)] < heap.data[index] {
                if heap.data[rightIndex(index)] < heap.data[leftIndex(index)] {
                    heap.data[index], heap.data[rightIndex(index)] = heap.data[rightIndex(index)], heap.data[index]
                    index = rightIndex(index)
                } else if heap.data[rightIndex(index)] > heap.data[leftIndex(index)] {
                    heap.data[leftIndex(index)], heap.data[index] = heap.data[index], heap.data[leftIndex(index)]
                    index = leftIndex(index)
                }
            } else {
                break
            }
        } else {
            break
        }
    }
    return true
}

func (heap *binaryHeap) CheckHeap() bool {
    lastNotLeafIndex := parentIndex(heap.Len())
    for i := lastNotLeafIndex; i > 0; i-- {
        if (rightIndex(i) <= heap.Len() && heap.data[rightIndex(i)] < heap.data[i]) ||
            heap.data[leftIndex(i)] < heap.data[i] {
            return false
        }
    }
    return true
}

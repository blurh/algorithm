package sort

func HeapSort(arr []int) ([]int, error) {
    heap := heapify(arr)
    heapSize := len(arr) - 1
    for heapPos := heapSize; heapPos > 0; heapPos-- {
        siftDown(heap, heapPos, 0) // 交换完后, 堆大小要减一, 即下次循环的 heapPos 值
        heap[heapPos], heap[0] = heap[0], heap[heapPos]
    }
    return arr, nil
}

func heapify(arr []int) []int {
    lastNotLeafIndex := len(arr) * 2 // parent index: (i + 1) * 2
    // 0 即 root 节点, 也需要下沉
    for i := lastNotLeafIndex; i >= 0; i-- {
        siftDown(arr, len(arr)-1, i)
    }
    return arr
}

func leftIndex(index int) int {
    return index*2 + 1
}

func rightIndex(index int) int {
    return index*2 + 2
}

func siftDown(heap []int, heapPos, index int) bool {
    for {
        if rightIndex(index) > heapPos && leftIndex(index) <= heapPos {
            if heap[leftIndex(index)] >= heap[index] {
                heap[leftIndex(index)], heap[index] = heap[index], heap[leftIndex(index)]
                index = leftIndex(index)
            } else {
                break
            }
        } else if rightIndex(index) <= heapPos && leftIndex(index) <= heapPos {
            if heap[leftIndex(index)] >= heap[index] || heap[rightIndex(index)] >= heap[index] {
                if heap[leftIndex(index)] > heap[rightIndex(index)] {
                    heap[leftIndex(index)], heap[index] = heap[index], heap[leftIndex(index)]
                    index = leftIndex(index)
                } else if heap[leftIndex(index)] <= heap[rightIndex(index)] {
                    heap[rightIndex(index)], heap[index] = heap[index], heap[rightIndex(index)]
                    index = rightIndex(index)
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

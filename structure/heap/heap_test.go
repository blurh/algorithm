package heap

import (
    "testing"
)

func TestHeap(t *testing.T) {
    checkSort := func(arr []int) {
        lastValue := 0
        for _, v := range arr {
            if v > lastValue && lastValue != 0 {
                t.Errorf("check heap fail")
            }
            lastValue = v
        }
    }
    t.Run("test for binary heap", func(t *testing.T) {
        testArr := []int{3, 9, 19, 36, 17, 12, 25, 5, 100, 15, 6, 11, 13, 8, 1, 4, 20}
        heap := Heapify(testArr)
        if !heap.CheckHeap() {
            t.Errorf("heapify heap fail")
        }

        heap = InitBinaryHeap()
        checkArr := []int{}
        for _, v := range testArr {
            heap.Push(v)
        }
        for i := 0; i <= len(testArr); i++ {
            checkArr = append(checkArr, heap.Pop())
        }
        checkSort(checkArr)
    })
}

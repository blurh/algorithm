package sort

func InsertionSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    for offset := 1; offset <= len(lt) - 1; offset++ {
        // 往左移动游标 i, 游标所在值跟 offsetValue 相比较, 将 offsetValue 左移插入到相应的位置
        // 2, 4, 8, 9, 30, 1, 7, 20, 17, 40
        //                 ↑
        //               offset
        //             ←i
        offsetValue := lt[offset]
        for i := offset - 1; i >= 0 && lt[i] > offsetValue; i-- {
            lt[i], lt[i + 1] = lt[i + 1], lt[i]
        }
    }
    return lt, nil
}

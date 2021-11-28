package sort

func SelectionSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    for offset := 0; offset <= len(lt) - 1; offset ++ {
        min := offset
        for i := offset; i <= len(lt) - 1; i++ {
            if lt[i] <= lt[min] {
                min = i
            }
        }
        lt[offset], lt[min] = lt[min], lt[offset]
    }
    return lt, nil
}

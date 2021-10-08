package sort

func ShellSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    // gap 可以取 2, 3, 2^k - 1
    gap := 2
    step := len(lt) / gap
    for step >= 1 {
        for i := 0; i <= len(lt)/step-1; i++ {
            lt, _ = insertionSortByStep(lt, i, step)
        }
        step /= gap
    }
    return lt, nil
}

func insertionSortByStep(lt []int, pos, step int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    for offset := pos; offset <= len(lt)-1; offset += step {
        offsetValue := lt[offset]
        for i := offset - step; i >= pos && lt[i] > offsetValue; i -= step {
            lt[i], lt[i+step] = lt[i+step], lt[i]
        }
    }
    return lt, nil
}

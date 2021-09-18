package sort

func BubbleSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    for range lt {
        flag := true
        for i := 0; i < len(lt) - 1; i++ {
            if lt[i] > lt[i + 1] {
                lt[i], lt[i + 1] = lt[i + 1], lt[i]
                flag = false
            }
        }
        if flag {
            break
        }
    }
    return lt, nil
}

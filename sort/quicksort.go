package sort

func QuickSort(lt []int) ([]int, error){
    if len(lt) <= 1 {
        return lt, nil
    }
    head, tail := 0, len(lt) - 1
    midValue := lt[0]
    for i := 1; i <= tail; {
        if lt[i] > midValue {
            lt[tail], lt[i] = lt[i], lt[tail]
            tail --
        }else{
            lt[head], lt[i] = lt[i], lt[head]
            head ++
            i ++
        }
    }
    QuickSort(lt[:head])
    QuickSort(lt[head + 1:])
    return lt, nil
}

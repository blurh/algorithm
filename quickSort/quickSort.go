//usr/bin/env go run "$0" "$@"; exit
package quickSort

import (
    "fmt"
)

func main() {
    lt := []int{1, 3, 5, 2, 10, 7}
    quickSortLt, _ := QuickSort(lt)
    fmt.Println(quickSortLt)
}

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

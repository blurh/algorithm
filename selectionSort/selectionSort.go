//usr/bin/env go run "$0" "$@"; exit
package selectionSort
import "fmt"

func main() {
    lt := []int{9, 5, 3, 7, 2}
    sortLt, _ := SelectionSort(lt)
    fmt.Println(sortLt)
}

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

//usr/bin/env go run $0 $@; exit
package countingSort
import "fmt"

func main() {
    lt := []int{2, 7, 8, 2, 3, 1, 20, 3, 6, 3}
    sortLt, _ := CountingSort(lt)
    fmt.Println(sortLt)
}

func CountingSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    max := lt[0]
    for i := 1; i <= len(lt) - 1; i++ {
        if lt[i] >= lt[i - 1] && lt[i] > max {
            max = lt[i]
        }
    }
    tmpLt := make([]int, max + 1)
    for _, v := range lt {
        tmpLt[v] += 1
    }
    offset := 0
    for index, value := range tmpLt {
        for i := 0; i <= value - 1; i++ {
            lt[offset] = index
            offset ++
        }
    }
    return lt, nil
}
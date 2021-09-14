//usr/bin/env go run $0 $@; exit
package insertionSort
import "fmt"

func main(){
    lt := []int{30, 9, 8, 2, 4, 1, 7, 20, 17, 40}
    sortLt, _ := InsertionSort(lt)
    fmt.Println(sortLt)
}

func InsertionSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    for offset := 1; offset <= len(lt) - 1; offset++ {
        for i := 0; i <= offset - 1; i++ {
            offsetValue := lt[offset]
            if offsetValue <= lt[i] {
                lt = append(lt[:offset], lt[offset + 1:]...)
                lt = append(lt[:i], append([]int{offsetValue}, lt[i:]...)...)
                break
            }
        }
    }
    return lt, nil
}

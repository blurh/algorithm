//usr/bin/env go run $0 $@; exit
package megreSort

import "fmt"

func main(){
    lt := []int{2, 4, 4, 3, 7, 9}
    sortLt, _ := MegreSort(lt)
    fmt.Println(sortLt)
}

func MegreSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    halfLenLt := len(lt) / 2
    leftLt, _ := MegreSort(lt[:halfLenLt])
    rightLt, _ := MegreSort(lt[halfLenLt:])
    return megreLt(leftLt, rightLt), nil
}

func megreLt(leftLt, rightLt []int) ([]int) {
    leftLen := len(leftLt)
    rightLen := len(rightLt)
    leftOffset, rightOffset := 0, 0
    lt := []int{}
    for leftOffset < leftLen && rightOffset < rightLen {
        if rightLt[rightOffset] < leftLt[leftOffset] {
            lt = append(lt, rightLt[rightOffset])
            rightOffset ++
        } else if leftLt[leftOffset] < rightLt[rightOffset] {
            lt = append(lt, leftLt[leftOffset])
            leftOffset ++
        } else if leftLt[leftOffset] == rightLt[rightOffset] {
            lt = append(lt, rightLt[rightOffset])
            lt = append(lt, leftLt[leftOffset])
            rightOffset ++
            leftOffset ++
        }
    }
    // 把剩下的直接添加到 lt
    lt = append(lt, rightLt[rightOffset:]...)
    lt = append(lt, leftLt[leftOffset:]...)
    return lt
}

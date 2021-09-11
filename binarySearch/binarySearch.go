//usr/bin/env go run "$0" "$@"; exit
package binarySearch

import "fmt"

func main(){
    lt := []int{1, 2, 4, 8, 9, 10, 17, 27, 30, 43, 55}
    searchValue := 17
    searchResult, _ := BinarySearch(lt, searchValue)
    fmt.Println(searchResult, lt[searchResult])
}

func BinarySearch(lt []int, value int) (int, error){
    if len(lt) <= 1 {
        return 0, nil
    }
    head, tail := 0, len(lt) - 1
    var middle, middleValue int
    for {
        middle = (head + tail) / 2
        middleValue = lt[middle]
        if value > middleValue {
            head = middle + 1
        }else if value < middleValue {
            tail = middle - 1
        }else if value == middleValue {
            break
        }else if head > tail {
            return -1, nil
        }
    }
    return middle, nil
}



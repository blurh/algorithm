//usr/bin/env go test; exit
package selectionSort

import (
    "fmt"
    "time"
    "math/rand"
    "testing"
)

func TestQuickSort(t *testing.T){
    assertMsg := func(t *testing.T, lt []int) {
        lastValue := 0
        checkResult := true
        for _, v := range lt {
            if v >= lastValue {
                lastValue = v
            }else{
                checkResult = false
                break
            }
        }
        if ! checkResult {
            t.Error(lt)
        }else {
            fmt.Println(lt)
        }
    }
    t.Run("test of selection sort", func(t *testing.T){
        var lt [20]int
        rand.Seed(time.Now().Unix())
        for i := 0; i<=19; i++ {
            lt[i] = rand.Intn(100)
        }
        fmt.Println(lt)
        sortLt, _ := SelectionSort(lt[:])
        assertMsg(t, sortLt)
    })
}

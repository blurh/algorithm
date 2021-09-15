//usr/bin/env go test; exit
package radixSort

import (
    "fmt"
    "time"
    "math/rand"
    "testing"
)

func TestRadixSort(t *testing.T){
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
    t.Run("test of radix sort", func(t *testing.T){
        var lt [30]int
        rand.Seed(time.Now().Unix())
        for i := 0; i<=29; i++ {
            lt[i] = rand.Intn(1000)
        }
        fmt.Println(lt)
        sortLt, _ := RadixSort(lt[:])
        assertMsg(t, sortLt)
    })
}

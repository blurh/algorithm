//usr/bin/env go test ; exit
package sort

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

func TestSort(t *testing.T) {
    var lt [30]int
    rand.Seed(time.Now().Unix())
    for i := 0; i <= 29; i++ {
        lt[i] = rand.Intn(1000)
    }
    fmt.Println("lt:", lt)

    assertMsg := func(t *testing.T, lt []int) {
        lastValue := 0
        checkResult := true
        for _, v := range lt {
            if v >= lastValue {
                lastValue = v
            } else {
                checkResult = false
                break
            }
        }
        if !checkResult {
            t.Error(lt)
        } else {
            fmt.Println(lt)
        }
    }
    t.Run("test of bubble sort", func(t *testing.T) {
        sortLt, _ := BubbleSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of bucket sort", func(t *testing.T) {
        sortLt, _ := BucketSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of counting sort", func(t *testing.T) {
        sortLt, _ := CountingSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of insertion sort", func(t *testing.T) {
        sortLt, _ := InsertionSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of megre sort", func(t *testing.T) {
        sortLt, _ := MegreSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of quickSort", func(t *testing.T) {
        sortLt, _ := QuickSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of radix sort", func(t *testing.T) {
        sortLt, _ := RadixSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of selection sort", func(t *testing.T) {
        sortLt, _ := SelectionSort(lt[:])
        assertMsg(t, sortLt)
    })
    t.Run("test of shell sort", func(t *testing.T) {
        sortLt, _ := ShellSort(lt[:])
        assertMsg(t, sortLt)
    })
}

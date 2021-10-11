package search

import (
    "math/rand"
    "sort"
    "testing"
    "time"
)

func TestSearch(t *testing.T) {
    assertSearch := func(index, searchValue int) {
        if index != searchValue {
            t.Errorf("index '%d' search return '%d', fail", index, searchValue)
        }
    }
    const length = 9999
    var lt [length]int
    var randValue int
    rand.Seed(time.Now().Unix())
    for i := 0; i <= length-1; i++ {
        randValue = rand.Intn(length)
        if i == 0 {
            lt[i] = randValue
        } else {
            lt[i] = randValue + lt[i-1]
        }
    }
    sort.Sort(sort.IntSlice(lt[:]))
    index := rand.Intn(length)
    indexValue := lt[index]
    t.Run("test of linear", func(t *testing.T) {
        linearResult, _ := Linear(lt[:], indexValue)
        assertSearch(index, linearResult)
    })
    t.Run("test of binary search", func(t *testing.T) {
        binaryResult, _ := BinarySearch(lt[:], indexValue)
        assertSearch(index, binaryResult)
    })
    t.Run("test of interpolation search", func(t *testing.T) {
        interpolationResult, _ := Interpolation(lt[:], indexValue)
        assertSearch(index, interpolationResult)
    })
    t.Run("test of fibonacci search", func(t *testing.T) {
        fibonacciResult, _ := Fibonacci(lt[:], indexValue)
        assertSearch(index, fibonacciResult)
    })
    t.Run("test of block search", func(t *testing.T) {
        blockResult, _ := Block(lt[:], indexValue)
        assertSearch(index, blockResult)
    })
}

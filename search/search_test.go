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
    var lt [20]int
    var randValue int
    rand.Seed(time.Now().Unix())
    for i := 0; i <= 19; i++ {
        randValue = rand.Intn(100)
        if i == 0 {
            lt[i] = randValue
        } else {
            lt[i] = randValue + lt[i-1]
        }
    }
    sort.Sort(sort.IntSlice(lt[:]))
    index := rand.Intn(20)
    indexValue := lt[index]
    t.Run("test of linear", func(t *testing.T) {
        linearSearch, _ := Linear(lt[:], indexValue)
        assertSearch(index, linearSearch)
    })
    t.Run("test of binary search", func(t *testing.T) {
        binarySearch, _ := BinarySearch(lt[:], indexValue)
        assertSearch(index, binarySearch)
    })
    t.Run("test of interpolation search", func(t *testing.T) {
        interpolationSearch, _ := Interpolation(lt[:], indexValue)
        assertSearch(index, interpolationSearch)
    })
}

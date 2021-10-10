//usr/bin/env go test; exit
package search

import (
    "math/rand"
    "sort"
    "testing"
    "time"
)

func TestSearch(t *testing.T) {
    assertSearch := func(index, binarySearchRet int) {
        if index != binarySearchRet {
            t.Errorf("index '%d' binarySearchRet '%d', fail", index, binarySearchRet)
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
}

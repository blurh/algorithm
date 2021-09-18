//usr/bin/env go test; exit
package binarySearch
import (
    "fmt"
    "time"
    "sort"
    "testing"
    "math/rand"
)
func TestSearch(t *testing.T) {
    assertMsg := func (t *testing.T, index, binarySearchRet int){
        if index != binarySearchRet {
            t.Errorf("index '%d' binarySearchRet '%d'", index, binarySearchRet)
        }else{
            fmt.Printf("index '%d' binarySearchRet '%d'\n\n", index, binarySearchRet)
        }
    }
    t.Run("test of binarySearch", func(t *testing.T){
        var lt [20]int
        var randValue int
        rand.Seed(time.Now().Unix())
        for i := 0; i <= 19; i++ {
            randValue = rand.Intn(100)
            if i == 0 {
                lt[i] = randValue
            }else{
                lt[i] = randValue + lt[i - 1]
            }
        }
        sort.Sort(sort.IntSlice(lt[:]))
        fmt.Println(lt)
        index := rand.Intn(20)
        indexValue := lt[index]
        binarySearch, _ := BinarySearch(lt[:], indexValue)
        assertMsg(t, index, binarySearch)
    })
}


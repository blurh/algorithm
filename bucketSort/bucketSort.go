//usr/bin/env go run "$0" "$@"; exit
package bucketSort
import (
    "fmt"
    "math"
)

func main() {
    lt := []int{15, 49, 48, 46, 58, 17, 80, 78, 24, 42, 56, 22, 64, 6, 6, 31, 15, 55, 41, 66}
    sortLt, _ := BucketSort(lt)
    fmt.Println(sortLt)
}

func BucketSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    // 分桶一般 k^2 = len(lt)
    bucketSize := int(math.Floor(math.Log(float64(len(lt)))))
    min, max := lt[0], lt[0]
    for i := 1; i <= len(lt) - 1; i++ {
        if lt[i] > max {
            max = lt[i]
        }
        if lt[i] < min {
            min = lt[i]
        }
    }
    bucketNum := (max - min) / bucketSize
    // 创建二维数组
    bucketLt := make([][]int, bucketNum + 1)
    for _, v := range lt {
        bucketCount := (v - min - 1)/ bucketSize
        bucketLt[bucketCount] = append(bucketLt[bucketCount], v)
    }
    // 排序各个桶进行合并
    for bucket := 0; bucket <= bucketNum; bucket++ {
        bucketLt[bucket], _ = quickSort(bucketLt[bucket])
        if bucket == 0 {
            lt = bucketLt[bucket]
        } else {
            lt = append(lt, bucketLt[bucket]...)
        }
    }
    return lt, nil
}

func quickSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    left, right := 0, len(lt) - 1
    midValue := lt[0]
    for i := 1; i <= right; {
        if lt[i] > midValue {
            lt[i], lt[right] = lt[right], lt[i]
            right --
        } else {
            lt[i], lt[left] = lt[left], lt[i]
            left ++
            i ++
        }
    }
    quickSort(lt[:left])
    quickSort(lt[left + 1:])
    return lt, nil
}

//usr/bin/env go run "$0" "$@"; exit
package bucketSort
import (
    "fmt"
    "math"
)

func main() {
    lt := []int{13, 28, 51, 14, 29, 86, 34, 36, 94, 67, 48, 49, 5, 92, 79, 27, 0, 27, 80, 19}
    sortLt, _ := BucketSort(lt)
    fmt.Println(lt)
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
    // 入桶时插入排序
    for _, v := range lt {
        bucketCount := (v - min - 1)/ bucketSize
        if len(bucketLt[bucketCount]) == 0 {
            bucketLt[bucketCount] = append(bucketLt[bucketCount], v)
            continue
        }
        for bucketIndex, bucketValue := range bucketLt[bucketCount] {
            if v <= bucketValue {
                bucketLt[bucketCount] = append(bucketLt[bucketCount][:bucketIndex],
                  append([]int{v}, bucketLt[bucketCount][bucketIndex:]...)...)
                break
            } else if v > bucketValue && bucketIndex == len(bucketLt[bucketCount]) - 1 {
                bucketLt[bucketCount] = append(bucketLt[bucketCount], v)
            }
        }
    }
    // 合并桶
    lt = bucketLt[0]
    for bucket := 1; bucket <= bucketNum; bucket++ {
        lt = append(lt, bucketLt[bucket]...)
    }
    return lt, nil
}

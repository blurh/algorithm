package sort

import "math"

func RadixSort(lt []int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    max := lt[0]
    for _, v := range lt {
        if v > max {
            max = v
        }
    }
    radix := 0
    for i := 1; i <= max; {
        i *= 10
        radix ++
    }
    for i := 0; i < radix; i++ {
        lt, _ = radixBucketSort(lt, i)
    }
    return lt, nil
}

func radixBucketSort(lt []int, radix int) ([]int, error) {
    if len(lt) <= 1 {
        return lt, nil
    }
    // 10 的 radix 次方数
    radixValue := int(math.Pow(10, float64(radix)))
    // 取相应的位数, 如 978 取百位数即: 978 / 100 % 10 = 9
    max := lt[0] / radixValue % 10
    for i := 1; i <= len(lt) - 1; i++ {
        if lt[i] / radixValue % 10 > max {
            max = lt[i] / radixValue % 10
        }
    }
    tmpLt := make([][]int, max + 1)
    for _, v := range lt {
        tmpLt[v / radixValue % 10] = append(tmpLt[v / radixValue % 10], v)
    }
    for i := range tmpLt {
        if i == 0 {
            lt = tmpLt[i]
        } else {
            lt = append(lt, tmpLt[i]...)
        }
    }
    return lt, nil
}

package search

func Fibonacci(arr []int, value int) (int, error) {
    if len(arr) < 1 {
        return -1, nil
    }
    low, high := 0, len(arr)-1
    for low <= high {
        mid := fibonacci(high-low) + low - 1
        if arr[mid] == value {
            return mid, nil
        } else if arr[mid] < value {
            low = mid - 1
        } else if arr[mid] > value {
            high = mid + 1
        }
    }
    return -1, nil
}

func fibonacci(max int) int {
    x, y := 0, 1
    for i := 0; ; i++ {
        x, y = y, x+y
        if y >= max {
            return x
        }
    }
}

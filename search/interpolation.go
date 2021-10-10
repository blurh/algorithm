package search

import "errors"

func Interpolation(arr []int, value int) (int, error) {
    errMsg := errors.New("value not found")
    if len(arr) <= 1 {
        return -1, errMsg
    }
    low, high := 0, len(arr)-1
    for low < high {
        mid := low + int(float64(float64((value-arr[low])*(high-low))/float64(arr[high]-arr[low])))
        if arr[mid] == value {
            return mid, nil
        } else if arr[mid] > value {
            high = mid - 1
        } else if arr[mid] < value {
            low = mid + 1
        }
    }
    return -1, errMsg
}

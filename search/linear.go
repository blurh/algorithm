package search

import "errors"

func Linear(arr []int, value int) (int, error) {
    for i, v := range arr {
        if v == value {
            return i, nil
        }
    }
    errMsg := errors.New("not found")
    return -1, errMsg
}

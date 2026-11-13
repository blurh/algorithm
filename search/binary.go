package search

func BinarySearch(lt []int, value int) (int, error){
    if len(lt) == 0 {
        return -1, nil
    }
    head, tail := 0, len(lt) - 1
    var middle, middleValue int
    for head <= tail {
        middle = (head + tail) / 2
        middleValue = lt[middle]
        if value > middleValue {
            head = middle + 1
        }else if value < middleValue {
            tail = middle - 1
        }else {
            return middle, nil
        }
    }
    return -1, nil
}



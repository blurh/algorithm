package binarySearch

func BinarySearch(lt []int, value int) (int, error){
    if len(lt) <= 1 {
        return 0, nil
    }
    head, tail := 0, len(lt) - 1
    var middle, middleValue int
    for {
        middle = (head + tail) / 2
        middleValue = lt[middle]
        if value > middleValue {
            head = middle + 1
        }else if value < middleValue {
            tail = middle - 1
        }else if value == middleValue {
            break
        }else if head > tail {
            return -1, nil
        }
    }
    return middle, nil
}



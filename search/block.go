package search

const blockNum = 10

type Index struct {
    start, end int
}

var indeices [blockNum]*Index

func Block(arr []int, value int) (int, error) {
    if len(arr) < 1 {
        return -1, nil
    }
    for i := 0; i < blockNum; i++ {
        index := new(Index)
        index.start = i * len(arr) / blockNum
        index.end = (i+1)*len(arr)/blockNum - 1
        indeices[i] = index
    }
    for i := 0; i < blockNum; i++ {
        if arr[indeices[i].start] <= value && arr[indeices[i].end] >= value {
            for j := indeices[i].start; i <= indeices[i].end; j++ {
                if value == arr[j] {
                    return j, nil
                }
            }
        }
    }
    return -1, nil
}

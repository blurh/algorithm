package hashing

func DJB(str string) uint64 {
    hash := uint64(5381)
    for i := 0; i <= len(str)-1; i++ {
        hash = hash<<5 + hash + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}

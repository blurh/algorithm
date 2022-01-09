package hashing

func JS(str string) uint64 {
    hash := uint64(1315423911)
    for i := 0; i <= len(str)-1; i++ {
        hash ^= ((hash << 5) + (hash >> 2) + uint64(str[i]))
    }
    return hash & 0x7FFFFFFF
}

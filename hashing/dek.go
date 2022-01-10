package hashing

func DEK(str string) uint64 {
    hash := uint64(1315423911)
    for i := 0; i <= len(str)-1; i++ {
        hash = ((hash << 5) ^ (hash >> 27)) ^ uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}

package hashing

func AP(str string) uint64 {
    hash := uint64(0)
    for i := 0; i <= len(str)-1; i++ {
        if (i & 1) == 0 {
            hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
        } else {
            hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5))
        }
    }
    return hash & 0x7FFFFFFF
}

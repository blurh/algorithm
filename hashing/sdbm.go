package hashing

func SDBM(str string) uint64 {
    hash := uint64(0)
    for i := 0; i <= len(str)-1; i++ {
        hash = hash<<6 + hash<<16 - hash + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}

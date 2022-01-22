package hashing

func BKDR(str string) uint64 {
    const seed = uint64(31) // 31 131 1313 13131 131313 etc..
    hash := uint64(0)
    for i := 0; i <= len(str)-1; i++ {
        hash = hash*seed + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}

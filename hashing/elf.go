package hashing

func ELF(str string) uint64 {
    hash := uint64(0)
    for i := 0; i <= len(str)-1; i++ {
        hash = (hash << 4) + uint64(str[i])
        x := hash & 0xF0000000
        if x != 0 {
            hash ^= (x >> 24)
            hash &= ^x
        }
    }
    return hash & 0x7FFFFFFF
}

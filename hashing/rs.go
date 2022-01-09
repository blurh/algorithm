package hashing

func RS(str string) uint64 {
    a := uint64(63689)
    b := uint64(378551)
    hash := uint64(0)
    for i := 0; i <= len(str)-1; i++ {
        hash = hash*a + uint64(str[i])
        a = a * b
    }
    return hash & 0x7FFFFFFF
}

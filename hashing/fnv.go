package hashing

// 32 位 FNV-1a
func FNV(str string) uint64 {
    prime := uint64(16777619)
    hash := uint64(2166136261) // offset
    for i := 0; i <= len(str)-1; i++ {
        hash ^= uint64(str[i])
        hash *= prime
    }
    return hash & 0x7FFFFFFF
}

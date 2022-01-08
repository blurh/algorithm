package hashing

func DJB(word string) uint64 {
    hash := uint64(5381)
    for i := 0; i <= len(word)-1; i++ {
        hash = hash<<5 + hash + uint64(word[i])
    }
    return hash & 0x7FFFFFFF
}

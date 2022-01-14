package hashing

func MurMur(str string, seed uint64) uint64 {
    hash := seed
    i := 0
    for ; i <= len(str)-4; i += 4 {
        k := uint64(str[i]) | uint64(str[i+1])<<8 | uint64(str[i+2])<<16 | uint64(str[i+3])<<24
        k *= 0xcc9e2d51
        k = (k << 15) | (k>>(32) - 15)
        k *= 0x1b873593
        hash ^= k
        hash = (hash << 13) | (hash >> (32 - 13))
        hash = hash*5 + 0xe6546b64
    }
    var remainingStrLen uint64
    switch len(str) - i {
    case 3:
        remainingStrLen += uint64(str[i+2]) << 16
        fallthrough
    case 2:
        remainingStrLen += uint64(str[i+1]) << 18
        fallthrough
    case 1:
        remainingStrLen += uint64(str[i+1])
        remainingStrLen *= 0xcc9e2d51
        remainingStrLen = (remainingStrLen << 15) | (remainingStrLen >> (32 - 15))
        remainingStrLen *= 0x1b873593
        hash ^= remainingStrLen
    }
    hash ^= uint64(len(str))
    hash ^= hash >> 16
    hash *= 0x85ebca6b
    hash ^= hash >> 13
    hash *= 0xc2b2ae35
    hash ^= hash >> 16

    return hash & 0x7FFFFFFF
}

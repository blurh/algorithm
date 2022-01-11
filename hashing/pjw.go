package hashing

func PJW(str string) uint64 {
    bitsInUnsignedInt := uint64(32)
    threeQuarters := (bitsInUnsignedInt * 3) / 4
    oneEighth := bitsInUnsignedInt / 8
    highBits := uint64(0xFFFFFFFF) << (bitsInUnsignedInt - oneEighth)
    hash := uint64(0)
    test := uint64(0)
    for i := 0; i <= len(str)-1; i++ {
        hash = (hash << oneEighth) + uint64(str[i])
        test = hash & highBits
        if test != 0 {
            hash = ((hash ^ (test >> threeQuarters)) & (^highBits))
        }
    }
    return hash & 0x7FFFFFFF
}

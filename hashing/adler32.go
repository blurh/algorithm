package hashing

func Adler32(str string) uint64 {
    var (
        mod uint64 = 65521
        a   uint64 = 1
        b   uint64 = 0
    )

    for i := 0; i <= len(str)-1; i++ {
        a = (a + uint64(str[i])) % mod
        b = (b + a) % mod
    }

    hash := (b << 16) | a
    return hash & 0x7FFFFFFF
}

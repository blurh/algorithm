package hashing

func SIP(str string, k0, k1 uint64) uint64 {
    p := []byte(str)
    // Initialization.
    v0 := k0 ^ 0x736f6d6570736575
    v1 := k1 ^ 0x646f72616e646f6d
    v2 := k0 ^ 0x6c7967656e657261
    v3 := k1 ^ 0x7465646279746573
    t := uint64(len(p)) << 56

    // Compression.
    for len(p) >= 8 {
        m := uint64(p[0]) | uint64(p[1])<<8 | uint64(p[2])<<16 | uint64(p[3])<<24 |
            uint64(p[4])<<32 | uint64(p[5])<<40 | uint64(p[6])<<48 | uint64(p[7])<<56
        v3 ^= m

        // Round 1.
        v0 += v1
        v1 = v1<<13 | v1>>(64-13)
        v1 ^= v0
        v0 = v0<<32 | v0>>(64-32)

        v2 += v3
        v3 = v3<<16 | v3>>(64-16)
        v3 ^= v2

        v0 += v3
        v3 = v3<<21 | v3>>(64-21)
        v3 ^= v0

        v2 += v1
        v1 = v1<<17 | v1>>(64-17)
        v1 ^= v2
        v2 = v2<<32 | v2>>(64-32)

        // Round 2.
        v0 += v1
        v1 = v1<<13 | v1>>(64-13)
        v1 ^= v0
        v0 = v0<<32 | v0>>(64-32)

        v2 += v3
        v3 = v3<<16 | v3>>(64-16)
        v3 ^= v2

        v0 += v3
        v3 = v3<<21 | v3>>(64-21)
        v3 ^= v0

        v2 += v1
        v1 = v1<<17 | v1>>(64-17)
        v1 ^= v2
        v2 = v2<<32 | v2>>(64-32)

        v0 ^= m
        p = p[8:]
    }

    // Compress last block.
    switch len(p) {
    case 7:
        t |= uint64(p[6]) << 48
        fallthrough
    case 6:
        t |= uint64(p[5]) << 40
        fallthrough
    case 5:
        t |= uint64(p[4]) << 32
        fallthrough
    case 4:
        t |= uint64(p[3]) << 24
        fallthrough
    case 3:
        t |= uint64(p[2]) << 16
        fallthrough
    case 2:
        t |= uint64(p[1]) << 8
        fallthrough
    case 1:
        t |= uint64(p[0])
    }

    v3 ^= t

    // Round 1.
    v0 += v1
    v1 = v1<<13 | v1>>(64-13)
    v1 ^= v0
    v0 = v0<<32 | v0>>(64-32)

    v2 += v3
    v3 = v3<<16 | v3>>(64-16)
    v3 ^= v2

    v0 += v3
    v3 = v3<<21 | v3>>(64-21)
    v3 ^= v0

    v2 += v1
    v1 = v1<<17 | v1>>(64-17)
    v1 ^= v2
    v2 = v2<<32 | v2>>(64-32)

    // Round 2.
    v0 += v1
    v1 = v1<<13 | v1>>(64-13)
    v1 ^= v0
    v0 = v0<<32 | v0>>(64-32)

    v2 += v3
    v3 = v3<<16 | v3>>(64-16)
    v3 ^= v2

    v0 += v3
    v3 = v3<<21 | v3>>(64-21)
    v3 ^= v0

    v2 += v1
    v1 = v1<<17 | v1>>(64-17)
    v1 ^= v2
    v2 = v2<<32 | v2>>(64-32)

    v0 ^= t

    // Finalization.
    v2 ^= 0xff

    // Round 1.
    v0 += v1
    v1 = v1<<13 | v1>>(64-13)
    v1 ^= v0
    v0 = v0<<32 | v0>>(64-32)

    v2 += v3
    v3 = v3<<16 | v3>>(64-16)
    v3 ^= v2

    v0 += v3
    v3 = v3<<21 | v3>>(64-21)
    v3 ^= v0

    v2 += v1
    v1 = v1<<17 | v1>>(64-17)
    v1 ^= v2
    v2 = v2<<32 | v2>>(64-32)

    // Round 2.
    v0 += v1
    v1 = v1<<13 | v1>>(64-13)
    v1 ^= v0
    v0 = v0<<32 | v0>>(64-32)

    v2 += v3
    v3 = v3<<16 | v3>>(64-16)
    v3 ^= v2

    v0 += v3
    v3 = v3<<21 | v3>>(64-21)
    v3 ^= v0

    v2 += v1
    v1 = v1<<17 | v1>>(64-17)
    v1 ^= v2
    v2 = v2<<32 | v2>>(64-32)

    // Round 3.
    v0 += v1
    v1 = v1<<13 | v1>>(64-13)
    v1 ^= v0
    v0 = v0<<32 | v0>>(64-32)

    v2 += v3
    v3 = v3<<16 | v3>>(64-16)
    v3 ^= v2

    v0 += v3
    v3 = v3<<21 | v3>>(64-21)
    v3 ^= v0

    v2 += v1
    v1 = v1<<17 | v1>>(64-17)
    v1 ^= v2
    v2 = v2<<32 | v2>>(64-32)

    // Round 4.
    v0 += v1
    v1 = v1<<13 | v1>>(64-13)
    v1 ^= v0
    v0 = v0<<32 | v0>>(64-32)

    v2 += v3
    v3 = v3<<16 | v3>>(64-16)
    v3 ^= v2

    v0 += v3
    v3 = v3<<21 | v3>>(64-21)
    v3 ^= v0

    v2 += v1
    v1 = v1<<17 | v1>>(64-17)
    v1 ^= v2
    v2 = v2<<32 | v2>>(64-32)

    hash := v0 ^ v1 ^ v2 ^ v3
    return hash & 0x7FFFFFFF
}
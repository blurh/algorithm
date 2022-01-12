package hashing

func Jenkins(num uint64) uint64 {
    num = (^num) + (num << 21)
    num = num ^ (num >> 24)
    num = (num + (num << 3)) + (num << 8)
    num = num ^ (num >> 14)
    num = (num + (num << 2)) + (num << 4)
    num = num ^ (num >> 28)
    num = num + (num << 31)
    return num
}

func JenkinsInvert(num uint64) uint64 {
    var tmp uint64

    // Invert num = num + (num << 31)
    tmp = num - (num << 31)
    num = num - (tmp << 31)

    // Invert num = num ^ (num >> 28)
    tmp = num ^ num>>28
    num = num ^ tmp>>28

    // Invert num *= 21
    num *= uint64(14933078535860113213)

    // Invert num = num ^ (num >> 14)
    tmp = num ^ num>>14
    tmp = num ^ tmp>>14
    tmp = num ^ tmp>>14
    num = num ^ tmp>>14

    // Invert num *= 265
    num *= uint64(15244667743933553977)

    // Invert num = num ^ (num >> 24)
    tmp = num ^ num>>24
    num = num ^ tmp>>24

    // Invert num = (^num) + (num << 21)
    tmp = ^num
    tmp = ^(num - (tmp << 21))
    tmp = ^(num - (tmp << 21))
    num = ^(num - (tmp << 21))

    return num
}

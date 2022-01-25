package hashing

func RK(str string) uint64 {
    const coefficient = uint64(1)
    const mod = uint64(65521)
    const windows = 4
    hash := uint64(0)
    // 计算初始窗口的 hash
    for i := 0; i <= windows-1; i++ {
        hash += (uint64(str[i]) * coefficient) % mod
    }
    for i := 1; i <= len(str)-1; i++ {
        // 减去左边滚动出去的值
        hash -= (uint64(str[i-1]) * coefficient) % mod
        // 加上右边滚动进来的值
        if i+windows <= len(str)-1 {
            hash += (uint64(str[i+windows-1]) * coefficient) % mod
        }
    }

    return hash
}

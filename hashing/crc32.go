package hashing

func CRC32(data string) uint64 {
    crc := uint64(0xFFFFFFFF)
    for i := 0; i <= len(data)-1; i++ {
        crc ^= uint64(data[i])
        for j := 0; j < 8; j++ {
            if (crc & 1) != 0 {
                crc = (crc >> 1) ^ 0xEDB88320 // 0xEDB88320 = reverse 0x04C11DB7
            } else {
                crc = (crc >> 1)
            }
        }
    }
    return ^crc & 0x7FFFFFFF
}

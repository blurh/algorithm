# Hash 算法

## 基础概念

**Hash 函数**: 将任意长度的输入值转变成固定长度的值输出，该值称为散列值（通常为字母与数字组合）

**类型**: 
- 整数 hash: 直接取余法、乘积取整法、平方取中法
- 字符串 hash: 可看成 256 进制（ANSI 为 128 进制）的大整数

**为什么常用素数取模?**

质数的因子只有 1 和它本身，对比合数可以获得更好的散列效果：

```
mod = 6 (合数):
  2的倍数: 2→2, 4→4, 6→0, 8→2, 10→4, 12→0  (碰撞多)
  3的倍数: 3→3, 6→0, 9→3, 12→0

mod = 7 (质数):
  2的倍数: 2→2, 4→4, 6→6, 8→1, 10→3, 12→5  (分布均匀)
  3的倍数: 3→3, 6→6, 9→2, 12→5
```

**技巧**: `hash & 0x7FFFFFFF` 可将 hash 限制到 31 位正数范围

---

## 字符串 Hash 算法

### BKDR (Brian Kernighan & Dennis Ritchie)
- **应用**: Java 字符串哈希
- **公式**: `h = h * seed + c`  (seed 常用 31, 131, 1313, 13131...)

### DJB (Daniel J. Bernstein) 
- **别名**: Times33 算法
- **应用**: ElasticSearch 索引文档 key 哈希
- **公式**: 
```go
hash = 5381
h = (h << 5) + h + c  // h = h * 33 + c
```

### SDBM
- **来源**: SDBM 开源数据库引擎
- **公式**: `h = (h << 6) + (h << 16) - h + c`  // h = h * 65599 + c

### RS (Robert Sedgwicks)
- **来源**: Algorithms in C 一书
- **公式**:
```go
a = 63689
b = 378551
hash = hash * a + c
a = a * b
```

### JS (Justin Sobel)
- **公式**:
```go
h = 1315423911
h ^= ((h << 5) + c + (h >> 2))
```

### AP (Arash Partow)
- **公式**:
```go
if (i & 1) == 0 {
    hash ^= (hash << 7) ^ uint64(c) ^ (hash >> 3)
} else {
    hash ^= ^((hash << 11) ^ uint64(c) ^ (hash >> 5))
}
```

### DEK (Donald E. Knuth)
- **来源**: The Art of Computer Programming Volume 3
- **公式**: 
```go
h = 1315423911
h = ((h << 5) ^ (h >> 27)) ^ c
```

### FNV (Fowler-Noll-Vo)
- **发明人**: Glenn Fowler, Landon Curt Noll, Phong Vo
- **应用**: Unix 系统、微软 hash_map
- **公式**:
```go
offset = 2166136261
prime = 16777619
hash ^= c
hash *= prime
```

### ELF (Executable and Linkable Format)
- **应用**: Linux 内核 (linux 2.4.0 irqueue.c)
- **公式**:
```go
hash = (hash << 4) + uint64(c)
x := hash & 0xF0000000
if x != 0 {
    hash ^= (x >> 24)
    hash &= ^x
}
```

### PJW (Peter J. Weinberger)
- **来源**: AT&T 贝尔实验室论文

---

## 高级 Hash 算法

### RK / Rolling Hash (Rabin-Karp)
- **特点**: 移动窗口计算 hash 值
- **用途**: 字符串匹配，常与其他哈希函数配合使用

### Jenkins (Bob Jenkins)
- **特点**: 雪崩性（Avalanche）、可逆性

### MurMur (Multiply and Rotate)
- **特点**: 对规律性较强的 key 有良好的随机分布
- **应用**: Redis、Memcache、Cassandra、Lucene

### SipHash
- **发明**: Jean-Philippe Aumasson & Daniel J. Bernstein (2012)
- **用途**: 防御 hash-flooding DoS 攻击

---

## 校验和算法

### Adler-32 (Mark Adler)
- **特点**: 计算快但不安全，容易被伪造
- **应用**: rsync

### CRC32 (Cyclic Redundancy Check)
- **特点**: 循环冗余校验
- **初始值**: 0xFFFFFFFF
- **多项式**: x³² + x²⁶ + x²³ + x²² + x¹⁶ + x¹² + x¹¹ + x¹⁰ + x⁸ + x⁷ + x⁵ + x⁴ + x² + x + 1
# hash
- 概念:  
将任意长度的输入值转变成固定长度的值输出, 该值称为散列值, 输出值通常为字母与数字组合  
hash & 0x7FFFFFFF 把 hash 限制到 31 位
- 类型:   
整数 hash 函数常用方法有三种：直接取余法、乘积取整法、平方取中法  
字符串可以看成 256 进制(ANSI 字符串为 128 进制)的大整数  
- hash 常用素数取模的原因:  
对合数取模, 那么所有该函数的因子的倍数冲突的概率会增大, 而质数的因子只有 1 和它本身, 所以对特定倍数数字来说, 会有更好散列效果, 如:  
若 mod 为 6, 对于 2 的倍数 2、4、6、8、10、12 的 hash 值是 2、4、0、2、4、0, 对于 3 的倍数 3、6、9、12 的 hash 值是 3、0、3、0
若 mod 为 7，对于 2、4、6、8、10、12 的 hash 值是 2、4、6、1、3、5, 对于 3 的倍数 3、6、9、12 的 hash 值是 3、6、2、5
可以看出, 如果 mod 是质数的话会得到更好的散列效果

<br />

- BKDR(Brian Kernighan & Dennis Ritchie):  
在 Java 字符串哈希值计算应用到  
```golang
h = h * seed + c  // seed = 31 131 1313 13131...
```

- DJB(Daniel J. Bernstein):  
俗称 Times33 算法, ElasticSearch 利用 DJB2 对要索引文档的指定 key 进行哈希  
```golang
hash = 5381
h = h << 5 + h + c // 即 h = h * 33 + c
```

- SDBM:  
在开源数据库引擎项目 SDBM 中被应用而得名
```golang
h = h << 6 + h << 16 - h + c // 即 h = h * 65599 + c
```

- RS(Robert Sedgwicks):  
因 Robert Sedgwicks 在其 Algorithms in C 一书中展示而得名
```golang
a = 63689
b = 378551
hash = hash*a + c
a = a * b
```

- JS(Justin Sobel):  
由 Justin Sobel 发明  
```golang
h = 1315423911
h ^= ((h << 5) + c + (h >> 2))
```

- AP(Arash Partow):  
由 Arash Partow 发明
```golang
if (i & 1) == 0 {
    hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
} else {
    hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5))
}
```

- DEK(Donald E. Knuth):  
由 Donald E. Knuth 在 Art Of Computer Programming Volume 3 中展示而得名
```golang 
h := uint64(1315423911)
h = ((h << 5) ^ (h >> 27)) ^ c)
```

- FNV(Fowler-Noll-Vo):  
以三位发明人 Glenn Fowler, Landon Curt Noll, Phong Vo 名字命名  
Unix system 系统中使用的一种著名 hash 算法, 后来微软也在其 hash_map 中实现  
```golang 
offset = 2166136261
prime = 16777619
hash ^= c
hash *= prime
```

- ELF(Executable and Linkable Format):  
ELF 在 Linux 中使用较多, linux 内核 ELF: [linux 2.4.0 -> irqueue.c]
```golang
hash = (hash << 4) + uint64(str[i])
x := hash & 0xF0000000
if x != 0 {
    hash ^= (x >> 24)
    hash &= ^x
}
```

- PJW(Peter J. Weinberger):  
基于 AT&T 贝尔实验室的 Peter J. Weinberger 的论文而发明的 hash 算法  

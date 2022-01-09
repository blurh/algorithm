# hash
- 概念:  
将任意长度的输入值转变成固定长度的值输出, 该值称为散列值, 输出值通常为字母与数字组合  
hash & 0x7FFFFFFF 把 hash 限制到 31 位
- 类型:   
整数 hash 函数常用方法有三种：直接取余法、乘积取整法、平方取中法  
字符串可以看成 256 进制(ANSI 字符串为 128 进制)的大整数  

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
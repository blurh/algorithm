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
h = h << 5 + h + c
```
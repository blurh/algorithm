# 堆(heap)
用数组实现的二叉树, 没有使用父指针或者子指针. 堆根据"堆属性"进行排序, "堆属性"决定了树中节点的位置  

- 用途  
构建优先队列  
支持堆排序  
快速找出一个集合中的最大/小值  

- 结构图
```
    100  
     └───┬───┐  
     ┊   19  36  
     ┊   │   └───────────┬───┐  
     ┊   └───────┬───┐   │   │  
     ┊   ┊   ┊   17  12  25  5  
     ┊   ┊   ┊   │   │   │   └───────────────────────────┬───┐  
     ┊   ┊   ┊   │   │   └───────────────────────┬───┐   │   │  
     ┊   ┊   ┊   │   └───────────────────┬───┐   │   │   │   │  
     ┊   ┊   ┊   └───────────────┬───┐   │   │   │   │   │   │  
     ┊   ┊   ┊   ┊   ┊   ┊   ┊   9   15  6   11  13  8   1   4  
     ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊   ┊  
  ┌────────────────────────────────────────────────────────────┐  
  │ 100  19  36  17  12  25  5   9   15  6   11  13  8   1   4 │  
  └────────────────────────────────────────────────────────────┘  
     0   1   2   3   4   5   6   7   8   9   10  11  12  13  14
        root
     1   2   3   4   5   6   7   8   9   10  11  12  13  14  15
    root

    root 节点为下标为 0 的情况下:   
      父节点索引: (i -1) / 2  
      左节点索引: 2 * i + 1  
      右节点索引: 2 * i + 2

    root 节点下标为 1 的情况下: 
      父节点索引: i / 2  
      左节点索引: 2 * i  
      右节点索引: 2 * i + 1
```

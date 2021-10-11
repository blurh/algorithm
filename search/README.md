# 查找算法
- 线性搜索算法(linear search):  
就是遍历搜索...

- 二分查找算法(binary search):  
时间复杂度: O(log n)

- 插值查找(interpolation search):   
对于数据分布均匀的数组查找效率高, 对于数据分布不均匀的效率反而不高  
二分查找查找中查找点:   
mid = (low+high)/2, 即: mid = low + 1/2*(high-low)  
将查找的点优化为自适应:   
mid = low + (value-a[low])/(a[high]-a[low])*(high-low)  

- 斐波那契查找(fibonacci search):  
查找点使用黄金分割点进行查找, 斐波那契查找的优点是它只涉及加法和减法运算, 除法比加减法要占用更多的时间, 斐波那契查找的运行时间理论上比二分法查找小  

- 树表查找(Tree table lookup):  
通过 AVLTree, RBTree, BTree 来查找, 略  

- 分块查找(block search):  
分块查找要求索引表有序的, 对块内节点无排序要求, 适合于节点动态变化的情况

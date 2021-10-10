# 查找算法
- 线性搜索算法(linear search):  
就是遍历搜索...

- 二分查找算法(binary search):  
时间复杂度: O(log n)

- 插值查找():   
对于数据分布均匀的数组查找效率高, 对于数据分布不均匀的效率反而不高  
二分查找查找中查找点:   
mid = (low+high)/2, 即: mid = low + 1/2*(high-low)  
将查找的点优化为自适应:   
mid = low + (value-a[low])/(a[high]-a[low])*(high-low)  

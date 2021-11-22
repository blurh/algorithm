package structure

import "tree"

type TreeMap struct {
    data    *tree.RBTree
    count    int
}

func InitTreeMap() *TreeMap {
    treeMap := new(TreeMap)
    treeMap.data = tree.InitRBTree()
    treeMap.count = 0
    return treeMap
}

func (treeMap *TreeMap) Set(key int, value interface{}) bool {
    if treeMap.Get(key) == nil {
        treeMap.count++
    }
    treeMap.data.Set(key, value)
    return true
}

func (treeMap *TreeMap) Get(key int) interface{} {
    return treeMap.data.Get(key)
}

func (treeMap *TreeMap) Remove(key int) bool {
    if treeMap.Get(key) == nil {
        return false
    }
    treeMap.data.Remove(key)
    treeMap.count--
    return true
}

func (treeMap *TreeMap) Keys() []int {
    return treeMap.data.Order()
}

func (treeMap *TreeMap) Values() []interface{} {
    values := []interface{}{}
    keys := treeMap.Keys()
    if len(keys) == 0 {
        return values
    }
    for _, key := range keys {
        values = append(values, treeMap.Get(key))
    }
    return values
}

func (treeMap *TreeMap) Count() int {
    return treeMap.count
}

func (treeMap *TreeMap) Clear() {
    treeMap.data = tree.InitRBTree()
    treeMap.count = 0
}

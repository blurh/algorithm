package structure

type Set struct {
    data *HashMap
    size int
}

func InitSet() *Set {
    hashmap := InitHashMap()
    set := new(Set)
    set.data = hashmap
    return set
}

func (set *Set) AddValue(key interface{}) bool {
    set.data.Set(key, struct{}{})
    return true
}

func (set *Set) RemoveValue(key interface{}) bool {
    set.data.Remove(key)
    return true
}

func (set *Set) Order() []interface{} {
    arr := set.data.Keys()
    return arr
}

func (set *Set) Clear() bool {
    set.data = InitHashMap()
    return true
}

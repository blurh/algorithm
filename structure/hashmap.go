package structure

// 需要存储 key 跟 value, 所以重新实现一个链表
type hashLink struct {
    key   int
    value interface{}
    next  *hashLink
}

type HashMap struct {
    size      int
    hashLinks []*hashLink
}

// ----------------- 链表 -------------------------
func initHashLink() *hashLink {
    l := new(hashLink)
    return l
}

func (link *hashLink) ExistsKey(key int) bool {
    for link != nil {
        if link.key == key {
            return true
        }
        link = link.next
    }
    return false
}

func (link *hashLink) GetLinkValue(key int) interface{} {
    for link != nil {
        if link.key == key {
            return link.value
        }
        link = link.next
    }
    return nil
}

func (link *hashLink) RemoveLinkKey(key int) bool {
    if !link.ExistsKey(key) {
        return false
    }
    var parentNode *hashLink
    for link != nil {
        if link.key != key {
            parentNode = link
            link = link.next
        } else {
            parentNode.next = link.next
            link = nil
            return true
        }
    }
    return false
}

func (link *hashLink) Append(key int, value interface{}) bool {
    for {
        if link.next == nil {
            link.next = &hashLink{key: key, value: value}
            break
        }
        link = link.next
    }
    return true
}

// ----------------------------------------------

func InitHashMap() *HashMap {
    hashMap := new(HashMap)
    hashMap.size = 0
    for i := 0; i <= 10; i++ {
        linkHead := initHashLink()
        hashMap.hashLinks = append(hashMap.hashLinks, linkHead)
    }
    return hashMap
}

// hash 函数
func getIndex(key int) int {
    return key % 10
}

func (hash *HashMap) Set(key int, value interface{}) bool {
    index := getIndex(key)
    link := hash.hashLinks[index]
    if link.ExistsKey(key) {
        for link.key != key {
            link = link.next
        }
        link.value = value
        return true
    }
    hash.size++
    link.Append(key, value)
    return true
}

func (hash *HashMap) Get(key int) interface{} {
    index := getIndex(key)
    link := hash.hashLinks[index]
    if !link.ExistsKey(key) {
        return nil
    }
    return link.GetLinkValue(key)
}

func (hash *HashMap) Remove(key int) bool {
    index := getIndex(key)
    link := hash.hashLinks[index]
    if !link.ExistsKey(key) {
        return false
    }
    hash.size--
    link.RemoveLinkKey(key)
    return true
}

func (hash *HashMap) Keys() []interface{} {
    var arr []interface{}
    for i := 0; i <= 9; i++ {
        link := hash.hashLinks[i]
        for link != nil {
            if link.value != nil {
                arr = append(arr, link.key)
            }
            link = link.next
        }
    }
    return arr
}

func (hash *HashMap) Values() []interface{} {
    var arr []interface{}
    for i := 0; i <= 9; i++ {
        link := hash.hashLinks[i]
        for link != nil {
            if link.value != nil {
                arr = append(arr, link.value)
            }
            link = link.next
        }
    }
    return arr
}

func (hash *HashMap) Size(key interface{}) int {
    return hash.size
}

func (hash *HashMap) Clear() bool {
    for i := 0; i <= 9; i++ {
        hash.hashLinks[i] = initHashLink()
    }
    hash.size = 0
    return true
}

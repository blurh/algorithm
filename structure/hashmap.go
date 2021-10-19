package structure

import "strings"

// 需要存储 key 跟 value, 所以重新实现一个链表
type hashLink struct {
    key   interface{}
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

func (link *hashLink) ExistsKey(key interface{}) bool {
    for link != nil {
        if link.key == key {
            return true
        }
        link = link.next
    }
    return false
}

func (link *hashLink) GetLinkValue(key interface{}) interface{} {
    for link != nil {
        if link.key == key {
            return link.value
        }
        link = link.next
    }
    return nil
}

func (link *hashLink) RemoveLinkKey(key interface{}) bool {
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

func (link *hashLink) Append(key interface{}, value interface{}) bool {
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

// 将字符串专为数字
func convertNum(data interface{}) (num int) {
    switch data.(type) {
    case int:
        num = data.(int)
    case string:
        letterArr := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
            "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
            "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D",
            "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
            "T", "U", "V", "W", "X", "Y", "Z"}
        stringSplit := strings.Split(data.(string), "")
        num = 0
        for _, v := range stringSplit {
            for i, letter := range letterArr {
                if letter == v {
                    num += i
                }
            }
        }
    default:
        num = 0
    }
    return
}

// hash 函数
func getIndex(key interface{}) int {
    return convertNum(key) % 10
}

func (hash *HashMap) Set(key interface{}, value interface{}) bool {
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

func (hash *HashMap) Get(key interface{}) interface{} {
    index := getIndex(key)
    link := hash.hashLinks[index]
    if !link.ExistsKey(key) {
        return nil
    }
    return link.GetLinkValue(key)
}

func (hash *HashMap) Remove(key interface{}) bool {
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

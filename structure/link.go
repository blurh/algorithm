package structure

import (
    "errors"
    "fmt"
)

type linkNode struct {
    data interface{}
    next *linkNode
}

type Link struct {
    length   int
    linkHead *linkNode
}

func InitLink() *Link {
    node := new(linkNode)
    node.data = "head"
    l := new(Link)
    l.length = 0
    l.linkHead = node
    return l
}

func (l *Link) Length() int {
    return l.length
}

func (l *Link) GetNode(index int) *linkNode {
    node := l.linkHead
    for i := 0; i < index; i++ {
        node = node.next
    }
    return node
}

func (l *Link) Get(index int) interface{} {
    node := l.GetNode(index)
    return node.data
}

func (l *Link) Insert(index int, data interface{}) error {
    if index <= 0 || index > l.length+1 {
        errMsg := fmt.Sprintf("index: %d is out of ranage: 0 ~ %d", index, l.Length()+1)
        return errors.New(errMsg)
    }
    node := new(linkNode)
    node.data = data
    preNode := l.GetNode(index - 1)
    node.next = preNode.next
    preNode.next = node
    l.length++
    return nil
}

func (l *Link) Append(value interface{}) error {
    endOfLink := l.Length() + 1
    l.Insert(endOfLink, value)
    return nil
}

func (l *Link) Remove(index int) error {
    preNode := l.GetNode(index - 1)
    node := l.GetNode(index)
    preNode.next = node.next
    node = nil
    l.length--
    return nil
}

func (l *Link) Set(index int, value interface{}) error {
    node := l.GetNode(index)
    node.data = value
    return nil
}

func (l *Link) Exists(value interface{}) bool {
    node := l.linkHead
    for node != nil {
        if node.data == value {
            return true
        }
        node = node.next
    }
    return false
}

func (l *Link) Order() []interface{} {
    var orderArr []interface{}
    node := l.linkHead.next
    for node.next != nil {
        orderArr = append(orderArr, node.data)
        node = node.next
    }
    return append(orderArr, node.data)
}

func (l *Link) Invert() bool {
    firstNode := l.linkHead.next
    node := firstNode
    var next, nextNext *linkNode
    for node != nil && node.next != nil {
        tmpNext := node.next
        node.next = next // 反转第一个节点指向
        next = tmpNext
        nextNext = next.next
        next.next = node // 反转第二个节点指向
        node = nextNext // 用于执行下一个循环
    }
    if node == nil {
        l.linkHead.next = next
    } else if node.next == nil {
        node.next = next // 反转
        l.linkHead.next = node
    }
    firstNode.next = nil
    return true
}

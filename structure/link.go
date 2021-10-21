package structure

import (
    "errors"
    "fmt"
)

type linkNode struct {
    data interface{}
    next *linkNode
}

type link struct {
    length   int
    linkHead *linkNode
}

func InitLink() *link {
    node := new(linkNode)
    node.data = "head"
    l := new(link)
    l.length = 0
    l.linkHead = node
    return l
}

func (l *link) GetLinkLength() int {
    return l.length
}

func (l *link) GetIndexNode(index int) (*linkNode, error) {
    node := l.linkHead
    for i := 0; i < index; i++ {
        node = node.next
    }
    return node, nil
}

func (l *link) GetIndexData(index int) (interface{}, error) {
    if index < 0 || index > l.GetLinkLength() {
        err := fmt.Sprintf("index: %d is out of ranage: 0 ~ %d", index, l.GetLinkLength()+1)
        errMsg := errors.New(err)
        return nil, errMsg
    }
    node, _ := l.GetIndexNode(index)
    return node.data, nil
}

func (l *link) InsertNode(index int, data interface{}) error {
    if index <= 0 || index > l.length+1 {
        errMsg := fmt.Sprintf("index: %d is out of ranage: 0 ~ %d", index, l.GetLinkLength()+1)
        return errors.New(errMsg)
    }
    node := new(linkNode)
    node.data = data
    // 获取到上一个节点, 让他的 next 指向插入的节点
    preNode, _ := l.GetIndexNode(index - 1)
    // 插入节点的 next 等于原节点的 next
    node.next = preNode.next
    // 原节点 next 指向插入的节点
    preNode.next = node
    l.length++
    return nil
}

func (l *link) AppendNode(value interface{}) error {
    endOfLink := l.GetLinkLength() + 1
    l.InsertNode(endOfLink, value)
    l.length++
    return nil
}

func (l *link) DeleteIndexNode(index int) error {
    preNode, _ := l.GetIndexNode(index - 1)
    node, _ := l.GetIndexNode(index)
    preNode.next = node.next
    node = nil
    l.length--
    return nil
}

func (l *link) SetIndexNodeData(index int, value interface{}) error {
    node, err := l.GetIndexNode(index)
    if err != nil {
        return err
    }
    node.data = value
    return nil
}

func (l *link) GetValue(value interface{}) bool {
    node := l.linkHead
    for node != nil {
        if node.data == value {
            return true
        }
        node = node.next
    }
    return false
}

package list

import (
	"fmt"
	"errors"
)

type Node struct{
    data    interface{}
    next    *Node
}

type List struct{
    length  int
    listHead    *Node
}

func initList() (*List) {
	node := new(Node)
	node.data = "head"
	l := new(List)
	l.length = 0
	l.listHead = node
    return l
}

func (l *List) GetListLength() (int) {
	return l.length
}

func (l *List) GetIndexNode(index int) (*Node, error) {
	node := l.listHead
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}

func (l *List) GetIndexData(index int) (interface{}, error) {
	if index < 0 || index > l.GetListLength() {
		err := fmt.Sprintf("index: %d is out of ranage: 0 ~ %d", index, l.GetListLength() + 1)
		errMsg := errors.New(err)
		return nil, errMsg
	}
	node, _ := l.GetIndexNode(index)
	return node.data, nil
}

func (l *List) InsertNode(index int, data interface{}) (error) {
	if index <= 0 || index > l.length + 1 {
		errMsg := fmt.Sprintf("index: %d is out of ranage: 0 ~ %d", index, l.GetListLength() + 1)
		return errors.New(errMsg)
    }
	node := new(Node)
	node.data = data
	// 获取到上一个节点, 让他的 next 指向插入的节点
	preNode, _ := l.GetIndexNode(index - 1)
	// 插入节点的 next 等于原节点的 next
	node.next = preNode.next
	// 原节点 next 指向插入的节点
	preNode.next = node
	l.length ++
	return nil
}

func (l *List) AppendNode(value interface{}) (error) {
	endOfList := l.GetListLength() + 1
	l.InsertNode(endOfList, value)
	l.length ++
	return nil
}

func (l *List) DeleteIndexNode(index int) (error) {
	preNode, _ := l.GetIndexNode(index - 1)
	node, _ := l.GetIndexNode(index)
	preNode.next = node.next
	node = nil
	l.length --
	return nil
}

func (l *List) SetIndexNodeData(index int, value interface{}) (error) {
	node, err := l.GetIndexNode(index)
	if err != nil {
		return err
	}
	node.data = value
	return nil
}

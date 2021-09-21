package structure

import (
	"errors"
	"fmt"
)

type Node struct {
	data      int
	leftNode  *Node
	rightNode *Node
}

func initBinaryTree() *Node {
	root := new(Node)
	return root
}

func (node *Node) AddLeft(data int) error {
	if node.leftNode != nil {
		errMsg := errors.New("this left node is already exists, add fail")
		return errMsg
	}
	node.leftNode = &Node{data: data}
	return nil
}

func (node *Node) AddRight(data int) error {
	if node.rightNode != nil {
		errMsg := errors.New("this right node is already exists, add fail")
		return errMsg
	}
	node.rightNode = &Node{data: data}
	return nil
}

func (node *Node) GetNodeData() int {
	return node.data
}

// DLR
func (node *Node) PreOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	node.leftNode.PreOrder()
	node.rightNode.PreOrder()
}

// LDR
func (node *Node) MiddleOrder() {
	if node == nil {
		return
	}
	node.leftNode.MiddleOrder()
	fmt.Println(node.data)
	node.rightNode.MiddleOrder()
}

// LRD
func (node *Node) PostOrder() {
	if node == nil {
		return
	}
	node.leftNode.PostOrder()
	node.rightNode.PostOrder()
	fmt.Println(node.data)
}

// 层次遍历(广度优先)

// 获取树高
func (node *Node) GetTreeHeight() int {
	if node == nil {
		return 0
	}
	leftHeight := node.leftNode.GetTreeHeight()
	rightHeight := node.rightNode.GetTreeHeight()
	if leftHeight >= rightHeight {
		leftHeight++
		return leftHeight
	} else {
		rightHeight++
		return rightHeight
	}
}

// 获取树的叶子节点个数
// 不计算根节点
func (node *Node) GetLeafNum() int {
	if node == nil {
		return 0
	}
	rightLeafNum := node.rightNode.GetLeafNum()
	leftLeafNum := node.leftNode.GetLeafNum()
	if node.leftNode != nil {
		leftLeafNum++
	}
	if node.rightNode != nil {
		rightLeafNum++
	}
	leafNum := leftLeafNum + rightLeafNum
	return leafNum
}

// 搜索

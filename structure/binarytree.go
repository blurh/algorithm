package structure

import "errors"

type treeNode struct {
    data      int
    leftNode  *treeNode
    rightNode *treeNode
}

func InitBinaryTree() *treeNode {
    root := new(treeNode)
    return root
}

func (node *treeNode) AddLeft(data int) error {
    if node.leftNode != nil {
        errMsg := errors.New("this left node is already exists, add fail")
        return errMsg
    }
    node.leftNode = &treeNode{data: data}
    return nil
}

func (node *treeNode) AddRight(data int) error {
    if node.rightNode != nil {
        errMsg := errors.New("this right node is already exists, add fail")
        return errMsg
    }
    node.rightNode = &treeNode{data: data}
    return nil
}

func (node *treeNode) GetNodeData() int {
    return node.data
}

// DLR
func (node *treeNode) PreOrder() []int {
    arr := []int{}
    if node == nil {
        return arr
    }
    nodeArr := []int{node.data}
    leftArr := node.leftNode.PreOrder()
    rightArr := node.rightNode.PreOrder()
    arr = append(nodeArr, append(leftArr, rightArr...)...)
    return arr
}

// LDR
func (node *treeNode) MiddleOrder() []int {
    arr := []int{}
    if node == nil {
        return arr
    }
    nodeArr := []int{node.data}
    leftArr := node.leftNode.MiddleOrder()
    rightArr := node.rightNode.MiddleOrder()
    arr = append(leftArr, append(nodeArr, rightArr...)...)
    return arr
}

// LRD
func (node *treeNode) PostOrder() []int {
    arr := []int{}
    if node == nil {
        return arr
    }
    nodeArr := []int{node.data}
    leftArr := node.leftNode.PostOrder()
    rightArr := node.rightNode.PostOrder()
    arr = append(leftArr, append(rightArr, nodeArr...)...)
    return arr
}

// 层次遍历(广度优先)
func (node *treeNode) BreadthFirstSearch() []int {
    if node == nil {
        return nil
    }
    nodes := []*treeNode{node}
    result := []int{}
    for len(nodes) > 0 {
        node := nodes[0]
        result = append(result, node.data)
        if node.leftNode != nil {
            nodes = append(nodes, node.leftNode)
        }
        if node.rightNode != nil {
            nodes = append(nodes, node.rightNode)
        }
        nodes = nodes[1:]
    }
    return result
}

// 获取树高
func (node *treeNode) GetTreeHeight() int {
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
func (node *treeNode) GetLeafNum() int {
    num := 0
    if node == nil {
        return num
    }
    if node.leftNode == nil && node.rightNode == nil {
        num++
    }
    leftNum := node.rightNode.GetLeafNum()
    rightNum := node.leftNode.GetLeafNum()
    num += leftNum + rightNum
    return num
}

// 搜索
func (node *treeNode) SearchValue(value int) bool {
    nodes := []*treeNode{node}
    for len(nodes) > 0 {
        node := nodes[0]
        data := node.data
        if data == value {
            return true
        }
        if node.leftNode != nil {
            nodes = append(nodes, node.leftNode)
        }
        if node.rightNode != nil {
            nodes = append(nodes, node.rightNode)
        }
        nodes = nodes[1:]
    }
    return false
}

// 翻转
func Invert(node *treeNode) *treeNode {
    if node == nil {
        return nil
    }
    node.leftNode, node.rightNode = node.rightNode, node.leftNode
    Invert(node.leftNode)
    Invert(node.rightNode)
    return node
}

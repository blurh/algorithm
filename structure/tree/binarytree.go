package tree

import "errors"

type treeNode struct {
    data      int
    leftNode  *treeNode
    rightNode *treeNode
}

func InitBinaryTree(value int) *treeNode {
    tree := new(treeNode)
    tree.data = value
    return tree
}

func (tree *treeNode) AddLeft(data int) error {
    if tree.leftNode != nil {
        errMsg := errors.New("this left node is already exists, add fail")
        return errMsg
    }
    tree.leftNode = &treeNode{data: data}
    return nil
}

func (tree *treeNode) AddRight(data int) error {
    if tree.rightNode != nil {
        errMsg := errors.New("this right node is already exists, add fail")
        return errMsg
    }
    tree.rightNode = &treeNode{data: data}
    return nil
}

func (tree *treeNode) GetNodeData() int {
    return tree.data
}

// DLR
func (tree *treeNode) PreOrder() []int {
    arr := []int{}
    if tree == nil {
        return arr
    }
    nodeArr := []int{tree.data}
    leftArr := tree.leftNode.PreOrder()
    rightArr := tree.rightNode.PreOrder()
    arr = append(nodeArr, append(leftArr, rightArr...)...)
    return arr
}

// LDR
func (tree *treeNode) MiddleOrder() []int {
    arr := []int{}
    if tree == nil {
        return arr
    }
    nodeArr := []int{tree.data}
    leftArr := tree.leftNode.MiddleOrder()
    rightArr := tree.rightNode.MiddleOrder()
    arr = append(leftArr, append(nodeArr, rightArr...)...)
    return arr
}

// LRD
func (tree *treeNode) PostOrder() []int {
    arr := []int{}
    if tree == nil {
        return arr
    }
    nodeArr := []int{tree.data}
    leftArr := tree.leftNode.PostOrder()
    rightArr := tree.rightNode.PostOrder()
    arr = append(leftArr, append(rightArr, nodeArr...)...)
    return arr
}

// 层次遍历(广度优先)
func (tree *treeNode) BreadthFirstSearch() []int {
    if tree == nil {
        return nil
    }
    nodes := []*treeNode{tree}
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

// 深度优先
func (t *treeNode) DepthFirstSearch() []int {
    depthStack := []*treeNode{t}
    result := []int{}
    for len(depthStack) > 0 {
        node := depthStack[0]
        depthStack = depthStack[1:]
        result = append(result, node.data)
        if node.rightNode != nil {
            depthStack = append([]*treeNode{node.rightNode}, depthStack...)
        }
        if node.leftNode != nil {
            depthStack = append([]*treeNode{node.leftNode}, depthStack...)
        }
    }
    return result
}

// 获取树高
func (tree *treeNode) GetTreeHeight() int {
    if tree == nil {
        return 0
    }
    leftHeight := tree.leftNode.GetTreeHeight()
    rightHeight := tree.rightNode.GetTreeHeight()
    if leftHeight >= rightHeight {
        leftHeight++
        return leftHeight
    } else {
        rightHeight++
        return rightHeight
    }
}

// 获取树的叶子节点个数
func (tree *treeNode) GetLeafNum() int {
    num := 0
    if tree == nil {
        return num
    }
    if tree.leftNode == nil && tree.rightNode == nil {
        num++
    }
    leftNum := tree.rightNode.GetLeafNum()
    rightNum := tree.leftNode.GetLeafNum()
    num += leftNum + rightNum
    return num
}

// 搜索
func (tree *treeNode) SearchValue(value int) bool {
    nodes := []*treeNode{tree}
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
func Invert(tree *treeNode) *treeNode {
    if tree == nil {
        return nil
    }
    tree.leftNode, tree.rightNode = tree.rightNode, tree.leftNode
    Invert(tree.leftNode)
    Invert(tree.rightNode)
    return tree
}

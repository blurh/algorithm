package tree

type searchTree struct {
    data      int
    leftNode  *searchTree
    rightNode *searchTree
}

func InitSearchTree(value int) *searchTree {
    tree := new(searchTree)
    tree.data = value
    return tree
}

func (tree *searchTree) SearchValue(value int) bool {
    if tree == nil {
        return false
    } else if tree.data == value {
        return true
    }
    if value > tree.data {
        tree = tree.rightNode
    } else {
        tree = tree.leftNode
    }
    searchRet := tree.SearchValue(value)
    return searchRet
}

func (tree *searchTree) InsertNode(value int) bool {
    valueExists := tree.SearchValue(value)
    if valueExists {
        return false
    }
    if value > tree.data {
        if tree.rightNode == nil {
            tree.rightNode = &searchTree{data: value}
            return true
        } else {
            tree = tree.rightNode
            tree.InsertNode(value)
        }
    } else {
        if tree.leftNode == nil {
            tree.leftNode = &searchTree{data: value}
            return true
        } else {
            tree = tree.leftNode
            tree.InsertNode(value)
        }
    }
    return true
}

func (tree *searchTree) MaxOfSearchTree() int {
    for tree.rightNode != nil {
        tree = tree.rightNode
    }
    return tree.data
}

func (tree *searchTree) MinOfSearchTree() int {
    for tree.leftNode != nil {
        tree = tree.leftNode
    }
    return tree.data
}

// DLR
func (tree *searchTree) PreOrder() []int {
    arr := []int{}
    if tree == nil {
        return arr
    }
    selfArr := []int{tree.data}
    leftArr := tree.leftNode.PreOrder()
    rightArr := tree.rightNode.PreOrder()
    arr = append(selfArr, append(leftArr, rightArr...)...)
    return arr
}

// LDR
func (tree *searchTree) MiddleOrder() []int {
    arr := []int{}
    if tree == nil {
        return nil
    }
    treeArr := []int{tree.data}
    leftArr := tree.leftNode.MiddleOrder()
    rightArr := tree.rightNode.MiddleOrder()
    arr = append(leftArr, append(treeArr, rightArr...)...)
    return arr
}

// LRD
func (tree *searchTree) PostOrder() []int {
    arr := []int{}
    if tree == nil {
        return nil
    }
    treeArr := []int{tree.data}
    leftArr := tree.leftNode.PostOrder()
    rightArr := tree.rightNode.PostOrder()
    arr = append(leftArr, append(rightArr, treeArr...)...)
    return arr
}

func (tree *searchTree) RemoveNode(value int) bool {
    _, removeResult := remove(tree, value)
    return removeResult
}

func remove(tree *searchTree, value int) (*searchTree, bool) {
    if tree == nil {
        return nil, false
    }
    var exists bool
    if value > tree.data {
        tree.rightNode, exists = remove(tree.rightNode, value)
        return tree, exists
    }
    if value < tree.data {
        tree.leftNode, exists = remove(tree.leftNode, value)
        return tree, exists
    }
    exists = true
    if tree.leftNode == nil && tree.rightNode == nil {
        tree = nil
        return tree, exists
    } else if tree.leftNode == nil {
        tree = tree.rightNode
        return tree, exists
    } else if tree.rightNode == nil {
        tree = tree.leftNode
        return tree, exists
    } else {
        // 左分支寻找最大值, 或右分支寻找最小值
        leftMaxValue := tree.leftNode.MaxOfSearchTree()
        tree.data = leftMaxValue
        tree.leftNode, _ = remove(tree.leftNode, leftMaxValue)
        return tree, exists
    }
}

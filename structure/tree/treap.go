package tree

type treap struct {
    data      int
    priority  int
    leftNode  *treap
    rightNode *treap
}

func InitTreap(value int, priority int) *treap {
    tree := new(treap)
    tree.data = value
    tree.priority = priority
    return tree
}

// LDR
func (tree *treap) Order() []int {
    if tree == nil {
        return nil
    }
    thisNodeArr := []int{tree.data}
    leftNodeArr := tree.leftNode.Order()
    rightNodeArr := tree.rightNode.Order()
    arr := append(leftNodeArr, append(thisNodeArr, rightNodeArr...)...)
    return arr
}

//
//        |                      |
//        a                      b
//       / \        ->          / \
//      b   c       <-         d   a
//     / \                        / \
//    d   e                      e   c
//
func (tree *treap) RightRotate() *treap {
    if tree == nil {
        return tree
    }
    root := tree.leftNode
    tree.leftNode = root.rightNode
    root.rightNode = tree
    return root
}

func (tree *treap) LeftRotate() *treap {
    if tree == nil {
        return tree
    }
    root := tree.rightNode
    tree.rightNode = root.leftNode
    root.leftNode = tree
    return root
}

func (tree *treap) MaxOfTreap() *treap {
    if tree == nil {
        return tree
    }
    for tree.rightNode != nil {
        tree = tree.rightNode
    }
    return tree
}

func (tree *treap) MinOfTreap() *treap {
    if tree == nil {
        return tree
    }
    for tree.leftNode != nil {
        tree = tree.leftNode
    }
    return tree
}

func (tree *treap) SearchNode(value int) *treap {
    if tree == nil {
        return tree
    }
    if value == tree.data {
        return tree
    }
    if value > tree.data {
        tree = tree.rightNode.SearchNode(value)
    } else {
        tree = tree.leftNode.SearchNode(value)
    }
    return tree
}

func (tree *treap) SearchValue(value int) bool {
    node := tree.SearchNode(value)
    if node == nil {
        return false
    }
    if node.data == value {
        return true
    }
    return false
}

func (tree *treap) InsertValue(value, priority int) *treap {
    if tree.SearchValue(value) {
        return tree
    }
    if tree == nil {
        tree = &treap{data: value, priority: priority}
    }
    if value > tree.data {
        if tree.rightNode == nil {
            tree.rightNode = &treap{data: value, priority: priority}
            if priority < tree.priority {
                tree = tree.LeftRotate()
            }
            return tree
        }
        tree.rightNode = tree.rightNode.InsertValue(value, priority)
        if tree.rightNode.priority < tree.priority {
            tree = tree.LeftRotate()
        }
        return tree
    } else if value < tree.data {
        if tree.leftNode == nil {
            tree.leftNode = &treap{data: value, priority: priority}
            if priority < tree.priority {
                tree = tree.RightRotate()
            }
            return tree
        }
        tree.leftNode = tree.leftNode.InsertValue(value, priority)
        if tree.leftNode.priority < tree.priority {
            tree = tree.RightRotate()
        }
        return tree
    }
    return tree
}

func (tree *treap) RemoveValue(value int) *treap {
    if !tree.SearchValue(value) {
        return tree
    }
    if value > tree.data {
        tree.rightNode = tree.rightNode.RemoveValue(value)
    } else if value < tree.data {
        tree.leftNode = tree.leftNode.RemoveValue(value)
    } else if value == tree.data {
        // 直接删除
        if tree.leftNode == nil && tree.rightNode == nil {
            tree = nil
            return tree
        }
        // 提升右节点
        if tree.leftNode == nil && tree.rightNode != nil {
            tree = tree.rightNode
            return tree
        }
        // 提升左节点
        if tree.leftNode != nil && tree.rightNode == nil {
            tree = tree.leftNode
            return tree
        }
        if tree.leftNode != nil && tree.rightNode != nil {
            if tree.leftNode.priority > tree.rightNode.priority {
                minOfRightNode := tree.rightNode.MinOfTreap()
                tree.data = minOfRightNode.data
                tree.rightNode = tree.rightNode.RemoveValue(minOfRightNode.data)
            } else {
                maxOfLeftNode := tree.leftNode.MaxOfTreap()
                tree.data = maxOfLeftNode.data
                tree.leftNode = tree.leftNode.RemoveValue(maxOfLeftNode.data)
            }
        }
    }
    return tree
}

func (tree *treap) CheckTreap() bool {
    if tree == nil {
        return true
    }
    if tree.leftNode == nil && tree.rightNode != nil {
        if tree.rightNode.data < tree.data {
            return false
        }
        if tree.rightNode.priority < tree.priority {
            return false
        }
    } else if tree.leftNode != nil && tree.leftNode == nil {
        if tree.leftNode.data > tree.data {
            return false
        }
        if tree.leftNode.priority < tree.priority {
            return false
        }
    } else if tree.leftNode != nil && tree.rightNode != nil {
        if tree.leftNode.data > tree.data || tree.data > tree.rightNode.data {
            return false
        }
        if tree.leftNode.priority < tree.priority || tree.priority > tree.rightNode.priority {
            return false
        }
    }
    checkLeftNode := tree.leftNode.CheckTreap()
    checkRightNode := tree.rightNode.CheckTreap()
    return checkLeftNode && checkRightNode
}

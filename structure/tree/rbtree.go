package tree

/*
  红黑树遵循的性质
  1. 所有叶子节点都是黑色节点(叶子是 NIL 节点)
  2. 每个红色节点必须有两个黑色的子节点 从每个叶子到根的所有路径上不能有两个连续的红色节点
  3. 从任意节点到其每个叶子节点的所有简单路径都包含相同数目的黑色节点
*/

const (
    RED   bool = true
    BLACK bool = false
)

type rbTree struct {
    data      int
    color     bool
    parent    *rbTree
    leftNode  *rbTree
    rightNode *rbTree
}

func newNode(value int, parent *rbTree) *rbTree {
    node := &rbTree{data: value, parent: parent, color: RED}
    return node
}

func InitRBTree(value int) *rbTree {
    tree := new(rbTree)
    tree.color = BLACK
    tree.data = value
    return tree
}

func (tree *rbTree) MaxOfRBTree() int {
    if tree == nil {
        return -1
    }
    for tree.rightNode != nil {
        tree = tree.rightNode
    }
    return tree.data
}

func (tree *rbTree) MinOfRBTree() int {
    if tree == nil {
        return -1
    }
    for tree.leftNode != nil {
        tree = tree.leftNode
    }
    return tree.data
}

// LDR
func (tree *rbTree) Order() []int {
    var arr []int
    if tree == nil {
        return arr
    }
    treeNodeArr := []int{tree.data}
    leftNodeArr := tree.leftNode.Order()
    rightNodeArr := tree.rightNode.Order()
    arr = append(leftNodeArr, append(treeNodeArr, rightNodeArr...)...)
    return arr
}

func (tree *rbTree) OrderLeaf() []int {
    var arr []int
    if tree == nil {
        return arr
    }
    treeNodeArr := []int{}
    if tree.leftNode == nil && tree.rightNode == nil {
        treeNodeArr = []int{tree.data}
    }
    leftNodeArr := tree.leftNode.OrderLeaf()
    rightNodeArr := tree.rightNode.OrderLeaf()
    arr = append(leftNodeArr, append(treeNodeArr, rightNodeArr...)...)
    return arr
}

func (tree *rbTree) SearchValue(value int) bool {
    if tree == nil {
        return false
    }
    if value == tree.data {
        return true
    }
    if value > tree.data {
        tree = tree.rightNode
    } else {
        tree = tree.leftNode
    }
    searchResult := tree.SearchValue(value)
    return searchResult
}

func (tree *rbTree) SearchNode(value int) *rbTree {
    if tree == nil {
        return nil
    }
    if value == tree.data {
        return tree
    }
    if value > tree.data {
        return tree.rightNode.SearchNode(value)
    } else {
        return tree.leftNode.SearchNode(value)
    }
    return tree
}

func (tree *rbTree) CheckBST() bool {
    if tree == nil {
        return true
    }
    if tree.leftNode != nil && tree.rightNode == nil {
        if tree.leftNode.data < tree.data {
            result := tree.leftNode.CheckBST()
            return result
        } else {
            return false
        }
    } else if tree.leftNode == nil && tree.rightNode != nil {
        if tree.rightNode.data > tree.data {
            result := tree.rightNode.CheckBST()
            return result
        } else {
            return false
        }
    } else if tree.leftNode != nil && tree.rightNode != nil {
        if tree.rightNode.data > tree.data && tree.data > tree.leftNode.data {
            leftResult := tree.leftNode.CheckBST()
            rightResult := tree.rightNode.CheckBST()
            return leftResult || rightResult
        }
    }
    return true
}

func (tree *rbTree) CheckBlackHeightBalance() bool {
    listOfLeaf := tree.OrderLeaf()
    height := 0
    blackHeight := 0
    for _, v := range listOfLeaf {
        for tree != nil {
            if v > tree.data {
                if tree.color == BLACK {
                    blackHeight++
                }
                tree = tree.rightNode
            } else if v < tree.data {
                if tree.color == BLACK {
                    blackHeight++
                }
                tree = tree.leftNode
            } else if v == tree.data {
                // 第一个 height 为 0 跳过不比较
                if height != 0 && height != blackHeight {
                    return false
                } else if height == 0 {
                    height = blackHeight
                }
                blackHeight = 0 // 重置累加的高度
                tree = nil      // 结束循环
            }
        }
    }
    return true
}

func (tree *rbTree) OrderCheckBlackHeightBalance() bool {
    if tree == nil {
        return true
    }
    treeCheck := tree.CheckBlackHeightBalance()
    leftNodeCheck := tree.leftNode.OrderCheckBlackHeightBalance()
    rightNodeCheck := tree.rightNode.OrderCheckBlackHeightBalance()
    return treeCheck && leftNodeCheck && rightNodeCheck
}

func (tree *rbTree) CheckRBTree() bool {
    checkBST := tree.CheckBST()
    checkBlackHeightBalance := tree.OrderCheckBlackHeightBalance()
    return checkBST && checkBlackHeightBalance
}

// 左旋: b a e 位置发生变化
//
//        |                    |
//        b                    a
//       / \                  / \
//      d   a       ->       b   c
//         / \              / \
//        e   c            d   e
//
func (tree *rbTree) LeftRotate() *rbTree {
    if tree == nil {
        return tree
    }
    root := tree.rightNode
    root.parent = tree.parent
    tree.rightNode = root.leftNode
    if tree.rightNode != nil {
        tree.rightNode.parent = tree
    }
    root.leftNode = tree
    tree.parent = root
    return root
}

// 右旋: b a e 位置发生变化
//
//             |                     |
//             a                     b
//            / \         ->        / \
//           b   c                 d   a
//          / \                       / \
//         d   e                     e   c
//
func (tree *rbTree) RightRotate() *rbTree {
    if tree == nil {
        return tree
    }
    root := tree.leftNode
    root.parent = tree.parent
    tree.leftNode = root.rightNode
    if tree.leftNode != nil {
        tree.leftNode.parent = tree
    }
    root.rightNode = tree
    tree.parent = root
    return root
}

func (tree *rbTree) LeftRightRotate() *rbTree {
    tree.leftNode = tree.leftNode.LeftRotate()
    tree = tree.RightRotate()
    return tree
}

func (tree *rbTree) RightLeftRotate() *rbTree {
    tree.rightNode = tree.rightNode.RightRotate()
    tree = tree.LeftRotate()
    return tree
}

func (tree *rbTree) InsertValue(value int) *rbTree {
    var caseValue int
    if exists := tree.SearchValue(value); exists {
        return tree
    }
    // case 1: 空树
    if tree.parent == nil && tree == nil {
        tree.data = value
        tree.color = BLACK
        caseValue = 1
        return tree
    }
    if value > tree.data {
        if tree.rightNode != nil {
            tree.rightNode = tree.rightNode.InsertValue(value)
            tree, caseValue = tree.Adjust(value)
        }
        // case 2: 插入节点的父节点为黑色, 直接插入即可
        if tree.rightNode == nil {
            tree.rightNode = newNode(value, tree)
            caseValue = 2
            return tree
        }
    } else if value < tree.data {
        if tree.leftNode != nil {
            tree.leftNode = tree.leftNode.InsertValue(value)
            tree, caseValue = tree.Adjust(value)
        }
        // case 2: 插入节点的父节点为黑色, 直接插入即可
        if tree.leftNode == nil {
            tree.leftNode = newNode(value, tree)
            caseValue = 2
            return tree
        }
    }
    // 回溯 修复 case 5 的变色
    if caseValue == 5 {
        node := tree.SearchNode(value)
        // tree.parent == nil 确保 tree 为 root, 即 InsertValue 的最外层调用
        for node.parent != nil && node.parent.parent != nil && node.parent != tree && tree.parent == nil {
            adjustValue := node.data
            node = node.parent.parent
            node, _ = node.Adjust(adjustValue)
        }
    }
    return tree
}

func (tree *rbTree) Adjust(value int) (retTree *rbTree, caseNum int) {
    // case 3: 插入节点的父节点为红, 叔节点黑/空, 子&父&祖父节点在同一直线上
    // 父节点涂黑, 祖父节点涂红, 祖父节点旋转
    if (tree.leftNode != nil && tree.leftNode.leftNode != nil && tree.leftNode.leftNode.data == value) &&
        tree.leftNode.color == RED && (tree.rightNode == nil || tree.rightNode.color == BLACK) {
        tree.leftNode.color = BLACK
        tree.color = RED
        tree = tree.RightRotate()
        return tree, 3
    } else if (tree.rightNode != nil && tree.rightNode.rightNode != nil && tree.rightNode.rightNode.data == value) &&
        tree.rightNode.color == RED && (tree.leftNode == nil || tree.leftNode.color == BLACK) {
        tree.rightNode.color = BLACK
        tree.color = RED
        tree = tree.LeftRotate()
        return tree, 3
    }
    // case 4: 插入节点的父节点为红, 叔节点黑/空, 子&父&祖父节点不在同一直线上
    // 插入节点涂黑, 祖父节点涂红, 父节点旋转, 祖父节点旋转
    if (tree.rightNode != nil && tree.rightNode.leftNode != nil && tree.rightNode.leftNode.data == value) &&
        tree.rightNode.color == RED && (tree.leftNode == nil || tree.leftNode.color == BLACK) {
        tree.color = RED
        tree.rightNode.leftNode.color = BLACK
        tree = tree.RightLeftRotate()
        return tree, 4
    } else if (tree.leftNode != nil && tree.leftNode.rightNode != nil && tree.leftNode.rightNode.data == value) &&
        tree.leftNode.color == RED && (tree.rightNode == nil || tree.rightNode.color == BLACK) {
        tree.color = RED
        tree.leftNode.rightNode.color = BLACK
        tree = tree.LeftRightRotate()
        return tree, 4
    }
    // case 5: 插入节点的父节点为红, 叔节点为红
    // 父&叔节点变黑, 祖父节点变红, 需要回溯: 将祖父节点看作新插入节点回溯调整
    if (tree.leftNode != nil && tree.leftNode.color == RED) && tree.rightNode.color == RED {
        // 不为根节点则涂红
        if tree.parent != nil {
            tree.color = RED
        }
        tree.leftNode.color = BLACK
        tree.rightNode.color = BLACK
        return tree, 5
    } else if (tree.rightNode != nil && tree.rightNode.color == RED) && tree.leftNode.color == RED {
        // 不为根节点则涂红
        if tree.parent != nil {
            tree.color = RED
        }
        tree.leftNode.color = BLACK
        tree.rightNode.color = BLACK
        return tree, 5
    }
    return tree, 0
}

func (tree *rbTree) isLeft() bool {
    if tree == nil && tree.parent != nil {
        return false
    }
    if tree == tree.parent.leftNode {
        return true
    }
    return false
}

func (tree *rbTree) isRight() bool {
    if tree == nil && tree.parent != nil {
        return false
    }
    if tree == tree.parent.rightNode {
        return true
    }
    return false
}

func (tree *rbTree) GetBrother() *rbTree {
    if tree == nil && tree.parent == nil {
        return nil
    }
    if tree.isLeft() {
        return tree.parent.rightNode
    } else if tree.isRight() {
        return tree.parent.leftNode
    }
    return nil
}

func (tree *rbTree) GetFarNephew() *rbTree {
    if tree == nil && tree.parent == nil {
        return nil
    }
    brother := tree.GetBrother()
    if brother == nil {
        return nil
    }
    if tree.isLeft() {
        return brother.rightNode
    } else if tree.isRight() {
        return brother.leftNode
    }
    return nil
}

func (tree *rbTree) GetNearNephew() *rbTree {
    if tree == nil && tree.parent == nil {
        return nil
    }
    brother := tree.GetBrother()
    if brother == nil {
        return nil
    }
    if tree.isLeft() {
        return brother.leftNode
    } else if tree.isRight() {
        return brother.rightNode
    }
    return nil
}

func (tree *rbTree) RemoveValue(value int) *rbTree {
    if !tree.SearchValue(value) {
        return tree
    }
    if tree == nil {
        return tree
    }
    if value > tree.data {
        tree.rightNode = tree.rightNode.RemoveValue(value)
    } else if value < tree.data {
        tree.leftNode = tree.leftNode.RemoveValue(value)
    } else if value == tree.data {
        // case 1: 节点红色, 且是叶子节点, 则父节点必定黑
        // 直接删除即可
        if tree.color == RED && tree.leftNode == nil && tree.rightNode == nil {
            tree = nil
            return tree
        }
        // case 2: 节点红色, 且仅有左子树或右子树
        //         不符合红黑树性质, 不存在
        // case 3: 节点红色, 且同时有左子树和右子树
        //         用删除节点的后继节点进行数据替换, 颜色不变. 删除后继节点
        if tree.color == RED && tree.leftNode != nil && tree.rightNode != nil {
            minValue := tree.rightNode.MinOfRBTree()
            tree.data = minValue
            tree.rightNode = tree.rightNode.RemoveValue(minValue)
            return tree
        }
        // case 4: 节点黑色, 且仅有左子树或右子树
        //         提升子节点, 涂黑
        if tree.color == BLACK && (tree.leftNode == nil && tree.rightNode != nil) {
            tree = tree.rightNode
            tree.color = BLACK
            return tree
        } else if tree.color == BLACK && (tree.leftNode != nil && tree.rightNode == nil) {
            tree = tree.leftNode
            tree.color = BLACK
            return tree
        }
        // case 5: 节点黑色, 且为叶子节点
        if tree.color == BLACK && tree.leftNode == nil && tree.rightNode == nil {
            // 1) 删除节点为父的左节点, 兄节点红
            //    交换父&兄节点颜色, 父节点左旋, 然后再调用一次删除方法即可
            if tree.isLeft() && tree.GetBrother() != nil && tree.GetBrother().color == RED {
                tree.parent.color, tree.GetBrother().color = tree.GetBrother().color, tree.parent.color
                tree.parent = tree.parent.LeftRotate()
                tree = tree.RemoveValue(value)
                return tree
            }
            // 2) 删除节点为父的右节点, 兄节点红
            //    交换父&兄节点颜色, 父节点右旋, 然后再重新调用一次删除方法
            if tree.isRight() && tree.GetBrother() != nil && tree.GetBrother().color == RED {
                tree.parent.color, tree.GetBrother().color = tree.GetBrother().color, tree.parent.color
                tree.parent = tree.parent.RightRotate()
                tree = tree.RemoveValue(value)
                return tree
            }
            // 3) 删除节点为父的左节点, 兄节点黑, 远侄子节点红
            //    交换父&兄节点颜色, 父节点左旋, 远侄子节点涂黑, 删除节点
            if tree.isLeft() && tree.GetBrother() != nil && tree.GetBrother().color == BLACK &&
                tree.GetFarNephew() != nil && tree.GetFarNephew().color == RED {
                tree.parent.color, tree.GetBrother().color = tree.GetBrother().color, tree.parent.color
                tree.GetFarNephew().color = BLACK
                tree.parent = tree.parent.LeftRotate()
                tree = tree.RemoveValue(value)
                return tree
            }
            // 4) 删除节点为父的右节点, 兄节点黑, 远侄子节点红
            //    交换父&兄节点颜色, 父节点右旋, 远侄子节点涂黑, 删除节点
            if tree.isRight() && tree.GetBrother() != nil && tree.GetBrother().color == BLACK &&
                tree.GetFarNephew() != nil && tree.GetFarNephew().color == RED {
                tree.parent.color, tree.GetBrother().color = tree.GetBrother().color, tree.parent.color
                tree.GetFarNephew().color = BLACK
                tree.parent = tree.parent.RightRotate()
                tree = tree.RemoveValue(value)
                return tree
            }
            // 5) 删除节点为左孩子, 兄节点为黑, 远侄子节点为黑, 近侄子节点为红
            //    交换兄&近侄子节点颜色, 兄节点右旋, 变成 case 5-3 的情况, 继续删除即可
            if tree.isLeft() && tree.GetBrother() != nil && tree.GetBrother().color == BLACK &&
                tree.GetFarNephew() != nil && tree.GetFarNephew().color == BLACK &&
                tree.GetNearNephew() != nil && tree.GetNearNephew().color == RED {
                tree.GetBrother().color, tree.GetNearNephew().color = tree.GetNearNephew().color, tree.GetBrother().color
                tree.parent.rightNode = tree.parent.rightNode.RightRotate()
                tree = tree.RemoveValue(value)
                return tree
            }
            // 6) 删除节点为右孩子, 兄节点为黑, 远侄子节点为黑, 近侄子节点为红
            //    交换兄&近侄子节点颜色, 兄节点左旋, 变成 case 5-4 的情况, 继续执行删除
            if tree.isRight() && tree.GetBrother() != nil && tree.GetBrother().color == BLACK &&
                tree.GetFarNephew() != nil && tree.GetFarNephew().color == BLACK &&
                tree.GetNearNephew() != nil && tree.GetNearNephew().color == RED {
                tree.GetBrother().color, tree.GetNearNephew().color = tree.GetNearNephew().color, tree.GetBrother().color
                tree.parent.leftNode = tree.parent.leftNode.LeftRotate()
                tree = tree.RemoveValue(value)
                return tree
            }
            // 7) 父节点红, 兄节点&侄子(只能是 nil 节点)节点均黑
            //    父节点涂黑, 兄节点涂红, 删除节点
            if tree.parent.color == RED && tree.GetBrother() != nil && tree.GetBrother().color == BLACK &&
                tree.GetFarNephew() == nil && tree.GetNearNephew() == nil {
                tree.parent.color = BLACK
                tree.GetBrother().color = RED
                tree = tree.RemoveValue(value)
                return tree
            }
            // 8) 父&兄&侄子节点均为黑
            //    删除节点, 兄节点涂红, 父节点回溯调整
            if tree.parent.color == BLACK && tree.GetBrother() != nil && tree.GetBrother().color == BLACK &&
                tree.GetFarNephew() == nil && tree.GetNearNephew() == nil {
                tree.GetBrother().color = RED
                tree = tree.RemoveValue(value)
                if tree.parent != nil && tree.parent.parent != nil && tree.parent.parent.parent != nil {
                    tree.parent.parent.parent, _ = tree.parent.parent.parent.Adjust(tree.parent.data)
                }
                return tree
            }
        }
    }
    return tree
}

func (tree *rbTree) Clear() *rbTree {
    tree = nil
    return tree
}

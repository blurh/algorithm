package structure

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

var leaf = &rbTree{color: BLACK}

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
        tmpNode := tree.SearchNode(value)
        // 确保 tree 为 root(tree.parent == nil), 即最外层的循环
        for tmpNode.parent != nil && tmpNode.parent.parent != nil &&
            tmpNode.parent != tree && tree.parent == nil {
            tmpValue := tmpNode.data
            tmpNode = tmpNode.parent.parent
            tmpNode, _ = tmpNode.Adjust(tmpValue)
        }
    }
    return tree
}

func (tree *rbTree) Adjust(value int) (retTree *rbTree, caseValue int) {
    // case 3: 插入节点的父节点为红, 叔节点黑/空, 子&父&祖父节点在同一直线上
    // 父节点涂黑, 祖父节点涂红, 右旋
    if (tree.leftNode != nil && tree.leftNode.leftNode != nil && tree.leftNode.leftNode.data == value) && tree.leftNode.color == RED &&
        (tree.rightNode == nil || tree.rightNode.color == BLACK) {
        tree.leftNode.color = BLACK
        tree.color = RED
        tree = tree.RightRotate()
        return tree, 3
    } else if (tree.rightNode != nil && tree.rightNode.rightNode != nil && tree.rightNode.rightNode.data == value) && tree.rightNode.color == RED &&
        (tree.leftNode == nil || tree.leftNode.color == BLACK) {
        tree.rightNode.color = BLACK
        tree.color = RED
        tree = tree.LeftRotate()
        return tree, 3
    }
    // case 4: 插入节点的父节点为红, 叔节点黑/空, 子&父&祖父节点不在同一直线上
    // 插入节点涂黑, 祖父节点涂红, 右旋
    if (tree.rightNode.leftNode != nil && tree.rightNode.leftNode.data == value) && tree.rightNode.color == RED &&
        (tree.leftNode == nil || tree.leftNode.color == BLACK) {
        tree.color = RED
        child := tree.SearchNode(value)
        child.color = BLACK
        tree = tree.RightLeftRotate()
        return tree, 4
    } else if (tree.leftNode.rightNode != nil && tree.leftNode.rightNode.data == value) && tree.leftNode.color == RED &&
        (tree.rightNode == nil || tree.rightNode.color == BLACK) {
        tree.color = RED
        child := tree.SearchNode(value)
        child.color = BLACK
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

func (tree *rbTree) RemoveValue(value int) bool {
    // TODO
    return true
}

func (tree *rbTree) Clear() *rbTree {
    tree = nil
    return tree
}

package tree

import "fmt"

const (
    RED   bool = true
    BLACK bool = false
)

type rbTree struct {
    root  *rbTreeNode
    count int
}

func InitRBTree() *rbTree {
    tree := new(rbTree)
    tree.root = nil
    tree.count = 0
    return tree
}

func (tree *rbTree) Clear() bool {
    tree.root = nil
    tree.count = 0
    return true
}

func (tree *rbTree) Max() int {
    if tree.root == nil {
        return -1
    }
    return tree.root.MaxOfRBTree().index
}

func (tree *rbTree) Min() int {
    if tree.root == nil {
        return -1
    }
    return tree.root.MinOfRBTree().index
}

func (tree *rbTree) Order() []int {
    if tree.root == nil {
        return []int{}
    }
    return tree.root.Order()
}

func (tree *rbTree) Get(index int) interface{} {
    if tree.root == nil {
        return nil
    }
    return tree.root.Get(index)
}

func (tree *rbTree) Update(index int, value interface{}) bool {
    if tree == nil {
        return false
    }
    searchNode := tree.root.GetNode(index)
    if searchNode == nil {
        return false
    }
    searchNode.value = value
    return true
}

func (tree *rbTree) Set(index int, value interface{}) int {
    if tree.Get(index) == nil {
        tree.Insert(index, value)
        return 1
    }
    tree.Update(index, value)
    return 2
}

func (tree *rbTree) CheckRBTree() int {
    if tree.root == nil {
        return 0
    }
    if tree.root.parent != nil || tree.root.color == RED {
        return 1
    }
    if len(tree.Order()) != tree.count {
        return 2
    }
    if !tree.root.CheckBST() {
        return 3
    }
    if !tree.root.CheckDoubleRed() {
        return 4
    }
    if !tree.root.OrderCheckBlackHeightBalance() {
        return 5
    }
    return 0
}

func (tree *rbTree) Insert(index int, value interface{}) bool {
    if tree.Get(index) != nil {
        return false
    }
    if tree.root == nil {
        tree.root = InitRBTreeNode(index, value)
    } else {
        tree.root = tree.root.Insert(index, value) 
        tree.root = tree.root.Adjust(index)
    }
    tree.count++
    return true
}

func (tree *rbTree) Remove(index int) bool {
    // search 已经判断了 root 的 nil 了, 故不需要再判断
    if tree.Get(index) == nil {
        return false
    }
    tree.root = tree.root.Remove(index)
    tree.count--
    return true
}

func (tree *rbTree) Count() int {
    return tree.count
}

type rbTreeNode struct {
    index     int
    value     interface{}
    color     bool
    parent    *rbTreeNode
    leftNode  *rbTreeNode
    rightNode *rbTreeNode
}

func newNode(index int, value interface{}, parent *rbTreeNode) *rbTreeNode {
    node := &rbTreeNode{index: index, value: value, parent: parent, color: RED}
    return node
}

func InitRBTreeNode(index int, value interface{}) *rbTreeNode {
    node := new(rbTreeNode)
    node.index = index
    node.value = value 
    node.color = BLACK
    return node
}

func (node *rbTreeNode) MaxOfRBTree() *rbTreeNode {
    if node == nil {
        return nil
    }
    for node.rightNode != nil {
        node = node.rightNode
    }
    return node
}

func (node *rbTreeNode) MinOfRBTree() *rbTreeNode {
    if node == nil {
        return nil
    }
    for node.leftNode != nil {
        node = node.leftNode
    }
    return node
}

// LDR
func (node *rbTreeNode) Order() []int {
    var arr []int
    if node == nil {
        return arr
    }
    nodeArr := []int{node.index}
    leftNodeArr := node.leftNode.Order()
    rightNodeArr := node.rightNode.Order()
    arr = append(leftNodeArr, append(nodeArr, rightNodeArr...)...)
    return arr
}

func (node *rbTreeNode) OrderLeaf() []int {
    var arr []int
    if node == nil {
        return arr
    }
    nodeArr := []int{}
    if node.leftNode == nil && node.rightNode == nil {
        nodeArr = []int{node.index}
    }
    leftNodeArr := node.leftNode.OrderLeaf()
    rightNodeArr := node.rightNode.OrderLeaf()
    arr = append(leftNodeArr, append(nodeArr, rightNodeArr...)...)
    return arr
}

func (node *rbTreeNode) GetNode(index int) *rbTreeNode {
    if node == nil {
        return nil
    }
    if index == node.index {
        return node
    } else if index > node.index {
        return node.rightNode.GetNode(index)
    } else {
        return node.leftNode.GetNode(index)
    }
}

func (node *rbTreeNode) Get(index int) interface{} {
    if node == nil {
        return nil
    }
    searchNode := node.GetNode(index)
    if searchNode == nil {
        return nil
    }
    return searchNode.value
}

// ---------------------- this for test -----------------------

func (node *rbTreeNode) CheckBST() bool {
    if node == nil {
        return true
    }
    leftCheckResult, rightCheckResult := true, true
    if node.leftNode != nil {
        if node.leftNode.index < node.index {
            leftCheckResult = node.leftNode.CheckBST()
        } else {
            return false
        }
    }
    if node.rightNode != nil {
        if node.rightNode.index > node.index {
            rightCheckResult = node.rightNode.CheckBST()
        } else {
            return false
        }
    }
    return leftCheckResult && rightCheckResult
}

func (node *rbTreeNode) CheckDoubleRed() bool {
    if node == nil {
        return true
    } 
    leftCheckResult, rightCheckResult := true, true
    if node.leftNode != nil {
        if  node.color == RED && node.leftNode.color == RED {
            return false
        } else {
            leftCheckResult = node.leftNode.CheckDoubleRed()
        }
    }
    if node.rightNode != nil {
        if node.color == RED && node.rightNode.color == RED {
            return false
        } else {
            rightCheckResult = node.rightNode.CheckDoubleRed()
        }
    }
    return  leftCheckResult && rightCheckResult
}

func (node *rbTreeNode) CheckBlackHeightBalance() bool {
    listOfLeaf := node.OrderLeaf()
    height := 0
    blackHeight := 0
    for _, v := range listOfLeaf {
        for node != nil {
            if v > node.index {
                if node.color == BLACK {
                    blackHeight++
                }
                node = node.rightNode
            } else if v < node.index {
                if node.color == BLACK {
                    blackHeight++
                }
                node = node.leftNode
            } else if v == node.index {
                // 第一个 height 为 0 跳过不比较
                if height != 0 && height != blackHeight {
                    return false
                } else if height == 0 {
                    height = blackHeight
                }
                blackHeight = 0 // 重置累加的高度
                node = nil      // 结束循环
            }
        }
    }
    return true
}

func (node *rbTreeNode) OrderCheckBlackHeightBalance() bool {
    if node == nil {
        return true
    }
    nodeCheck := node.CheckBlackHeightBalance()
    leftNodeCheck := node.leftNode.OrderCheckBlackHeightBalance()
    rightNodeCheck := node.rightNode.OrderCheckBlackHeightBalance()
    return nodeCheck && leftNodeCheck && rightNodeCheck
}

// -------------------------------------------------------------

func (node *rbTreeNode) isLeft() bool {
    if node == nil || node.parent == nil {
        return false
    }
    if node == node.parent.leftNode {
        return true
    }
    return false
}

func (node *rbTreeNode) isRight() bool {
    if node == nil || node.parent == nil {
        return false
    }
    if node == node.parent.rightNode {
        return true
    }
    return false
}

func (node *rbTreeNode) GetBrother() *rbTreeNode {
    if node == nil || node.parent == nil {
        return nil
    }
    if node.isLeft() {
        return node.parent.rightNode
    } else if node.isRight() {
        return node.parent.leftNode
    }
    return nil
}

func (node *rbTreeNode) GetFarNephew() *rbTreeNode {
    if node == nil || node.parent == nil {
        return nil
    }
    brother := node.GetBrother()
    if brother == nil {
        return nil
    }
    if node.isLeft() {
        return brother.rightNode
    } else if node.isRight() {
        return brother.leftNode
    }
    return nil
}

func (node *rbTreeNode) GetNearNephew() *rbTreeNode {
    if node == nil || node.parent == nil {
        return nil
    }
    brother := node.GetBrother()
    if brother == nil {
        return nil
    }
    if node.isLeft() {
        return brother.leftNode
    } else if node.isRight() {
        return brother.rightNode
    }
    return nil
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
func (node *rbTreeNode) LeftRotate() *rbTreeNode {
    if node == nil {
        return node
    }
    root := node.rightNode
    root.parent = node.parent
    if node.isLeft() {
        root.parent.leftNode = root
    } else if node.isRight() {
        root.parent.rightNode = root
    }
    node.rightNode = root.leftNode
    if node.rightNode != nil {
        node.rightNode.parent = node
    }
    root.leftNode = node
    node.parent = root
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
func (node *rbTreeNode) RightRotate() *rbTreeNode {
    if node == nil {
        return node
    }
    root := node.leftNode
    root.parent = node.parent
    if node.isLeft() {
        root.parent.leftNode = root
    } else if node.isRight() {
        root.parent.rightNode = root
    }
    node.leftNode = root.rightNode
    if node.leftNode != nil {
        node.leftNode.parent = node
    }
    root.rightNode = node
    node.parent = root
    return root
}

func (node *rbTreeNode) LeftRightRotate() *rbTreeNode {
    node.leftNode = node.leftNode.LeftRotate()
    node = node.RightRotate()
    return node
}

func (node *rbTreeNode) RightLeftRotate() *rbTreeNode {
    node.rightNode = node.rightNode.RightRotate()
    node = node.LeftRotate()
    return node
}

func (root *rbTreeNode) Adjust(index int) *rbTreeNode {
    rootNode := root
    var node *rbTreeNode
    // 
    adjustNode := root.GetNode(index)
    if adjustNode != nil && adjustNode.parent != nil && adjustNode.parent.parent != nil {
        node = adjustNode.parent.parent
    } else {
        return root
    }
    // case 1: 空树, 直接插入, 不需要处理
    // case 2: 插入节点的父节点为黑, 不需要处理

    // case 3: 插入节点的父节点为红, 叔节点黑/空, 子&父&祖父节点在同一直线上
    // 父节点涂黑, 祖父节点涂红, 祖父节点旋转
    if (node.leftNode != nil && node.leftNode.leftNode != nil && node.leftNode.leftNode.index == index) &&
        node.leftNode.color == RED && (node.rightNode == nil || node.rightNode.color == BLACK) {
        fmt.Println("adjust case 3")
        node.leftNode.color = BLACK
        node.color = RED
        if node == rootNode {
            root = node.RightRotate()
        } else {
            node = node.RightRotate()
        }
    } else if (node.rightNode != nil && node.rightNode.rightNode != nil && node.rightNode.rightNode.index == index) &&
        node.rightNode.color == RED && (node.leftNode == nil || node.leftNode.color == BLACK) {
        node.rightNode.color = BLACK
        node.color = RED
        if node == rootNode {
            root = node.LeftRotate()
        } else {
            node = node.LeftRotate()
        }
    }
    // case 4: 插入节点的父节点为红, 叔节点黑/空, 子&父&祖父节点不在同一直线上
    // 插入节点涂黑, 祖父节点涂红, 父节点旋转, 祖父节点旋转
    if (node.rightNode != nil && node.rightNode.leftNode != nil && node.rightNode.leftNode.index == index) &&
        node.rightNode.color == RED && (node.leftNode == nil || node.leftNode.color == BLACK) {
        fmt.Println("adjust case 4")
        node.color = RED
        node.rightNode.leftNode.color = BLACK
        if rootNode == node {
            root = node.RightLeftRotate()
        } else {
            node = node.RightLeftRotate()
        }
    } else if (node.leftNode != nil && node.leftNode.rightNode != nil && node.leftNode.rightNode.index == index) &&
        node.leftNode.color == RED && (node.rightNode == nil || node.rightNode.color == BLACK) {
        node.color = RED
        node.leftNode.rightNode.color = BLACK
        if rootNode == node {
            root = node.LeftRightRotate()
        } else {
            node = node.LeftRightRotate()
        }
    }
    // case 5: 插入节点的父节点为红, 叔节点为红
    // 父&叔节点变黑, 祖父节点变红, 需要回溯: 将祖父节点看作新插入节点回溯调整
    if (node.leftNode != nil && node.leftNode.color == RED) && (node.rightNode != nil && node.rightNode.color == RED) {
        // 不为根节点则涂红
        fmt.Println("adjust case 5")
        if node != rootNode {
            node.color = RED
        }
        node.leftNode.color = BLACK
        node.rightNode.color = BLACK
        // 回溯
        root = root.Adjust(node.index)
    } else if (node.rightNode != nil && node.rightNode.color == RED) && node.leftNode.color == RED {
        // 不为根节点则涂红
        fmt.Println("adjust case 5")
        if node != rootNode {
            node.color = RED
        }
        node.leftNode.color = BLACK
        node.rightNode.color = BLACK
        // 回溯
        root = root.Adjust(node.index)
    }
    return root
}

func (node *rbTreeNode) Insert(index int, value interface{}) *rbTreeNode {
    if node.Get(index) != nil {
        return node
    }
    // case 1: 空树
    if node == nil && node.parent == nil {
        node = newNode(index, value, nil)
        node.color = BLACK
    } else if index > node.index {
        if node.rightNode != nil {
            node.rightNode = node.rightNode.Insert(index, value)
        } else {
            node.rightNode = newNode(index, value, node)
        }
    } else if index < node.index {
        if node.leftNode != nil {
            node.leftNode = node.leftNode.Insert(index, value)
        } else {
            node.leftNode = newNode(index, value, node)
        }
    }
    return node
}

// 删除的动作在单分支节点和叶子节点进行
// 对于有左右子树的节点, 删除节点指删除了其后继节点/前驱节点
func (root *rbTreeNode) Remove(index int) *rbTreeNode {
    if root == nil || root.Get(index) == nil {
        return root
    }
    node := root.GetNode(index)
    fmt.Println(node)
    // case 1: 节点红色, 且是叶子节点, 则父节点必定黑
    //         直接删除即可
    if node.color == RED && node.leftNode == nil && node.rightNode == nil {
        fmt.Println("remove case 1")
        node = nil
    } 
    // case 2: 节点红色, 且仅有左子树或右子树
    //         不符合红黑树性质, 不存在
    // ---
    // case 3: 节点红色, 且同时有左子树和右子树
    //         用删除节点的后继节点进行数据替换, 颜色不变. 删除后继节点
    if node.color == RED && node.leftNode != nil && node.rightNode != nil {
        fmt.Println("remove case 3")
        minIndexNode := node.rightNode.MinOfRBTree()
        node.index = minIndexNode.index
        node.value = minIndexNode.value
        root = root.Remove(minIndexNode.index)
    } 
    // case 4: 节点黑色, 且仅有左子树或右子树
    //         提升子节点, 涂黑
    if node.color == BLACK && (node.leftNode == nil && node.rightNode != nil) {
        fmt.Println("remove case 4")
        node = node.rightNode
        node.color = BLACK
    } else if node.color == BLACK && (node.leftNode != nil && node.rightNode == nil) {
        fmt.Println("remove case 4")
        node = node.leftNode
        node.color = BLACK
    } 
    // case 5: 节点黑色, 且为叶子节点
    if node.color == BLACK && node.leftNode == nil && node.rightNode == nil {
        fmt.Println("remove case 5")
        // 1) 删除节点为 父节点的左节点, 兄节点红
        //    交换父&兄节点颜色, 父节点左旋, 然后再调用一次删除方法即可
        if node.isLeft() && node.GetBrother() != nil && node.GetBrother().color == RED {
            fmt.Println("remove case 5-1")
            node.parent.color, node.GetBrother().color = node.GetBrother().color, node.parent.color
            node.parent = node.parent.LeftRotate()
            root = root.Remove(index)
        } 
        // 2) 删除节点为 父的右节点, 兄节点红
        //    交换父&兄节点颜色, 父节点右旋, 然后再重新调用一次删除方法
        if node.isRight() && node.GetBrother() != nil && node.GetBrother().color == RED {
            fmt.Println("remove case 5-2")
            node.parent.color, node.GetBrother().color = node.GetBrother().color, node.parent.color
            node.parent = node.parent.RightRotate()
            root = root.Remove(index)
        } 
        // 3) 删除节点为父的左节点, 兄节点黑, 远侄子节点红
        //    交换父&兄节点颜色, 父节点左旋, 远侄子节点涂黑, 删除节点
        if node.isLeft() && node.GetBrother() != nil && node.GetBrother().color == BLACK &&
            node.GetFarNephew() != nil && node.GetFarNephew().color == RED {
            fmt.Println("remove case 5-3")
            node.parent.color, node.GetBrother().color = node.GetBrother().color, node.parent.color
            node.GetFarNephew().color = BLACK
            node.parent = node.parent.LeftRotate()
            root = root.Remove(index)
        } 
        // 4) 删除节点为父的右节点, 兄节点黑, 远侄子节点红
        //    交换父&兄节点颜色, 父节点右旋, 远侄子节点涂黑, 删除节点
        if node.isRight() && node.GetBrother() != nil && node.GetBrother().color == BLACK && 
            node.GetFarNephew() != nil && node.GetFarNephew().color == RED {
            fmt.Println("remove case 5-4")
            node.parent.color, node.GetBrother().color = node.GetBrother().color, node.parent.color
            node.GetFarNephew().color = BLACK
            node.parent = node.parent.RightRotate()
            root = root.Remove(index)
        } 
        // 5) 删除节点为左孩子, 兄节点为黑, 远侄子节点为黑, 近侄子节点为红
        //    交换兄&近侄子节点颜色, 兄节点右旋, 变成 case 5-3 的情况, 继续删除即可
        if node.isLeft() && node.GetBrother() != nil && node.GetBrother().color == BLACK &&
            node.GetFarNephew() != nil && node.GetFarNephew().color == BLACK &&
            node.GetNearNephew() != nil && node.GetNearNephew().color == RED {
            fmt.Println("remove case 5-5")
            node.GetBrother().color, node.GetNearNephew().color = node.GetNearNephew().color, node.GetBrother().color
            node.parent.rightNode = node.parent.rightNode.RightRotate()
            root = root.Remove(index)
        }  
        // 6) 删除节点为右孩子, 兄节点为黑, 远侄子节点为黑, 近侄子节点为红
        //    交换兄&近侄子节点颜色, 兄节点左旋, 变成 case 5-4 的情况, 继续执行删除
        if node.isRight() && node.GetBrother() != nil && node.GetBrother().color == BLACK &&
            node.GetFarNephew() != nil && node.GetFarNephew().color == BLACK &&
            node.GetNearNephew() != nil && node.GetNearNephew().color == RED {
            fmt.Println("remove case 5-6")
            node.GetBrother().color, node.GetNearNephew().color = node.GetNearNephew().color, node.GetBrother().color
            node.parent.leftNode = node.parent.leftNode.LeftRotate()
            root = root.Remove(index)
        } 
        // 7) 父节点红, 兄节点&侄子(只能是 nil 节点)节点均黑
        //    父节点涂黑, 兄节点涂红, 删除节点
        if node.parent.color == RED && node.GetBrother() != nil && node.GetBrother().color == BLACK &&
            node.GetFarNephew() == nil && node.GetNearNephew() == nil {
            fmt.Println("remove case 5-7")
            node.parent.color = BLACK
            node.GetBrother().color = RED
            root = root.Remove(index)
        } 
        // 8) 父&兄&侄子节点均为黑
        //    删除节点, 兄节点涂红, 父节点回溯调整
        if node.parent.color == BLACK && node.GetBrother() != nil && node.GetBrother().color == BLACK &&
            node.GetFarNephew() == nil && node.GetNearNephew() == nil {
            fmt.Println("remove case 5-8")
            node.GetBrother().color = RED
            parent := node.parent
            node = nil
            root = root.Adjust(parent.index)
        }
    }
    return root
}

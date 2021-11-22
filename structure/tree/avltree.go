package tree

type AvlTree struct {
    root    *avlTreeNode
    count   int
}

type avlTreeNode struct {
    index     int
    value     interface{}
    height    int
    leftNode  *avlTreeNode
    rightNode *avlTreeNode
}

func initAvlTreeNode(index int, value interface{}) *avlTreeNode {
    node := new(avlTreeNode)
    node.index = index
    node.value = value
    node.height = 0
    return node
}

func InitAvlTree() *AvlTree {
    tree := new(AvlTree)
    // tree.root = initAvlTreeNode()
    tree.count = 0
    return tree
}

func (tree *AvlTree) Get(index int) interface{} {
    return tree.root.Get(index)
}

func (tree *AvlTree) Insert(index int, value interface{}) bool {
    if tree.root == nil {
        tree.root = initAvlTreeNode(index, value)
        return true
    } 
    tree.root = tree.root.Insert(index, value)
    return true
}

func (tree *AvlTree) Remove(index int) bool {
    tree.root = tree.root.Remove(index)
    return true
}

func (tree *AvlTree) Max() int {
    return tree.root.MaxOfAVLTree()
}

func (tree *AvlTree) Min() int {
    return tree.root.MinOfAVLTree()
}

func (tree *AvlTree) Count() int {
    return tree.count
}

func (tree *AvlTree) Keys() []int {
    return tree.root.Order()
}

func (tree *AvlTree) Values() []interface{} {
    values := []interface{}{}
    keys := tree.Keys()
    if len(keys) == 0 {
        return values
    }
    for _, key := range keys {
        values = append(values, tree.Get(key))
    }
    return values
}

func (tree *AvlTree) Clear() {
    tree.root = nil
    tree.count = 0
}

func (node *avlTreeNode) GetHeight() int {
    if node == nil {
        return 0
    }
    return node.height
}

func (node *avlTreeNode) Get(index int) interface{} {
    if node == nil {
        return nil
    }
    if index == node.index {
        return node.value
    } else if index > node.index {
        node = node.rightNode
    } else if index < node.index {
        node = node.leftNode
    }
    return node.Get(index)
}

// LDR
func (node *avlTreeNode) Order() []int {
    arr := []int{}
    if node == nil {
        return arr
    }
    nodeNodeArr := []int{node.index}
    leftNodeArr := node.leftNode.Order()
    rightNodeArr := node.rightNode.Order()
    arr = append(leftNodeArr, append(nodeNodeArr, rightNodeArr...)...)
    return arr
}

func (node *avlTreeNode) MaxOfAVLTree() int {
    for node.rightNode != nil {
        node = node.rightNode
    }
    return node.index
}

func (node *avlTreeNode) MinOfAVLTree() int {
    for node.leftNode != nil {
        node = node.leftNode
    }
    return node.index
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

// LL 右单旋转
//                       5
//         5            /           3
//        /    ->     3     ->    /   \
//      3            /           2     5
//                  2
//
//        5                    5 (失衡节点)        3
//      /   \                /   \              /    \
//     3     6     ->       3     6     ->     2      5
//   /   \                /   \              /      /   \
//  2     4              2     4            1      4     6
//                     /
//                    1
// 由上面可知, 3 5 4 三个节点位置发生变化, 3 5 高度发生变化
func (node *avlTreeNode) LLRotate() *avlTreeNode {
    root := node.leftNode
    node.leftNode = root.rightNode
    root.rightNode = node
    // adjust height
    // 用 GetHeight 的原因是 nil 的话也返回 0
    // root.rightNode.height = max(root.rightNode.leftNode.GetHeight(), root.rightNode.rightNode.GetHeight())
    node.height = max(node.leftNode.GetHeight(), node.rightNode.GetHeight()) + 1
    root.height = max(root.leftNode.GetHeight(), root.rightNode.GetHeight()) + 1
    return root
}

// RR 左单旋转
//                      5
//     5                 \               7
//      \        ->       7      ->     /  \
//       7                 \           5    8
//                          8
//
//        5                    5 (失衡节点)         7
//      /   \                /   \               /    \
//     3     7       ->     3     7      ->     5      8
//         /   \                /   \         /   \      \
//        6     8              6     8       3     6      9
//                                     \
//                                      9
// 位置发生变化的为: 7 5 6, 高度发生变化的为: 7 5
func (node *avlTreeNode) RRRotate() *avlTreeNode {
    root := node.rightNode
    node.rightNode = root.leftNode
    root.leftNode = node
    // adjust height
    // root.leftNode.height = max(root.leftNode.leftNode.GetHeight(), root.leftNode.rightNode.GetHeight())
    node.height = max(node.leftNode.GetHeight(), node.rightNode.GetHeight()) + 1
    root.height = max(root.leftNode.GetHeight(), root.rightNode.GetHeight()) + 1
    return root
}

// LR 先左后右
//                    5              5
//      5            /              /             4
//     /     ->     2      ->     4       ->    /   \
//    2              \           /             2     5
//                    4         2
//
//          5                    5                    5                   3
//        /   \                /   \                /   \               /   \
//       2     6     ->       2     6     ->       3     6     ->      2     5
//     /   \                /   \                /   \                /    /   \
//    1     3              1     3              2     4             1     4     6
//                                 \          /
//                                  4        1
//
//          5                    5                    5                   4
//        /   \                /   \                /   \               /   \
//       3     6     ->       3     6     ->       4     6     ->      3     5
//     /   \                /   \                /                   /   \     \
//    1     4              1     4              3                   1     2     6
//                             /              /   \
//                            2              1     2
func (node *avlTreeNode) LRRotate() *avlTreeNode {
    node.leftNode = node.leftNode.RRRotate()
    node = node.LLRotate()
    return node
}

// RL 先右后左
//                     5             5
//       5              \             \              6
//        \     ->       7     ->      6      ->    / \
//         7            /               \          5   7
//                     6                 7
//
//        4                    4 (失衡节点)         4                      6
//      /   \                /   \               /    \                 /    \
//     3     7       ->     3     7      ->     3      6       ->      4      7
//         /   \                /   \                /   \           /   \      \
//        6     9              6     9              5     7         3     5      9
//                           /                             \
//                          5                               9
//
//        5                    5 (失衡节点)         5                      6
//      /   \                /   \               /    \                 /    \
//     3     8       ->     3     8      ->     3      6       ->      5      8
//         /   \                /   \                    \           /      /   \
//        6     9              6     9                    8         3      7     9
//                               \                      /   \
//                                7                    7     9
func (node *avlTreeNode) RLRotate() *avlTreeNode {
    node.rightNode = node.rightNode.LLRotate()
    node = node.RRRotate()
    return node
}

func (node *avlTreeNode) Adjust() *avlTreeNode {
    // 判断树的形状进行相应的旋转
    // 左边高则需要右旋(LL) (LR)
    // 右边高则需要左旋(RR) (RL)
    if node.leftNode.GetHeight()-node.rightNode.GetHeight() == 2 {
        if node.leftNode.leftNode.GetHeight() > node.leftNode.rightNode.GetHeight() {
            node = node.LLRotate()
        } else {
            node = node.LRRotate()
        }
    } else if node.leftNode.GetHeight()-node.rightNode.GetHeight() == -2 {
        if node.rightNode.leftNode.GetHeight() < node.rightNode.rightNode.GetHeight() {
            node = node.RRRotate()
        } else {
            node = node.RLRotate()
        }
    }
    return node
}

func (node *avlTreeNode) Insert(index int, value interface{}) *avlTreeNode {
    if node.Get(index) != nil {
        return node
    }
    if node == nil {
        node = &avlTreeNode{index: index, value:value, height: 1}
        return node
    }
    if index < node.index {
        node.leftNode = node.leftNode.Insert(index, value)
        node = node.Adjust()
    } else {
        node.rightNode = node.rightNode.Insert(index, value)
        node = node.Adjust()
    }
    node.height = max(node.leftNode.GetHeight(), node.rightNode.GetHeight()) + 1
    return node
}

func (node *avlTreeNode) Remove(index int) *avlTreeNode {
    if node.Get(index) == nil {
        return node
    }
    if node == nil {
        return nil
    }
    if index < node.index {
        node.leftNode = node.leftNode.Remove(index)
    } else if index > node.index {
        node.rightNode = node.rightNode.Remove(index)
    } else if index == node.index {
        // 左右节点都为空, 直接删除
        // 右节点为空, 左节点不为空, 直接提升左节点
        // 左节点为空, 右节点不为空, 直接提升右节点
        // 左右节点不为空, 从左分支查找最大值进行提升
        if node.leftNode == nil && node.rightNode == nil {
            node = nil
        } else if node.leftNode != nil && node.rightNode == nil {
            node = node.leftNode
        } else if node.leftNode == nil && node.rightNode != nil {
            node = node.rightNode
        } else if node.leftNode != nil && node.rightNode != nil {
            if node.leftNode.GetHeight() > node.rightNode.GetHeight() {
                maxIndexOfLeft := node.leftNode.MaxOfAVLTree()
                node.index = maxIndexOfLeft
                node.leftNode = node.leftNode.Remove(maxIndexOfLeft)
            } else {
                minIndexOfRight := node.rightNode.MinOfAVLTree()
                node.index = minIndexOfRight
                node.rightNode = node.rightNode.Remove(minIndexOfRight)
            }
        }
    }
    // 自平衡及更新高度
    if node != nil {
        node.height = max(node.leftNode.GetHeight(), node.rightNode.GetHeight()) + 1
        node = node.Adjust()
    }
    return node
}

// ------------------------------- for test -----------------------------

func (node *avlTreeNode) CheckAVLTree() bool {
    if node == nil {
        return true
    }
    checkBalance := node.CheckBalance()
    checkBST := node.CheckBST()
    return checkBalance && checkBST
}

func (node *avlTreeNode) CheckBST() bool {
    if node == nil {
        return true
    }
    if node.leftNode != nil && node.rightNode == nil {
        if node.leftNode.index < node.index {
            result := node.leftNode.CheckBST()
            return result
        } else {
            return false
        }
    } else if node.leftNode == nil && node.rightNode != nil {
        if node.rightNode.index > node.index {
            result := node.rightNode.CheckBST()
            return result
        } else {
            return false
        }
    } else if node.leftNode != nil && node.rightNode != nil {
        if node.rightNode.index > node.index && node.index > node.leftNode.index {
            leftResult := node.leftNode.CheckBST()
            rightResult := node.rightNode.CheckBST()
            return leftResult || rightResult
        }
    }
    return true
}

func (node *avlTreeNode) GetAVLTreeHeight() int {
    if node == nil {
        return 0
    }
    leftTreeHeight := node.leftNode.GetAVLTreeHeight()
    rightTreeHeight := node.rightNode.GetAVLTreeHeight()
    if leftTreeHeight > rightTreeHeight {
        leftTreeHeight++
        return leftTreeHeight
    } else {
        rightTreeHeight++
        return rightTreeHeight
    }
}

func (node *avlTreeNode) CheckBalance() bool {
    if node == nil {
        return true
    }
    leftTreeHeight := node.leftNode.GetAVLTreeHeight()
    rightTreeHeight := node.rightNode.GetAVLTreeHeight()
    if leftTreeHeight-rightTreeHeight > 1 || leftTreeHeight-rightTreeHeight < -1 {
        return false
    }
    checkLeftResult := node.leftNode.CheckBalance()
    checkRightResult := node.rightNode.CheckBalance()
    return checkLeftResult || checkRightResult
}

// --------------------------------------------------------------------------------
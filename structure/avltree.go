package structure

type avlTree struct {
    data      int
    height    int
    leftNode  *avlTree
    rightNode *avlTree
}

func initAvlTree() *avlTree {
    tree := new(avlTree)
    tree.height = 1
    return tree
}

func (tree *avlTree) GetHeight() int {
    if tree == nil {
        return 0
    }
    return tree.height
}

func (tree *avlTree) SearchValue(value int) bool {
    if tree == nil {
        return false
    }
    if value == tree.data {
        return true
    } else if value > tree.data {
        tree = tree.rightNode
    } else if value < tree.data {
        tree = tree.leftNode
    }
    searchResult := tree.SearchValue(value)
    return searchResult
}

// LDR
func (tree *avlTree) MiddleOrder() []int {
    arr := []int{}
    if tree == nil {
        return arr
    }
    treeNodeArr := []int{tree.data}
    leftNodeArr := tree.leftNode.MiddleOrder()
    rightNodeArr := tree.rightNode.MiddleOrder()
    arr = append(leftNodeArr, append(treeNodeArr, rightNodeArr...)...)
    return arr
}

// ------------------------------- 用于测试 -----------------------------

func (tree *avlTree) CheckAVLTree() bool {
    if tree == nil {
        return true
    }
    checkBalance := tree.CheckBalance()
    checkBST := tree.CheckBST()
    return checkBalance || checkBST
}

func (tree *avlTree) CheckBST() bool {
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

func (tree *avlTree) GetAVLTreeHeight() int {
    if tree == nil {
        return 0
    }
    leftTreeHeight := tree.leftNode.GetAVLTreeHeight()
    rightTreeHeight := tree.rightNode.GetAVLTreeHeight()
    if leftTreeHeight > rightTreeHeight {
        leftTreeHeight++
        return leftTreeHeight
    } else {
        rightTreeHeight++
        return rightTreeHeight
    }
}

func (tree *avlTree) CheckBalance() bool {
    if tree == nil {
        return true
    }
    leftTreeHeight := tree.leftNode.GetAVLTreeHeight()
    rightTreeHeight := tree.rightNode.GetAVLTreeHeight()
    if leftTreeHeight-rightTreeHeight > 1 || leftTreeHeight-rightTreeHeight < -1 {
        return false
    }
    checkLeftResult := tree.leftNode.CheckBalance()
    checkRightResult := tree.rightNode.CheckBalance()
    return checkLeftResult || checkRightResult
}

// --------------------------------------------------------------------------------

func (tree *avlTree) MaxOfAVLTree() int {
    if tree.rightNode != nil {
        tree = tree.rightNode
    }
    return tree.data
}

func (tree *avlTree) MinOfAVLTree() int {
    if tree.leftNode != nil {
        tree = tree.leftNode
    }
    return tree.data
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
func (tree *avlTree) LLRotate() *avlTree {
    root := tree.leftNode
    tree.leftNode = root.rightNode
    root.rightNode = tree
    // adjust height
    // 用 GetHeight 的原因是 nil 的话也返回 0
    // root.rightNode.height = max(root.rightNode.leftNode.GetHeight(), root.rightNode.rightNode.GetHeight())
    tree.height = max(tree.leftNode.GetHeight(), tree.rightNode.GetHeight())
    root.height = max(root.leftNode.GetHeight(), root.rightNode.GetHeight())
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
func (tree *avlTree) RRRotate() *avlTree {
    root := tree.rightNode
    tree.rightNode = root.leftNode
    root.leftNode = tree
    // adjust height
    // root.leftNode.height = max(root.leftNode.leftNode.GetHeight(), root.leftNode.rightNode.GetHeight())
    tree.height = max(tree.leftNode.GetHeight(), tree.rightNode.GetHeight())
    root.height = max(root.leftNode.GetHeight(), root.rightNode.GetHeight())
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
func (tree *avlTree) LRRotate() *avlTree {
    tree.leftNode = tree.leftNode.RRRotate()
    tree = tree.LLRotate()
    return tree
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
func (tree *avlTree) RLRotate() *avlTree {
    tree.rightNode = tree.rightNode.LLRotate()
    tree = tree.RRRotate()
    return tree
}

// 找到最小不平衡树进行旋转
func (tree *avlTree) Adjust() *avlTree {
    // 判断树的形状进行相应的旋转
    // 左边高则需要右旋(LL) (LR)
    // 右边高则需要左旋(RR) (RL)
    if tree.leftNode.GetHeight()-tree.rightNode.GetHeight() == 2 {
        if tree.leftNode.leftNode.GetHeight() > tree.leftNode.rightNode.GetHeight() {
            tree = tree.LLRotate()
        } else {
            tree = tree.LRRotate()
        }
    } else if tree.leftNode.GetHeight()-tree.rightNode.GetHeight() == -2 {
        if tree.rightNode.leftNode.GetHeight() < tree.rightNode.rightNode.GetHeight() {
            tree = tree.RRRotate()
        } else {
            tree = tree.RLRotate()
        }
    }
    return tree
}

func (tree *avlTree) InsertValue(value int) *avlTree {
    if exists := tree.SearchValue(value); exists == true {
        return tree
    }
    if tree == nil {
        tree = &avlTree{data: value, height: 1}
        return tree
    }
    if value < tree.data {
        tree.leftNode = tree.leftNode.InsertValue(value)
        tree = tree.Adjust()
    } else {
        tree.rightNode = tree.rightNode.InsertValue(value)
        tree = tree.Adjust()
    }
    tree.height = max(tree.leftNode.GetHeight(), tree.rightNode.GetHeight()) + 1
    return tree
}

func (tree *avlTree) RemoveValue(value int) *avlTree {
    if tree == nil {
        return nil
    }
    if value < tree.data {
        tree.leftNode = tree.leftNode.RemoveValue(value)
    } else if value > tree.data {
        tree.rightNode = tree.rightNode.RemoveValue(value)
    } else if value == tree.data {
        // 左右节点都为空, 直接删除
        // 右节点为空, 左节点不为空, 直接提升左节点
        // 左节点为空, 右节点不为空, 直接题生右节点
        // 左右节点不为空, 从左分支查找最大值进行提升
        if tree.leftNode == nil && tree.rightNode == nil {
            tree = nil
        } else if tree.leftNode != nil && tree.rightNode == nil {
            tree = tree.leftNode
        } else if tree.leftNode == nil && tree.rightNode != nil {
            tree = tree.rightNode
        } else if tree.leftNode != nil && tree.rightNode != nil {
            maxValueOfLeft := tree.leftNode.MaxOfAVLTree()
            tree.data = maxValueOfLeft
            tree.leftNode = tree.leftNode.RemoveValue(maxValueOfLeft)
        }
        tree = tree.Adjust()
    }
    return tree
}

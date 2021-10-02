package structure

// TODO:
// - remove value

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

func (tree *avlTree) CheckAvlTree() bool {
	// TODO
	return true
}

func (tree *avlTree) CheckBalance() bool {
	// TODO
	return true
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

func (tree *avlTree) RemoveValue() bool {
	// TODO
	return false
}

package structure

// TODO:
// - height
// - RR LL LR RL
// - insert value
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

func (tree *avlTree) GetTreeHeight() int {
	if tree == nil {
		return 0
	}
	treeLeftHeight := tree.leftNode.GetTreeHeight()
	treeRgihtHeight := tree.rightNode.GetTreeHeight()
	if treeLeftHeight > treeRgihtHeight {
		treeLeftHeight++
		return treeLeftHeight
	} else {
		treeRgihtHeight++
		return treeRgihtHeight
	}
}

func (tree *avlTree) GetBF() int {
	if tree == nil {
		return 0
	}
	treeLeftHeight := tree.leftNode.GetTreeHeight()
	treeRgihtHeight := tree.rightNode.GetTreeHeight()
	return treeLeftHeight - treeRgihtHeight
}

func (tree *avlTree) CheckBF() bool {
	checkRelust := checkBF(tree)
	return checkRelust
}

func checkBF(tree *avlTree) bool {
	if tree == nil {
		return true
	}
	treeNodeBF := tree.GetBF()
	treeNodeBFCheck := true
	if treeNodeBF > 1 || treeNodeBF < -1 {
		treeNodeBFCheck = false
	}
	leftNodeBFCheck := checkBF(tree.leftNode)
	rightNodeBFCheck := checkBF(tree.rightNode)
	bfCheck := treeNodeBFCheck && leftNodeBFCheck && rightNodeBFCheck
	return bfCheck
}

func (tree *avlTree) CheckAvlTree() bool {
	if tree == nil {
		return true
	}
	bfCheck := tree.CheckBF()
	if !bfCheck {
		return false
	}
	middleArr := tree.MiddleOrder()
	lastValue := 0
	for _, v := range middleArr {
		if v > lastValue {
			lastValue = v
		} else {
			return false
		}
	}
	return true
}

// LL(指的是树结构) 右单旋转
//           5          5
//         /           /           3
//        3    ->    3     ->    /   \
//                  /           2     5
//                 2
//
//        5                    5 (失衡节点)        3
//      /   \                /   \              /    \
//     3     6     ->       3     6    ->      2      5
//   /   \                /   \              /      /   \
//  2     4              2     4            1      4     6
//                     /
//                    1
func (tree *avlTree) LL() *avlTree {
	// TODO
	root := tree.leftNode
	tmpNode := root.rightNode
	root.rightNode = tree
	root.rightNode.leftNode = tmpNode
	return root
}

// RR 左单旋转
//    5                 5
//     \                 \               7
//      7        ->       7      ->     /  \
//                         \           5    8
//                          8
//
//        5                    5 (失衡节点)         7
//      /   \                /   \               /    \
//     3     7       ->     3     7      ->     5      8
//         /   \                /   \         /   \      \
//        6     8              6     8       3     6      9
//                                     \
//                                      9
func (tree *avlTree) RR() *avlTree {
	// TODO
	root := tree.rightNode
	tmpNode := root.leftNode
	root.leftNode = tree
	root.leftNode.rightNode = tmpNode
	return root
}

// LR 先左后右
//       5            5              5
//     /             /              /             4
//    2      ->     2      ->     4       ->    /   \
//                   \           /             2     5
//                    4         2
//
//          5                    5                    5                   4
//        /   \                /   \                /   \               /   \
//       2     6     ->       2     6     ->       4     6     ->      2     5
//     /   \                /   \                /   \                /    /   \
//    1     4              1     4              2     3             1     3     6
//                                 \          /
//                                  3        1
//
//          5                    5                    5                   4
//        /   \                /   \                /   \               /   \
//       3     6     ->       3     6     ->       4     6     ->      3     5
//     /   \                /   \                /                   /   \     \
//    1     4              1     4              3                   1     2     6
//                             /              /   \
//                            2              1     2
func (tree *avlTree) LR() *avlTree {
	// TODO:
	return tree
}

// RL 先右后左
//      5              5             5
//       \              \             \              6
//        7     ->       7     ->      6      ->    / \
//                      /               \          5   7
//                     6                 7
//
//        5                    5 (失衡节点)         5                      6
//      /   \                /   \               /    \                 /    \
//     3     7       ->     3     7      ->     3      6       ->      5      7
//         /   \                /   \                /   \           /      /   \
//        6     9              6     9              4     7         3      4     9
//                           /                             \
//                          4                               9
//
//        5                    5 (失衡节点)         5                      6
//      /   \                /   \               /    \                 /    \
//     3     7       ->     3     8      ->     3      6       ->      5      8
//         /   \                /   \                    \           /      /   \
//        6     9              6     9                    8         3      7     9
//                               \                      /   \
//                                7                    7     9
func (tree *avlTree) RL() *avlTree {
	// TODO
	return tree
}

func (tree *avlTree) InsertValue() bool {
	// TODO
	if tree == nil {
		return true
	}
	return false
}

func (tree *avlTree) RemoveValue() bool {
	// TODO
	return false
}

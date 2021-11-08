package tree

import "fmt"

// 定义树的阶
const M int = 5

// 非 root 节点至少应有的 data 个数
const Min int = M / 2

type BTree struct {
    root  *BTreeNode
    count int
}

func InitBTree() *BTree {
    tree := new(BTree)
    tree.count = 0
    return tree
}

func (tree *BTree) GetCount() int {
    return tree.count
}

func (tree *BTree) Search(index int) interface{} {
    root := tree.root
    return root.Search(index)
}

func (tree *BTree) Insert(index int, value interface{}) bool {
    tree.root = tree.root.Insert(index, value)
    tree.count++
    return true
}

func (tree *BTree) Delete(index int) bool {
    if tree.Search(index) == nil {
        return false
    }
    tree.root = tree.root.Delete(index)
    tree.count--
    return true
}

func (tree *BTree) Update(index int, value interface{}) bool {
    indexNode := tree.root.SearchNode(index)
    if indexNode != nil {
        for i := 0; i <= len(indexNode.data)-1; i++ {
            if indexNode.data[i].index == index {
                indexNode.data[i].value = value
                return true
            }
        }
    }
    return false
}

func (tree *BTree) Order() []int {
    return tree.root.Order()
}

func (tree *BTree) MinIndexOfTree() *item {
    root := tree.root
    return root.MinIndex()
}

func (tree *BTree) MaxIndexOfTree() *item {
    root := tree.root
    return root.MaxIndex()
}

func (tree *BTree) CheckBTree() int {
    if len(tree.Order()) != tree.GetCount() {
        return 1
    }
    if tree.root.parent != nil {
        return 2
    }
    if !tree.root.CheckBTree() {
        return 3
    }
    return 0
}

type item struct {
    index int
    value interface{}
}

type BTreeNode struct {
    data   []*item      // M - 1
    child  []*BTreeNode // M
    parent *BTreeNode
}

func (node *BTreeNode) SearchNode(index int) *BTreeNode {
    if node == nil {
        return nil
    }
    for i := 0; i <= len(node.data)-1; i++ {
        if index == node.data[i].index {
            return node
        } else if index < node.data[0].index && len(node.child) != 0 {
            return node.child[0].SearchNode(index)
        } else if index >= node.data[i].index && (i != len(node.data)-1 && index < node.data[i+1].index) && len(node.child) != 0 {
            return node.child[i+1].SearchNode(index)
        } else if i == len(node.data)-1 && len(node.child) != 0 {
            return node.child[len(node.child)-1].SearchNode(index)
        }
    }
    return nil
}

func (node *BTreeNode) Search(index int) interface{} {
    if node == nil {
        return nil
    }
    nodeSearchResult := node.SearchNode(index)
    if nodeSearchResult == nil {
        return nil
    }
    for i := 0; i <= len(nodeSearchResult.data)-1; i++ {
        if index == nodeSearchResult.data[i].index {
            return nodeSearchResult.data[i].value
        }
    }
    return nil
}

// root 节点分裂时, 树高才会变高
//                 root                            root                          root
//                  |                               |                             |
//           [ 19 20 34 50 ]     ->         [ 19 20 34 40 50 ]     ->           [ 34 ]
//                                                                               /  \
//                                                                      [ 19 20 ]    [ 40 50 ]
//
//                [ 60 70 ]                       [ 60 70 ]                  [ 34 60 70 ]
//                /   |   \      ->               /   |   \     ->          /    |   \   \
//   [ 19 20 34 50 ]              [ 19 20 34 40 50 ]               [ 19 20 ] [ 40 50 ]
//    /  /  |  \  \                                                 /  |  \   /  |  \
//
//                [ 19 50 ]                  [ 19 50 ]                  [ 19 22 50 ]
//                /   |   \      ->          /   |   \        ->        /   /  \   \
//                    |                          |                         /    \
//             [ 20 21 22 34 ]           [ 20 21 22 34 40 ]          [ 20 21 ] [ 34 40 ]
func (node *BTreeNode) Split() *BTreeNode {
    if node == nil {
        return node
    }
    parent := node.parent
    if len(node.data) == M {
        // 上浮的 data
        indexOfSift := len(node.data) / 2
        siftData := node.data[indexOfSift]
        // 分裂 child // 深拷贝
        nodeNewLeftChild := append([]*BTreeNode{}, node.child[:len(node.child)/2]...)
        nodeNewRightChild := append([]*BTreeNode{}, node.child[len(node.child)/2:]...)
        // 分裂的 node 节点 // 深拷贝
        leftData := append([]*item{}, node.data[:indexOfSift]...)
        rightData := append([]*item{}, node.data[indexOfSift+1:]...)
        newLeftNode := &BTreeNode{parent: parent, data: leftData, child: nodeNewLeftChild}
        newRightNode := &BTreeNode{parent: parent, data: rightData, child: nodeNewRightChild}
        // 更新分裂出来后子节点的 parent 指向
        for i := 0; i <= len(newLeftNode.child)-1; i++ {
            newLeftNode.child[i].parent = newLeftNode
        }
        for i := 0; i <= len(newRightNode.child)-1; i++ {
            newRightNode.child[i].parent = newRightNode
        }

        // parent
        if parent == nil {
            parent = &BTreeNode{data: []*item{siftData}, child: []*BTreeNode{newLeftNode, newRightNode}}
            newLeftNode.parent = parent
            newRightNode.parent = parent
        } else {
            var nodeIndex int
            for i := 0; i <= len(parent.child)-1; i++ {
                if parent.child[i] == node {
                    nodeIndex = i
                    break
                }
            }
            parent.data = append(parent.data[:nodeIndex], append([]*item{siftData}, parent.data[nodeIndex:]...)...)
            parent.child = append(parent.child[:nodeIndex], append([]*BTreeNode{newLeftNode, newRightNode}, parent.child[nodeIndex+1:]...)...)
        }
    }
    return parent
}

func (node *BTreeNode) Insert(index int, value interface{}) *BTreeNode {
    if node == nil {
        node = &BTreeNode{}
        itemTmp := &item{index: index, value: value}
        node.data = []*item{itemTmp}
        return node
    }
    if len(node.child) == 0 { // 叶子节点
        for i := 0; i <= len(node.data)-1; i++ {
            if index < node.data[i].index {
                node.data = append(node.data[:i], append([]*item{&item{index: index, value: value}}, node.data[i:]...)...)
                break
            } else if index >= node.data[i].index && i == len(node.data)-1 {
                node.data = append(node.data, &item{index: index, value: value})
                break
            }
        }
        if len(node.data) >= M {
            node = node.Split()
        } else if node.parent != nil {
            return node.parent
        }
    } else if index < node.data[0].index && len(node.child) != 0 {
        node = node.child[0].Insert(index, value)
    } else if index >= node.data[len(node.data)-1].index && len(node.child) != 0 {
        node = node.child[len(node.child)-1].Insert(index, value)
    } else if len(node.child) != 0 {
        for i := 0; i < len(node.data)-1; i++ {
            if index >= node.data[i].index && index < node.data[i+1].index {
                node = node.child[i+1].Insert(index, value)
                break
            }
        }
    }
    for node != nil && len(node.data) >= M {
        node = node.Split()
    }
    if node != nil && node.parent == nil {
        return node
    }
    return node.parent
}

func (node *BTreeNode) Order() []int {
    arr := []int{}
    if node == nil {
        return arr
    }
    // 因为要遍历 child, 所以长度应为 len(node.data)
    for i := 0; i <= len(node.data)-1+1; i++ {
        childData := []int{}
        if len(node.child) != 0 {
            childData = node.child[i].Order()
        }
        if i != len(node.data) {
            arr = append(arr, append(childData, node.data[i].index)...)
        } else {
            arr = append(arr, childData...)
        }
    }
    return arr
}

func (node *BTreeNode) MaxIndex() *item {
    if node == nil {
        return nil
    }
    for len(node.child) != 0 {
        node = node.child[len(node.child)-1]
    }
    return node.data[len(node.data)-1]
}

func (node *BTreeNode) MinIndex() *item {
    if node == nil {
        return nil
    }
    for len(node.child) != 0 {
        node = node.child[0]
    }
    return node.data[0]
}

func (node *BTreeNode) CheckBTree() bool {
    if node == nil {
        return true
    }
    if len(node.data) >= M || len(node.data) < M/2 {
        return false
    }
    lastIndex := node.data[0].index
    for _, data := range node.data {
        if lastIndex > data.index {
            return false
        }
    }
    for i := 0; i <= len(node.child)-1; i++ {
        if !node.child[i].CheckBTree() {
            return false
        }
    }
    return true
}

// 删除后返回 root
func (node *BTreeNode) Delete(index int) *BTreeNode {
    if node.Search(index) == nil {
        // 值不存在情况下直接返回 root
		fmt.Println(index, "not exists")
		for node.parent != nil && len(node.parent.data) != 0 {
			node = node.parent 
		} 
		node.parent = nil
        return node
    }
    if index < node.data[0].index && len(node.child) != 0 {
        return node.child[0].Delete(index)
    } else if index > node.data[len(node.data)-1].index && len(node.child) != 0 {
        return node.child[len(node.child)-1].Delete(index)
    } else {
        for i := 0; i <= len(node.data)-1; i++ {
            if index == node.data[i].index && len(node.child) != 0 { // 非叶子节点
                // 从左右节点中, 子节点长度较长的节点中取 // b-tree 不存在有左无右或有右无左的情况
                if i != 0 && len(node.child[i].data) > len(node.child[i+1].data) {
                    maxValueOfLeftChild := node.child[i].MaxIndex()
                    node.data[i] = maxValueOfLeftChild
                    return node.child[i].Delete(maxValueOfLeftChild.index)
                } else {
                    minValueOfLeftChild := node.child[i+1].MinIndex()
                    node.data[i] = minValueOfLeftChild
                    return node.child[i+1].Delete(minValueOfLeftChild.index)
                }
            } else if len(node.child) != 0 && i < len(node.data)-1 && index >= node.data[i].index && index < node.data[i+1].index {
                return node.child[i+1].Delete(index)
            } else if index == node.data[i].index && len(node.child) == 0 { // 叶子节点
                node.data = append(node.data[:i], node.data[i+1:]...)
                // 从叶子节点往上递归调整
                for node.parent != nil && len(node.parent.data) != 0 {
                    node = node.Adjust()
                }
				// 如果父节点 data 长度为 0, 说明已经合并, parent 指向为 nil 再返回
				if node.parent != nil && len(node.parent.data) == 0 {
					node.parent = nil
				}
                return node
			} else if i == len(node.data)-1 { // TODO: delete this
				fmt.Println("delete else")
			}
        }
    }
	return node
}

func (node *BTreeNode) Adjust() *BTreeNode {
    // root 节点无需调整
    if node.parent == nil {
        return node
    }
	// TODO: delete this
	if node.parent != nil && len(node.parent.data) == 0 {
		node.parent = nil
		return node
	}
    // 非 root 节点的 data 不小于 M/2 不需要调整
    if node.parent != nil && len(node.data) >= M/2 {
        return node.parent
    }
    // 需要调整的情况:
    parent := node.parent
    childIndex := 0
    for childIndex <= len(parent.child)-1 {
        if parent.child[childIndex] == node {
            break
        }
        childIndex++
    }
    if node.parent.parent == nil && childIndex != len(parent.child)-1 { // 因为跟右子节点合并, 所以不为末 child 值
        // ) 需要考虑 root 的子节点合并, 导致 root 节点 data 数量为 0 
		fmt.Println("root child")
        rightBrother := parent.child[childIndex+1]
        node.data = append(node.data, append([]*item{parent.data[childIndex]}, rightBrother.data...)...)
        node.child = append(node.child, rightBrother.child...)
        if len(node.child) != 0 { // 更新 parent 指向
            for i := 0; i <= len(rightBrother.child)-1; i++ {
                rightBrother.child[i].parent = node
            }
        }
        // parent 删掉当前的 data, 右边的子节点
        parent.data = append(parent.data[:childIndex], parent.data[childIndex+1:]...)
        parent.child = append(parent.child[:childIndex+1], parent.child[childIndex+2:]...)
	} else if node.parent.parent == nil && childIndex == len(parent.child)-1 && childIndex != 0 {
		// ) root 的子节点, 跟左兄弟节点合并
		leftBrother := parent.child[childIndex-1]
		leftBrother.data = append(leftBrother.data, append([]*item{parent.data[childIndex-1]}, node.data...)...)
		leftBrother.child = append(leftBrother.child, node.child...)
		if len(node.child) != 0 { // 更新 parent 指向
			for i := 0; i <= len(node.child)-1; i++ {
				node.child[i].parent = leftBrother
			}
		}
		// parent 删掉当前节点, 当前的 child
		parent.data = append([]*item{}, parent.data[:childIndex-1]...)
		parent.child = append([]*BTreeNode{}, parent.child[:len(parent.child)-1]...)
    } else if childIndex != len(parent.child)-1 && len(parent.child[childIndex+1].data) >= (M/2)+1 {
        // 1) 右边 len(data) >= (M/2)+1 时从父节点 i+1 取, 右边取最小值上浮到父节点
        rightBrother := parent.child[childIndex+1]
        node.data = append(node.data, parent.data[childIndex])
        parent.data[childIndex] = rightBrother.data[0]
        rightBrother.data = append([]*item{}, rightBrother.data[1:]...)
        if len(node.child) != 0 { // 更新 parent 指向
            rightBrother.child[0].parent = node
            node.child = append(node.child, rightBrother.child[0])
            rightBrother.child = append([]*BTreeNode{}, rightBrother.child[1:]...)
        }
    } else if childIndex != 0 && childIndex != len(parent.child)-1 &&
        // 2) 右边 len(data) == M/2, 左边 len(data) >= (M/2)+1 时从父节点 i 取, 左边取最大值上浮到父节点
        len(parent.child[childIndex+1].data) <= M/2 && len(parent.child[childIndex-1].data) >= (M/2)+1 {
        leftBrother := parent.child[childIndex-1]
        node.data = append([]*item{parent.data[childIndex-1]}, node.data...)
        parent.data[childIndex-1] = leftBrother.data[len(leftBrother.data)-1]
        leftBrother.data = append([]*item{}, leftBrother.data[:len(leftBrother.data)-1]...)
        if len(node.child) != 0 { // 更新 parent 指向
            leftBrother.child[len(leftBrother.child)-1].parent = node
            node.child = append([]*BTreeNode{leftBrother.child[len(leftBrother.child)-1]}, node.child...)
            leftBrother.child = append([]*BTreeNode{}, leftBrother.child[:len(leftBrother.child)-1]...)
        }
    } else if childIndex == len(parent.child)-1 && childIndex != 0 && len(parent.child[childIndex-1].data) == M/2 {
        // 3) 如果是末尾 child, 左不足时, 跟左兄弟节点合并
        leftBrother := parent.child[childIndex-1]
        leftBrother.data = append(leftBrother.data, append([]*item{parent.data[childIndex-1]}, node.data...)...)
        leftBrother.child = append(leftBrother.child, node.child...)
        if len(node.child) != 0 { // 更新 parent 指向
            for i := 0; i <= len(node.child)-1; i++ {
                node.child[i].parent = leftBrother
            }
        }
        // parent 删掉当前节点, 当前的 child
        parent.data = append([]*item{}, parent.data[:childIndex-1]...)
        parent.child = append([]*BTreeNode{}, parent.child[:len(parent.child)-1]...)
    } else if childIndex != len(parent.child)-1 &&
        ((childIndex == 0) || (childIndex != 0 && len(parent.child[childIndex-1].data) <= M/2)) &&
        len(parent.child[childIndex+1].data) == M/2 {
        // 4) 左右 == M/2(不足以取节点), 且不为末节点(即有右节点), 跟右兄弟节点合并(因分裂时分裂出来的是右兄弟节点)
        rightBrother := parent.child[childIndex+1]
        node.data = append(node.data, append([]*item{parent.data[childIndex]}, rightBrother.data...)...)
        node.child = append(node.child, rightBrother.child...)
        if len(node.child) != 0 { // 更新 parent 指向
            for i := 0; i <= len(rightBrother.child)-1; i++ {
                rightBrother.child[i].parent = node
            }
        }
        // parent 删掉当前的 data, 右边的子节点
        parent.data = append(parent.data[:childIndex], parent.data[childIndex+1:]...)
        parent.child = append(parent.child[:childIndex+1], parent.child[childIndex+2:]...)
	// } else if len(node.data) <= M/2 { // TODO: delete
	// 	fmt.Println("value of data", node.data[0])
	}
	if len(node.parent.data) == 0 && node.parent.parent == nil {
		node.parent = nil
		return node
	}

    return node.parent
}

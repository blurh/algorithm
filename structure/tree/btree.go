package tree

// 定义树的阶
const M int = 5

// 节点可以有的最大键值数
const Min int = M - 1

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
    for tree.root.parent != nil {
        tree.root = tree.root.parent
    }
    tree.count += 1
    return true
}

func (tree *BTree) Delete() bool {
    tree.root = tree.root.Delete()
    return true
}

func (tree *BTree) Order() []int {
    return tree.root.Order()
}

func (tree *BTree) MinIndexOfTree() *item {
    root := tree.root
    for len(root.child) != 0 {
        root = root.child[0]
    }
    return root.data[0]
}

func (tree *BTree) MaxIndexOfTree() *item {
    root := tree.root
    for len(root.child) != 0 {
        root = root.child[len(root.child)-1]
    }
    return root.data[len(root.data)-1]
}

func (tree *BTree) CheckBTree() bool {
    return tree.root.CheckBTree()
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

func (node *BTreeNode) Search(index int) interface{} {
    if node == nil {
        return nil
    }
    for i := 0; i <= len(node.data)-1; i++ {
        if index == node.data[i].index {
            return node.data[i].value
        } else if index < node.data[0].index && len(node.child) != 0 {
            return node.child[0].Search(index)
        } else if index >= node.data[i].index && (i != len(node.data)-1 && index < node.data[i+1].index) && len(node.child) != 0 {
            return node.child[i+1].Search(index)
        } else if i == len(node.data)-1 && len(node.child) != 0 {
            return node.child[len(node.child)-1].Search(index)
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
        // 分裂 child
        nodeNewLeftChild := node.child[:len(node.child)/2]
        nodeNewRightChild := node.child[len(node.child)/2:]
        // 分裂的 node 节点
        newLeftNode := &BTreeNode{parent: parent, data: node.data[:indexOfSift], child: nodeNewLeftChild}
        newRightNode := &BTreeNode{parent: parent, data: node.data[indexOfSift+1:], child: nodeNewRightChild}
        // 修复分裂的子节点的 parent 指向
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
    // nil 插入不需要分裂
    if node == nil {
        node = &BTreeNode{}
        itemTmp := &item{index: index, value: value}
        node.data = []*item{itemTmp}
        return node
    }
    if len(node.child) == 0 { // 叶子节点
        for i := 0; i <= len(node.data)-1; i++ {
            if index <= node.data[i].index {
                node.data = append(node.data[:i], append([]*item{&item{index: index, value: value}}, node.data[i:]...)...)
                break
            } else if index > node.data[i].index && i == len(node.data)-1 {
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
            if index > node.data[i].index && index < node.data[i+1].index {
                node = node.child[i+1].Insert(index, value)
                break
            }
        }
    }
    if node != nil && len(node.data) >= M {
        node = node.Split()
    }
    return node
}

func (node *BTreeNode) Order() []int {
    arr := []int{}
    if node == nil {
        return arr
    }
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

func (node *BTreeNode) CheckBTree() bool {
    if node == nil {
        return true
    }
    if len(node.data) >= M {
        return false
    }
    var lastIndex int
    for i, data := range node.data {
        if i == 0 {
            lastIndex = data.index
        }
        if lastIndex > data.index {
            return false
        }
    }
    for i := 0; i <= len(node.child)-1; i++ {
        return node.child[i].CheckBTree()
    }
    return true
}

func (node *BTreeNode) Delete() *BTreeNode {
    // TODO
    return node
}

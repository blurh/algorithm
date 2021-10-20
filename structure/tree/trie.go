package tree

// TrieNode 支持大小写, 所以 52 叉, count 表示引用计数, 用于查找和删除
type TrieNode struct {
    endflag    bool
    letter     string
    count      int
    dictionary [52]*TrieNode
}

type TrieTree struct {
    root *TrieNode
}

func InitTrieTree() *TrieTree {
    node := new(TrieNode)
    node.count = 0
    tree := new(TrieTree)
    tree.root = node
    return tree
}

func (tree *TrieTree) AddWord(word string) bool {
    if tree.FindWord(word) {
        return false
    }
    node := tree.root
    baseNum := int('A') // 65
    for i := 0; i < len(word); i++ {
        var num int
        if int(word[i]) >= 97 {
            num = int(word[i]) - baseNum - 6
        } else {
            num = int(word[i]) - baseNum
        }
        if node.dictionary[num] == nil {
            node.dictionary[num] = &TrieNode{count: 1, letter: string(word[i])}
            node = node.dictionary[num]
        } else {
            node.count++
            node = node.dictionary[num]
        }
        if i == len(word)-1 {
            node.endflag = true
        }
    }
    return true
}

func (tree *TrieTree) FindWord(word string) bool {
    node := tree.root
    baseNum := int('A') // 65
    for i := 0; i < len(word); i++ {
        var num int
        if int(word[i]) >= 97 {
            num = int(word[i]) - baseNum - 6
        } else {
            num = int(word[i]) - baseNum
        }
        // 判断词末 endflag
        if i == len(word)-1 && node.dictionary[num] != nil && node.dictionary[num].endflag == false {
            return false
        }
        if node.dictionary[num] != nil {
            node = node.dictionary[num]
        } else {
            return false
        }
    }
    return true
}

func (tree *TrieTree) DelWord(word string) bool {
    if !tree.FindWord(word) {
        return false
    }
    node := tree.root
    baseNum := int('A') // 65
    for i := 0; i < len(word); i++ {
        var num int
        if int(word[i]) >= 97 {
            num = int(word[i]) - baseNum - 6
        } else {
            num = int(word[i]) - baseNum
        }
        node.count--
        if i == len(word)-1 {
            node.endflag = false
        }
        if node.dictionary[num] != nil {
            if node.count <= 0 {
                node.dictionary[num] = nil
                node = nil
                return true
            }
            node = node.dictionary[num]
        }
    }
    return true
}

func (tree *TrieTree) GetWordCount() int {
    count := len(tree.GetAllWord())
    return count
}

func (tree *TrieTree) GetAllWord() []string {
    return order(tree.root, "")
}

// 顺序遍历
func order(tree *TrieNode, letter string) []string {
    if tree == nil {
        return []string{}
    }
    if len(tree.letter) != 0 {
        letter += tree.letter
    }
    var arr []string
    // 搜索到词末标志时, 判断是否有后续节点
    if tree.endflag == true {
        flag := false
        for i := 0; i < 52; i++ {
            if tree.dictionary[i] != nil {
                // 判断是否第一次, 避免前缀重复
                // 如 [ abc ab abe ab ] 的 ab 重复
                if !flag {
                    arr = append(arr, append([]string{letter}, order(tree.dictionary[i], letter)...)...)
                } else {
                    arr = append(arr, order(tree.dictionary[i], letter)...)
                }
                flag = true
            }
        }
        if flag {
            return arr
        }
        return []string{letter}
    }
    // 非词末情况
    for i := 0; i < 52; i++ {
        orderArr := order(tree.dictionary[i], letter)
        arr = append(arr, orderArr...)
    }
    return arr
}

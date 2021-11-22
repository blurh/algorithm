package tree

import (
    "fmt"
    "sort"
    "strings"
)

type TrieNode struct {
    count    int
    endflag  bool
    word     string
    next     []*TrieNode
}

func indexOfNext(words []*TrieNode, word string) int {
    for i, v := range words {
        if v.word == word {
            return i
        }
    }
    return -1
}

type TrieTree struct {
    root    *TrieNode
    count   int
    sep     string
}

func InitTrieTree(sep string) *TrieTree {
    node := new(TrieNode)
    node.count = 0
    tree := new(TrieTree)
    tree.root = node
    tree.count = 0
    tree.sep = sep
    return tree
}

func (tree *TrieTree) Count() int {
    return tree.count
}

func (tree *TrieTree) Add(word string) bool {
    if tree.Get(word) {
        return false
    }
    sep := tree.sep
    node := tree.root
    splitWordArray := strings.Split(word, sep)
    for i := 0; i <= len(splitWordArray)-1; i++ {
        index := indexOfNext(node.next, splitWordArray[i])
        if index == -1 {
            node.next = append(node.next, &TrieNode{count: 1, word: splitWordArray[i], endflag: false})
            node = node.next[len(node.next)-1]
        } else {
            node = node.next[index]
            node.count++
        }
        if i == len(splitWordArray)-1 {
            node.endflag = true
        }
    }
    tree.count++
    return true
}

func (tree *TrieTree) Get(word string) bool {
    sep := tree.sep
    splitWordArray := strings.Split(word, sep)
    node := tree.root
    for i := 0; i <= len(splitWordArray)-1; i++ {
        index := indexOfNext(node.next, splitWordArray[i])
        if index == -1 {
            return false
        }
        node = node.next[index]
    }
    // 判断词末 endflag
    if node.endflag == false {
        return false
    }
    return true
}

func (tree *TrieTree) Delete(word string) bool {
    if !tree.Get(word) {
        return false
    }
    sep := tree.sep
    node := tree.root
    splitWordArray := strings.Split(word, sep)
    for i := 0; i <= len(splitWordArray)-1; i++ {
        index := indexOfNext(node.next, splitWordArray[i])
        node.next[index].count--
        if node.next[index].count <= 1 && i != 0 {
            node.next = append(node.next[:index], node.next[index+1:]...)
            tree.count--
            return true
        } 
        if i == len(splitWordArray)-1 {
            node.next[index].endflag = false
        }
        node = node.next[index]
    }
    tree.count--
    return true
}

func (tree *TrieTree) Words() []string {
    orderArr := tree.root.Order(tree.sep, tree.sep)
    sort.Strings(orderArr)
    return orderArr
}

func (node *TrieNode) Order(word, sep string) []string {
    if node == nil {
        return []string{}
    }
    if word == sep {
        word = fmt.Sprintf("%s%s", sep, node.word)
    } else {
        word = fmt.Sprintf("%s%s%s", word, sep, node.word)
    }
    var arr []string
    // 搜索到词末标志时, 判断是否有后续节点
    if node.endflag == true {
        flag := false
        for i := 0; i < len(node.next); i++ {
            // 判断是否第一次, 避免前缀重复
            // 如 [ abc ab abe ab ] 的 ab 重复
            if !flag {
                arr = append(arr, append([]string{word}, node.next[i].Order(word, sep)...)...)
            } else {
                arr = append(arr, node.next[i].Order(word, sep)...)
            }
            flag = true
        }
        if flag {
            return arr
        }
        return []string{word}
    }
    // 非词末情况
    for i := 0; i < len(node.next); i++ {
        orderArr := node.next[i].Order(word, sep)
        arr = append(arr, orderArr...)
    }
    return arr
}

func (tree *TrieTree) PartOrder() []string {
    return tree.root.PartOrder(tree.sep, tree.sep)
}

func (node *TrieNode) PartOrder(word, sep string) []string {
    if word == sep {
        word = fmt.Sprintf("%s%s", sep, node.word)
    } else {
        word = fmt.Sprintf("%s%s%s", word, sep, node.word)
    }
    arr := []string{word}
    for _, v := range node.next {
        arr = append(arr, v.Order(word, sep)...)
    }
    return arr
}
package tree

// TODO: 随机生成一串数组, 插入, 再根据特性 check

import (
    "fmt"
    "reflect"
    "testing"
)

func TestTree(t *testing.T) {
    t.Run("test of binary tree", func(t *testing.T) {
        tree := InitBinaryTree(123)
        if rootData := tree.GetNodeData(); rootData != 123 {
            t.Errorf("root data not equal 123, fail")
        }

        tree.AddLeft(1)

        if err := tree.AddLeft(2); err == nil {
            t.Error("fail")
        }
        if letfNodeData := tree.leftNode.GetNodeData(); letfNodeData != 1 {
            t.Errorf("left data is not 1, fail")
        }

        tree.AddRight(2)
        tree.leftNode.AddLeft(3)
        tree.leftNode.AddRight(4)
        tree.leftNode.leftNode.AddLeft(5)

        if treeHeight := tree.GetTreeHeight(); treeHeight != 4 {
            t.Errorf("height of tree is not 4, fail")
        }

        if leafNum := tree.GetLeafNum(); leafNum != 3 {
            t.Errorf("leaf num is not 3, fail")
        }

        preOrderArr := tree.PreOrder()
        if !reflect.DeepEqual(preOrderArr, []int{123, 1, 3, 5, 4, 2}) {
            t.Errorf("pre order fail")
        }

        middleOrderArr := tree.MiddleOrder()
        if !reflect.DeepEqual(middleOrderArr, []int{5, 3, 1, 4, 123, 2}) {
            t.Errorf("middle order fail")
        }

        postOrderArr := tree.PostOrder()
        if !reflect.DeepEqual(postOrderArr, []int{5, 3, 4, 1, 2, 123}) {
            t.Errorf("post order fail")
        }

        breadthFirstSearchArr := tree.BreadthFirstSearch()
        if !reflect.DeepEqual(breadthFirstSearchArr, []int{123, 1, 2, 3, 4, 5}) {
            t.Errorf("breadth first search fail")
        }

        invertRoot := Invert(tree)
        invertPreOrderArr := invertRoot.PreOrder()
        if !reflect.DeepEqual(invertPreOrderArr, []int{123, 2, 1, 4, 3, 5}) {
            t.Errorf("invert fail")
        }

        if !tree.SearchValue(5) {
            t.Errorf("search 5 fail, fail")
        }
        if tree.SearchValue(6) {
            t.Errorf("search 6 success, fail")
        }
    })
    t.Run("test of binary search tree", func(t *testing.T) {
        tree := InitSearchTree(8)
        treeArr := []int{3, 10, 1, 6, 4, 14, 7, 13}
        for _, v := range treeArr {
            tree.InsertNode(v)
        }
        if tree.InsertNode(13) {
            t.Errorf("insert 13 twice not false, fail")
        }
        if tree.SearchValue(18) {
            t.Errorf("search 18 not false, fail")
        }
        if maxValue := tree.MaxOfSearchTree(); maxValue != 14 {
            t.Errorf("max value is not 14, fail")
        }
        if minValue := tree.MinOfSearchTree(); minValue != 1 {
            t.Errorf("min value is not 1, fail")
        }

        preArr := tree.PreOrder()
        if !reflect.DeepEqual(preArr, []int{8, 3, 1, 6, 4, 7, 10, 14, 13}) {
            t.Errorf("pre order fail")
        }
        middleArr := tree.MiddleOrder()
        if !reflect.DeepEqual(middleArr, []int{1, 3, 4, 6, 7, 8, 10, 13, 14}) {
            t.Errorf("middle order fail")
        }
        postArr := tree.PostOrder()
        if !reflect.DeepEqual(postArr, []int{1, 4, 7, 6, 3, 13, 14, 10, 8}) {
            t.Errorf("middle order fail")
        }

        if !tree.SearchValue(13) {
            t.Errorf("search 13 false, fail")
        }

        tree.RemoveNode(13)
        middleArr = tree.MiddleOrder()
        if !reflect.DeepEqual(middleArr, []int{1, 3, 4, 6, 7, 8, 10, 14}) {
            t.Errorf("remove node fail")
        }
    })
    t.Run("test of AVL tree", func(t *testing.T) {
        checkAVL := func(tree *avlTree) {
            if checkResult := tree.CheckAVLTree(); !checkResult {
                t.Errorf("check AVL tree return false, fail")
            }
        }
        tree := InitAvlTree(4)
        for _, v := range []int{3, 4, 5, 6, 7, 9, 11, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23} {
            tree = tree.InsertValue(v)
            checkAVL(tree)
        }
        for _, v := range []int{23, 13, 4, 19, 16, 22, 21, 20, 19, 18, 17, 15, 9, 5, 3, 6} {
            tree = tree.RemoveValue(v)
            checkAVL(tree)
        }
        for i := 10000; i >= 0; i-- {
            tree = tree.InsertValue(i)
        }
        checkAVL(tree)
        for i := 0; i <= 10000; i++ {
            tree = tree.RemoveValue(i)
        }
        checkAVL(tree)
    })
    t.Run("test of red black tree", func(t *testing.T) {
        checkRBTree := func(tree *rbTree) {
            if tree.CheckRBTree() != 0 {
                t.Errorf("check red black tree return %d, fail", tree.CheckRBTree())
            }
        }
        tree := InitRBTree()
        tree.Insert(1, 111)
        tree.Remove(1)
        if tree.root != nil {
            t.Errorf("remove root key error")
        }
        // test for insert
        for _, v := range []int{10, 13, 11, 7, 6, 5, 3, 2, 1, 16, 17, 18, 15, 50, 25, 30, 32, 4} {
            tree.Insert(v, fmt.Sprintf("%d%d%d", v, v, v))
            checkRBTree(tree)
        }
        // test for update
        tree.Update(10, 101010)
        if tree.Get(10) != 101010 {
            t.Errorf("red black tree update value fail")
        }
        tree.Set(12, 121212)
        if tree.Get(12) != 121212 {
            t.Errorf("red black tree set value fail")
        }

        for _, v := range []int{13, 5, 6, 7, 4, 3, 15, 1, 16, 17, 11, 18, 2, 15, 4, 32, 30, 50, 10, 25} {
            tree.Remove(v)
            // checkRBTree(tree)
        }
        for i := 1000; i > 0; i-- {
            tree.Insert(i, i)
            checkRBTree(tree)
        }
        for i := 100; i <= 1000; i++ {
            tree.Remove(i)
            // checkRBTree(tree)
        }
        checkRBTree(tree)

        tree.Clear()
        if len(tree.Order()) != 0 {
            t.Errorf("clear red black tree fail")
        }
    })
    t.Run("test of treap", func(t *testing.T) {
        checkTreap := func(tree *treap) {
            if !tree.CheckTreap() {
                t.Errorf("check treap fail")
            }
        }
        tree := InitTreap(13, 1)
        for i, v := range []int{13, 5, 6, 7, 4, 3, 15, 1, 16, 17, 11, 18, 2, 15, 4, 32, 30, 50} {
            tree = tree.InsertValue(v, i)
            checkTreap(tree)
        }
        for _, v := range []int{5, 6, 4, 7, 13, 15, 1, 16, 17, 30, 11, 4, 3, 18, 15, 32, 50, 2} {
            tree = tree.RemoveValue(v)
            checkTreap(tree)
        }
        if tree != nil {
            t.Errorf("remove fail")
        }
    })
    t.Run("test of trie", func(t *testing.T) {
        tree := InitTrieTree()
        for _, word := range []string{"abc", "abd"} {
            tree.AddWord(word)
        }
        // find exists word
        if !tree.FindWord("abc") || !tree.FindWord("abd") {
            t.Errorf("find exists word return false, fail")
        }
        // find not exists word
        if tree.FindWord("abde") || tree.FindWord("abe") || tree.FindWord("ab") {
            t.Errorf("find not exists word return true, fail")
        }
        tree.AddWord("abe")
        // delete word
        tree.DelWord("abc")
        if tree.FindWord("abc") {
            t.Errorf("find deleted word return true, fail")
        }
        if !tree.FindWord("abd") {
            t.Errorf("find exists word return false, fail")
        }
        tree.AddWord("Abcdef")
        if tree.FindWord("Ab") {
            t.Errorf("find path of word return true, fail")
        }
        tree.AddWord("Ab")
        tree.AddWord("Abc")
        tree.AddWord("Abcd")
        tree.AddWord("Abe")
        if !reflect.DeepEqual(tree.GetAllWord(), []string{"Ab", "Abc", "Abcd", "Abcdef", "Abe", "abd", "abe"}) {
            t.Errorf("get all word fail")
        }
        if tree.GetWordCount() != 7 {
            t.Errorf("get word fail")
        }
    })
    t.Run("test of b-tree", func(t *testing.T) {
        var arr []int
        tree := InitBTree()
        for i := 30; i <= 1500; i++ {
            tree.Insert(i, fmt.Sprintf("%d%d%d", i, i, i))
            arr = append(arr, i)
        }
        for i := 1500; i > 1030; i-- {
            tree.Insert(i, fmt.Sprintf("%d%d%d", i, i, i))
            arr = append(arr, i)
        }
        for i := 1500; i > 1030; i-- {
            tree.Insert(1500, fmt.Sprintf("%d%d%d", 1500, 1500, 1500))
            arr = append(arr, 1500)
        }
        for i := 10000; i > 0; i-- {
            tree.Insert(i, fmt.Sprintf("%d%d%d", i, i, i))
            arr = append(arr, i)
        }
        for _, v := range []int{4, 3, 1, 10, 1000, 2, 7, 20, 8, 9, 19, 100, 1000} {
            if tree.Search(v) != fmt.Sprintf("%d%d%d", v, v, v) {
                t.Errorf("get value %d fail", v)
            }
        }
        // max and min of array
        max := func(arr []int) int {
            maxValue := 0
            for _, v := range arr {
                if v > maxValue {
                    maxValue = v
                }
            }
            return maxValue
        }
        min := func(arr []int) int {
            minValue := arr[0]
            for _, v := range arr {
                if v < minValue {
                    minValue = v
                }
            }
            return minValue
        }
        if tree.MaxIndexOfTree().index != max(arr) {
            t.Errorf("get max of b-tree fail")
        }
        if tree.MinIndexOfTree().index != min(arr) {
            t.Errorf("get min of b-tree fail")
        }
        // check count
        if tree.count != len(arr) {
            t.Errorf("get count of b-tree not match len of test arr, fail")
        }

        // check order
        orderArr := tree.Order()
        if len(orderArr) != len(arr) {
            t.Errorf("order b-tree fail")
        }
        lastIndex := orderArr[0]
        for _, v := range orderArr {
            if lastIndex > v {
                t.Errorf("order b-tree fail")
            }
        }

        // check b-tree after insert value
        if tree.CheckBTree() != 0 {
            t.Errorf("check b-tree return %d, fail", tree.CheckBTree())
        }

        for _, v := range arr {
            tree.Delete(v)
        }

        // check b-tree after delete value
        checkReturn := tree.CheckBTree()
        if  checkReturn != 0 {
            t.Errorf("check b-tree return %d, fail", checkReturn)
        }
    })
}

package tree

// TODO: 随机生成一串数组, 插入, 再根据特性 check

import (
	"reflect"
	"testing"
)

func TestStructure(t *testing.T) {
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

		if searchResultFound := tree.SearchValue(5); !searchResultFound {
			t.Errorf("search 5 fail, fail")
		}
		if searchResultNotFound := tree.SearchValue(6); searchResultNotFound {
			t.Errorf("search 6 success, fail")
		}
	})
	t.Run("test of binary search tree", func(t *testing.T) {
		tree := InitSearchTree(8)
		treeArr := []int{3, 10, 1, 6, 4, 14, 7, 13}
		for _, v := range treeArr {
			tree.InsertNode(v)
		}
		if insertRes := tree.InsertNode(13); insertRes {
			t.Errorf("insert 13 twice not false, fail")
		}
		if res := tree.SearchValue(18); res {
			t.Errorf("search 18 not false, fail")
		}
		if res := tree.SearchValue(14); !res {
			t.Errorf("search 14 false, fail")
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

		if searchValueRet13 := tree.SearchValue(13); !searchValueRet13 {
			t.Errorf("search 13 false, fail")
		}
		if searchValueRet18 := tree.SearchValue(18); searchValueRet18 {
			t.Errorf("search 18 not false, fail")
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
	})
	t.Run("test of red black tree", func(t *testing.T) {
		rotateTestTree := InitRBTree(10)
		rotateTestTree.leftNode = newNode(7, rotateTestTree)
		// 两节点右转
		rotateTestTree = rotateTestTree.RightRotate()
		if rotateTestTree.data != 7 || rotateTestTree.rightNode.data != 10 {
			t.Errorf("right rotate fail")
		}
		// 两节点左转
		rotateTestTree = rotateTestTree.LeftRotate()
		if rotateTestTree.data != 10 || rotateTestTree.leftNode.data != 7 {
			t.Errorf("left rotate fail")
		}
		rotateTestTree.leftNode.leftNode = newNode(5, rotateTestTree.leftNode)
		// 三节点右转
		rotateTestTree = rotateTestTree.RightRotate()
		if rotateTestTree.data != 7 || rotateTestTree.leftNode.data != 5 || rotateTestTree.rightNode.data != 10 {
			t.Errorf("right rotate fail")
		}
		// 构造右斜树测试左转
		rotateTestTree = rotateTestTree.RightRotate()
		// 三节点左转
		rotateTestTree = rotateTestTree.LeftRotate()
		if rotateTestTree.data != 7 || rotateTestTree.leftNode.data != 5 || rotateTestTree.rightNode.data != 10 {
			t.Errorf("left rotate fail")
		}
		// 构造标准的旋转的树
		rotateTestTree.leftNode.leftNode = newNode(3, rotateTestTree.leftNode)
		rotateTestTree.leftNode.rightNode = &rbTree{data: 6, color: RED, parent: rotateTestTree.leftNode}
		// 标准树的右转
		rotateTestTree = rotateTestTree.RightRotate()
		if rotateTestTree.data != 5 || rotateTestTree.leftNode.data != 3 || rotateTestTree.rightNode.data != 7 ||
			rotateTestTree.rightNode.leftNode.data != 6 || rotateTestTree.rightNode.rightNode.data != 10 {
			t.Errorf("right rotate fail")
		}
		// 标准树的左转
		rotateTestTree = rotateTestTree.LeftRotate()
		if rotateTestTree.data != 7 || rotateTestTree.rightNode.data != 10 || rotateTestTree.leftNode.data != 5 ||
			rotateTestTree.leftNode.leftNode.data != 3 || rotateTestTree.leftNode.rightNode.data != 6 {
			t.Errorf("left rotate fail")
		}
		// 左转然后右转
		lrRotateTree := rotateTestTree.LeftRightRotate()
		if lrRotateTree.data != 6 && lrRotateTree.leftNode.data != 5 &&
			lrRotateTree.leftNode.leftNode.data != 3 && lrRotateTree.rightNode.data != 7 &&
			lrRotateTree.rightNode.rightNode.data != 10 {
			t.Errorf("left right rotate fail")
		}
		// 右转然后左转
		rlRotateTree := InitRBTree(7)
		rlRotateTree.leftNode = &rbTree{data: 5, parent: rlRotateTree, color: BLACK}
		rlRotateTree.rightNode = &rbTree{data: 10, parent: rlRotateTree, color: BLACK}
		rlRotateTree.rightNode.leftNode = &rbTree{data: 9, parent: rlRotateTree.rightNode, color: RED}
		rlRotateTree.rightNode.rightNode = &rbTree{data: 11, parent: rlRotateTree.rightNode, color: RED}
		rlRotateTree = rlRotateTree.RightLeftRotate()
		if rlRotateTree.data != 9 && rlRotateTree.leftNode.data != 7 && rlRotateTree.rightNode.data != 10 &&
			rlRotateTree.rightNode.rightNode.data != 10 && rlRotateTree.rightNode.rightNode.rightNode.data != 11 {
			t.Errorf("right left rotate fail")
		}
		// 左转然后右转
		rlRotateTree = InitRBTree(6)
		rlRotateTree.leftNode = newNode(7, rlRotateTree)
		rlRotateTree.leftNode.leftNode = &rbTree{data: 5, parent: rlRotateTree.leftNode, color: BLACK}
		rlRotateTree.leftNode.rightNode = &rbTree{data: 10, parent: rlRotateTree.leftNode, color: BLACK}
		rlRotateTree.leftNode.rightNode.leftNode = &rbTree{data: 9, parent: rlRotateTree.leftNode.rightNode, color: RED}
		rlRotateTree.leftNode.rightNode.rightNode = &rbTree{data: 11, parent: rlRotateTree.leftNode.rightNode, color: RED}
		rlRotateTree.leftNode = rlRotateTree.leftNode.RightLeftRotate()
		// rlRotateTree.leftNode.rightNode = rlRotateTree.leftNode.rightNode.RightRotate()
		// rlRotateTree.leftNode = rlRotateTree.leftNode.LeftRotate()
		// end of rotate test

		checkRBTree := func(tree *rbTree) {
			if tree.parent != nil {
				t.Errorf("root's parent is not nil, fail")
			}
			if !tree.CheckRBTree() {
				t.Errorf("check red black fail")
			}
		}
		tree := InitRBTree(10)
		for _, v := range []int{13, 11, 7, 6, 5, 4, 3, 2, 1, 16, 17, 18, 15} {
			tree = tree.InsertValue(v)
			checkRBTree(tree)
		}
		tree = tree.Clear()
		if len(tree.Order()) != 0 {
			t.Errorf("clear red black tree fail")
		}
	})
}

//usr/bin/env go test .; exit
package structure

import (
	"reflect"
	"testing"
)

func TestStructure(t *testing.T) {
	t.Run("test of list", func(t *testing.T) {
		l := InitList()
		err := l.InsertNode(1, 100)
		if err != nil {
			t.Errorf("insert err, fail")
		}
		valueOfIndex1, err := l.GetIndexData(1)
		if valueOfIndex1 != 100 || err != nil {
			t.Errorf("get index of 1 err, fail")
		}
		err = l.AppendNode(200)
		if err != nil {
			t.Errorf("append value err, fail")
		}
		valueOfIndex2, err := l.GetIndexData(2)
		if valueOfIndex2 != 200 || err != nil {
			t.Errorf("get index of 2 err, fail")
		}
		err = l.SetIndexNodeData(1, 1000)
		setIndexNodeDataResult, _ := l.GetIndexData(1)
		if setIndexNodeDataResult != 1000 {
			t.Errorf("set index err, fail")
		}
		lenOfList := l.GetListLength()
		if lenOfList != 3 {
			t.Errorf("len of list is not 2, fail")
		}
		l.DeleteIndexNode(1)
		delIndexNodeGetData, _ := l.GetIndexData(1)
		if delIndexNodeGetData != 200 {
			t.Errorf("delete error, fail")
		}
	})
	t.Run("test of binary tree", func(t *testing.T) {
		root := InitBinaryTree()
		root.data = 123
		rootData := root.GetNodeData()
		if rootData != 123 {
			t.Errorf("root data not equal 123, fail")
		}

		root.AddLeft(1)
		if err := root.AddLeft(2); err == nil {
			t.Error("fail")
		}
		letfNodeData := root.leftNode.GetNodeData()
		if letfNodeData != 1 {
			t.Errorf("left data is not 1, fail")
		}

		root.AddRight(2)
		root.leftNode.AddLeft(3)
		root.leftNode.AddRight(4)
		root.leftNode.leftNode.AddLeft(5)

		treeHeight := root.GetTreeHeight()
		if treeHeight != 4 {
			t.Errorf("height of tree is not 4, fail")
		}

		leafNum := root.GetLeafNum()
		if leafNum != 3 {
			t.Errorf("leaf num is not 3, fail")
		}

		preOrderArr := root.PreOrder()
		if !reflect.DeepEqual(preOrderArr, []int{123, 1, 3, 5, 4, 2}) {
			t.Errorf("pre order fail")
		}

		middleOrderArr := root.MiddleOrder()
		if !reflect.DeepEqual(middleOrderArr, []int{5, 3, 1, 4, 123, 2}) {
			t.Errorf("middle order fail")
		}

		postOrderArr := root.PostOrder()
		if !reflect.DeepEqual(postOrderArr, []int{5, 3, 4, 1, 2, 123}) {
			t.Errorf("post order fail")
		}

		breadthFirstSearchArr := root.BreadthFirstSearch()
		if !reflect.DeepEqual(breadthFirstSearchArr, []int{123, 1, 2, 3, 4, 5}) {
			t.Errorf("breadth first search fail")
		}

		invertRoot := Invert(root)
		invertPreOrderArr := invertRoot.PreOrder()
		if !reflect.DeepEqual(invertPreOrderArr, []int{123, 2, 1, 4, 3, 5}) {
			t.Errorf("invert fail")
		}

		searchResultFound := root.SearchValue(5)
		if !searchResultFound {
			t.Errorf("search 5 fail, fail")
		}
		searchResultNotFound := root.SearchValue(6)
		if searchResultNotFound {
			t.Errorf("search 6 success, fail")
		}
	})
	t.Run("test of binary search tree", func(t *testing.T) {
		tree := InitSearchTree()
		tree.data = 8
		treeArr := []int{3, 10, 1, 6, 4, 14, 7, 13}
		for _, v := range treeArr {
			tree.InsertNode(v)
		}
		insertRes := tree.InsertNode(13)
		if insertRes {
			t.Errorf("insert 13 twice not false, fail")
		}
		res := tree.SearchValue(18)
		if res {
			t.Errorf("search 18 not false, fail")
		}
		res = tree.SearchValue(14)
		if !res {
			t.Errorf("search 14 false, fail")
		}
		maxValue := tree.MaxOfSearchTree()
		if maxValue != 14 {
			t.Errorf("max value is not 14, fail")
		}
		minValue := tree.MinOfSearchTree()
		if minValue != 1 {
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

		searchValueRet13 := tree.SearchValue(13)
		searchValueRet18 := tree.SearchValue(18)
		if !searchValueRet13 {
			t.Errorf("search 13 false, fail")
		}
		if searchValueRet18 {
			t.Errorf("search 18 not false, fail")
		}

		tree.RemoveNode(13)
		middleArr = tree.MiddleOrder()
		if !reflect.DeepEqual(middleArr, []int{1, 3, 4, 6, 7, 8, 10, 14}) {
			t.Errorf("remove node fail")
		}
	})
}

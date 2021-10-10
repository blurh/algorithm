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
		if err = l.AppendNode(200); err != nil {
			t.Errorf("append value err, fail")
		}
		valueOfIndex2, err := l.GetIndexData(2)
		if valueOfIndex2 != 200 || err != nil {
			t.Errorf("get index of 2 err, fail")
		}
		l.SetIndexNodeData(1, 1000)
		setIndexNodeDataResult, _ := l.GetIndexData(1)
		if setIndexNodeDataResult != 1000 {
			t.Errorf("set index err, fail")
		}
		if lenOfList := l.GetListLength(); lenOfList != 3 {
			t.Errorf("len of list is not 2, fail")
		}
		l.DeleteIndexNode(1)
		if delIndexNodeGetData, _ := l.GetIndexData(1); delIndexNodeGetData != 200 {
			t.Errorf("delete error, fail")
		}
	})
	t.Run("test of stack", func(t *testing.T) {
		s := InitStack()
		s.Push(1)
		s.Push(2)
		s.Push("testOfStack")
		if s.Len() != 3 {
			t.Errorf("cap of s is not 3, fail")
		}
		if value, _ := s.Pop(); value != "testOfStack" {
			t.Errorf("top of stack is not expections, fail")
		}
		s.Clear()
		if s.Len() != 0 {
			t.Errorf("after clear is not empty, fail")
		}
	})
	t.Run("test of queue", func(t *testing.T) {
		q := InitQueue()
		testArr := []int{2, 4, 7, 1, 3, 10, 24, 3}
		for _, v := range testArr {
			q.Push(v)
		}
		if q.Front() != testArr[len(testArr)-1] {
			t.Errorf("check front func fail")
		}
		if q.Back() != testArr[0] {
			t.Errorf("check back func fail")
		}
		checkArr := []int{}
		for range testArr {
			checkArr = append(checkArr, q.Pop())
		}
		if !reflect.DeepEqual(checkArr, testArr) {
			t.Errorf("check queue fail")
		}
	})
}

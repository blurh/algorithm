package structure

import (
    "reflect"
    "testing"
)

func TestStructure(t *testing.T) {
    t.Run("test of link", func(t *testing.T) {
        l := InitLink()
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
        if l.GetLinkLength() != 3 {
            t.Errorf("len of link is not 2, fail")
        }
        l.DeleteIndexNode(1)
        if delIndexNodeGetData, _ := l.GetIndexData(1); delIndexNodeGetData != 200 {
            t.Errorf("delete error, fail")
        }
        if !l.GetValue(200) {
            t.Errorf("get value fail")
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
    t.Run("test of hash map", func(t *testing.T) {
        hashmap := InitHashMap()

        tmpValue := "xxx"
        testValue := "22xx"
        for _, v := range []interface{}{1, 2, 22, "a", 22} {
            hashmap.Set(v, tmpValue)
        }
        // change value
        hashmap.Set(22, testValue)
        if hashmap.Get(22) != testValue {
            t.Errorf("get hash map value !=" + testValue + "fail")
        }
        // get value
        if hashmap.Get("a") != tmpValue {
            t.Errorf("get hash map value !=" + tmpValue + "fail")
        }
        // remove value
        if !hashmap.Remove(1) && hashmap.Get(1) != nil {
            t.Errorf("remove hash map value fail, fail")
        }
        // get keys
        if !reflect.DeepEqual(hashmap.Keys(), []interface{}{"a", 2, 22}) {
            t.Errorf("get hash map key fail")
        }
        // get value
        if !reflect.DeepEqual(hashmap.Values(), []interface{}{tmpValue, tmpValue, testValue}) {
            t.Errorf("get hash map value fail")
        }
        // empty
        hashmap.Clear()
        if hashmap.size != 0 || len(hashmap.Keys()) != 0 {
            t.Errorf("clear hash map fail")
        }
    })
    t.Run("test of set", func(t *testing.T) {
        set := InitSet()
        for _, v := range []interface{}{"a", "ab", 2, 2, 3, 5, 4} {
            set.AddValue(v)
        }
        if !reflect.DeepEqual(set.Order(), []interface{}{"a", "ab", 2, 3, 4, 5}) {
            t.Errorf("set add value fail")
        }
        set.RemoveValue(5)
        set.RemoveValue("a")
        if !reflect.DeepEqual(set.Order(), []interface{}{"ab", 2, 3, 4}) {
            t.Errorf("set remove value fail")
        }
        set.Clear()
        if len(set.Order()) != 0 {
            t.Errorf("clear set fail")
        }
    })
}

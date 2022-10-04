package diskstructure

import (
    "os"
    "reflect"
    "testing"
)

func TestDiskStructure(t *testing.T) {
    t.Run("test of disklink", func(t *testing.T) {
        Assert := func(err error) {
            if err != nil {
                t.Errorf(err.Error())
            }
        }

        firstStr := "abgpiou"
        secStr := "pboigdj"
        thirdStr := "owbhiruvf"
        fourthStr := "vwdfjkio"

        l := NewLink()

        err := l.Add(1, firstStr)
        if err != nil && err != AlReadyExistsErr {
            t.Errorf("test add %s fail", firstStr)
        }
        a, err := l.Get(1)
        if a != firstStr {
            t.Errorf("test add or get %s fail", firstStr)
        }
        err = l.Add(2, secStr)
        if err != nil && err != AlReadyExistsErr {
            t.Errorf("test add %s fail", secStr)
        }
        err = l.Add(3, thirdStr)
        if err != nil && err != AlReadyExistsErr {
            t.Errorf("test add %s fail", thirdStr)
        }
        b, _ := l.Get(2)
        if b != secStr {
            t.Errorf("test add %s fail", secStr)
        }
        err = l.Del(2)
        Assert(err)

        b, err = l.Get(2)
        if err != NotFoundErr {
            t.Errorf("test del or get %d fail", 2)
        }

        err = l.Add(2, secStr)
        Assert(err)

        err = l.Add(4, fourthStr)
        Assert(err)

        err = l.Set(4, "304958")
        Assert(err)

        err = l.Invert()
        Assert(err)

        keys, _ := l.Keys()
        if !reflect.DeepEqual(keys, []uint32{4, 2, 3, 1}) {
            t.Errorf("test get keys fail")
        }

        l.Close()
        os.Remove(DBFileName)
    })
}

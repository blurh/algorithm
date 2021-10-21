package strings

import (
    "reflect"
    "testing"
)

func TestTree(t *testing.T) {
    text := "dfndflnaasidfnadfn"
    //       ↑          ↑   ↑
    pattren := "dfn"

    assertMatch := func(matchResult []int) {
        if !reflect.DeepEqual([]int{0, 11, 15}, matchResult) {
            t.Errorf("match fail")
        }
    }
    t.Run("test of naive match", func(t *testing.T) {
        matchResult := Naive(text, pattren)
        assertMatch(matchResult)
    })
}

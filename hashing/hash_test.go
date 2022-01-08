package hashing

import (
    "testing"
)

func TestHash(t *testing.T) {
    str := "oapfjasasfn"
    t.Run("test for BKDR Hash", func(t *testing.T) {
        hash := BKDRHash(str)
        if hash != 1946341454 {
            t.Errorf("BKDR Hash fail")
        }
    })
}

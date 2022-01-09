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
    t.Run("test for DJB Hash", func(t *testing.T) {
        hash := DJB(str)
        if hash != 696472849 {
            t.Errorf("DJB Hash fail")
        }
    })
    t.Run("test for SDBM Hash", func(t *testing.T) {
        hash := SDBM(str)
        if hash != 1951602670 {
            t.Errorf("DJB Hash fail")
        }
    })
}

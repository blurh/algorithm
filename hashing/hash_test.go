package hashing

import (
    "testing"
)

func TestHash(t *testing.T) {
    str := "oapfjasasfn"
    t.Run("test for RK Hash", func(t *testing.T) {
        hash := RK(str)
        if hash != 0 {
            t.Errorf("RK Hash fail")
        }
    })
    t.Run("test for BKDR Hash", func(t *testing.T) {
        hash := BKDR(str)
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
    t.Run("test for RS Hash", func(t *testing.T) {
        hash := RS(str)
        if hash != 1955139542 {
            t.Errorf("RS Hash fail")
        }
    })
    t.Run("test for JS Hash", func(t *testing.T) {
        hash := JS(str)
        if hash != 1424885267 {
            t.Errorf("JS Hash fail")
        }
    })
    t.Run("test for AP Hash", func(t *testing.T) {
        hash := AP(str)
        if hash != 1251890205 {
            t.Errorf("AP Hash fail")
        }
    })
    t.Run("test for DEK Hash", func(t *testing.T) {
        hash := DEK(str)
        if hash != 87438148 {
            t.Errorf("DEK Hash fail")
        }
    })
    t.Run("test for FNV Hash", func(t *testing.T) {
        hash := FNV(str)
        if hash != 550411091 {
            t.Errorf("FNV Hash fail")
        }
    })
    t.Run("test for ELF Hash", func(t *testing.T) {
        hash := ELF(str)
        if hash != 16518942 {
            t.Errorf("ELF Hash fail")
        }
    })
    t.Run("test for PJW Hash", func(t *testing.T) {
        hash := PJW(str)
        if hash != 16518942 {
            t.Errorf("PJW Hash fail")
        }
    })
    t.Run("test for Jenkins Hash", func(t *testing.T) {
        num := ELF("apofn")
        hash := Jenkins(num)
        invertHash := JenkinsInvert(hash)
        if num != invertHash {
            t.Errorf("Jenkins Hash fail")
        }
    })
    t.Run("test for MurMur Hash", func(t *testing.T) {
        hash := MurMur(str, 1)
        if hash != 405891513 {
            t.Errorf("MurMur Hash fail")
        }
    })
    t.Run("test for SIP Hash", func(t *testing.T) {
        hash := SIP(str, 17390, 23405)
        if hash != 1899414923 {
            t.Errorf("SIP Hash fail")
        }
    })
    t.Run("test for Adler32 Hash", func(t *testing.T) {
        hash := Adler32(str)
        if hash != 456918157 {
            t.Errorf("Adler32 Hash fail")
        }
    })
    t.Run("test for CRC32 Hash", func(t *testing.T) {
        hash := CRC32(str)
        if hash != 1848985071 {
            t.Errorf("CRC32 Hash fail")
        }
    })
}

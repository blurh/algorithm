package strings

func RK(text, pattren string) []int {
    var matchResult []int
    hashNum := 0
    pattrenHash := 0
    for i := 0; i < len(pattren); i++ {
        hashNum += int(text[i])
        pattrenHash += int(pattren[i])
    }
    for position := 0; position < len(text)-len(pattren)+1; position++ {
        if hashNum == pattrenHash {
            for posOfPattren := 0; posOfPattren < len(pattren); posOfPattren++ {
                if pattren[posOfPattren] != text[position+posOfPattren] {
                    break
                }
                if posOfPattren == len(pattren)-1 {
                    matchResult = append(matchResult, position)
                }
            }
        }
        if position < len(text)-len(pattren) {
            hashNum = hashNum - int(text[position]) + int(text[position+len(pattren)])
        }
    }
    return matchResult
}

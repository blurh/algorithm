package strings

func Naive(text, pattren string) []int {
    var match []int
    for position := 0; position < len(text)-len(pattren)+1; position++ {
        for pattrenPos := 0; pattrenPos < len(pattren); pattrenPos++ {
            if text[position+pattrenPos] != pattren[pattrenPos] {
                break
            }
            // 到词末了没有 break 说明匹配
            if pattrenPos == len(pattren)-1 {
                match = append(match, position)
            }
        }
    }
    return match
}

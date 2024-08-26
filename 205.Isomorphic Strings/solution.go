func isIsomorphic(s string, t string) bool {

    mappingS := make(map[byte]byte)
    mappingT := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        charS := s[i]
        charT := t[i]

        if mappedChar, exists := mappingS[charS]; exists {
            if mappedChar != charT {
                return false
            }
        } else {
            mappingS[charS] = charT
        }


        if mappedChar, exists := mappingT[charT]; exists {
            if mappedChar != charS {
                return false
            }
        } else {
            mappingT[charT] = charS
        }
    }

    return true
}

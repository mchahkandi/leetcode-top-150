func wordPattern(pattern string, s string) bool {

    charToWord := make(map[byte]string)
    wordToChar := make(map[string]byte)

    fields := strings.Fields(s)

    if len(pattern) != len(fields) {
        return false
    }

    for index, val := range fields {
        pChar := pattern[index]


        if mappedWord, exists := charToWord[pChar]; exists {
            if val != mappedWord {
                return false
            }
        } else {
            charToWord[pChar] = val
        }


        if mappedChar, exists := wordToChar[val]; exists {
            if mappedChar != pChar {
                return false
            }
        } else {
            wordToChar[val] = pChar
        }
    }

    return true
}

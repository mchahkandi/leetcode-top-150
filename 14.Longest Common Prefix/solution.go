func longestCommonPrefix(strs []string) string {

    res := ""

    for i := 0 ; i < len(strs[0]) ; i++ {
        for j := 0 ; j < len(strs) ; j++ {
            if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
                return res
            }
        }
        res += string(strs[0][i])
    }

    return res
}

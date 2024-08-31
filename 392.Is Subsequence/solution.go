func isSubsequence(s string, t string) bool {

    i, j := 0, 0
    var res string

    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            res += string(s[i])
            i++; j++
        } else {
            for j < len(t) && s[i] != t[j] { j++ }
        }
    }

    return res == s
}

func isAnagram(s string, t string) bool {
    s = sortString(s)
    t = sortString(t)

    if s == t {
        return true
    }

    return false

}

func sortString(word string) string {
    s := []rune(word)
    sort.Slice(s, func(i int,j int) bool { return s[i] < s[j] })
    return string(s)
}

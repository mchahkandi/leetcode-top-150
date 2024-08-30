func groupAnagrams(strs []string) [][]string {

    groups := map[string][]string{}

    for i := 0 ; i < len(strs) ; i++ {
        val := sorted(strs[i])

        groups[val] = append(groups[val], strs[i])

    }

    result := [][]string{}

    for _ , value := range groups {
        result = append(result,value)
    }

    return result
}

func sorted(word string) string {
    s := []rune(word)
    sort.Slice(s, func(i int,j int) bool { return s[i] < s[j] })
    return string(s)
}

func lengthOfLongestSubstring(s string) int {

    maxl := 0

    for i := 0 ; i< len(s) ; i++ {

        hashMap :=  map[byte]bool{}

        for j := i ; j < len(s) ; j++ {
            _ , exists := hashMap[s[j]]

            if exists {
                break
            }else {
                hashMap[s[j]] = true
                maxl = max(maxl, j - i + 1)
            }
        }
    }
    return maxl
}

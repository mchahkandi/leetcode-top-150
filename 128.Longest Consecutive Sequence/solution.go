func longestConsecutive(nums []int) int {

    hashMap := map[int]bool{}

    for _, val := range nums {
        hashMap[val] = true
    }

    maxLength := 0

    for val := range hashMap {

        if _, ok := hashMap[val-1]; !ok {

            currLength := 1

            for i := val + 1; ; i++ {
                if _, exists := hashMap[i]; !exists {
                    break
                }
                currLength++
            }
            maxLength = max(maxLength, currLength)
        }
    }
    return maxLength
}

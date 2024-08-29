func containsNearbyDuplicate(nums []int, k int) bool {

    hashMap := map[int]int{}

    for index, num := range nums {
        key, exists := hashMap[num]

        if !exists {
            hashMap[num] = index
        }else if index - key <= k {
            return true
        }else{
            hashMap[num] = index
        }
    }

    return false
}

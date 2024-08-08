func majorityElement(nums []int) int {
    hashTable := map[int]int{}

    for i:= 0 ; i < len(nums) ; i++ {
        if _,ok := hashTable[nums[i]] ; ok {
            hashTable[nums[i]] += 1
        }else{
            hashTable[nums[i]] = 1
        }
        
    }

    number := 0
    majorityvalue := 0

    for index, item := range(hashTable) {
        if item > number {
            number = item
            majorityvalue = index
        }
    }

    return majorityvalue
}

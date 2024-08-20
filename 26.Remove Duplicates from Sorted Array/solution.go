func removeDuplicates(nums []int) int {

    pointer := 0

    for i := 1 ; i < len(nums) ; i++ {
        if nums[pointer] != nums[i]  {
            nums[pointer + 1] = nums[i]
            pointer++
        }
    }

    return pointer + 1

}

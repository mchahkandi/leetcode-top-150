func canJump(nums []int) bool {

    jumps := 0

    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] > jumps {
            jumps = nums[i]
        }
        if jumps == 0 && i != len(nums) -1 {
            return false
        }
        jumps--
    }

    return true
}

func rotate(nums []int, k int)  {

    n := len(nums)
    k = k % n

    if k == 0 {
        return
    }

    for i,j := 0, len(nums) - k - 1  ; i < j ; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }

    for i,j := len(nums) - k  , len(nums) - 1 ; i < j ; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }

    for i,j := 0, len(nums) - 1 ; i < j ; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}

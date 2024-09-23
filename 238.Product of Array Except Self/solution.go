func productExceptSelf(nums []int) []int {
    out := []int{}

    prefix := 1
    for i := 0 ; i < len(nums) ; i++ {
        out = append(out,prefix)
        prefix *= nums[i]
    }

    postfix := 1
    for j := len(nums) - 1 ; j >=0 ; j-- {
        out[j] *= postfix
        postfix *= nums[j]
    }
    return out
}

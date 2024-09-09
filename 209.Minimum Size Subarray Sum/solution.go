func minSubArrayLen(target int, nums []int) int {
    mint := math.MaxInt32
    sum := 0
    j := 0

    for i := 0 ; i < len(nums) ; i++ {

        sum += nums[i]
        for sum >= target {
            mint = min(mint,i - j + 1)
            sum -= nums[j]
            j++
        }

    }
    if mint == math.MaxInt32 {
        return 0
    }else{
        return mint
    }
}

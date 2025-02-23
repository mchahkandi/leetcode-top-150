func maxSubArray(nums []int) int {

    maxSum := nums[0]
    currSum := 0

    for _,n := range nums {
        currSum = max(currSum,0)
        currSum += n
        maxSum = max(maxSum, currSum)
    }

    return maxSum
}

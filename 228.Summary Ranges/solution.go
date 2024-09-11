func summaryRanges(nums []int) []string {
    i := 0
    res := []string{}

    for i < len(nums) {
        j := i
        for j + 1 < len(nums) && nums[j + 1] == nums[j] + 1 {
            j = j + 1
        }
        if i != j {
            res = append(res, fmt.Sprintf("%d->%d", nums[i],nums[j]))
        }else{
            res = append(res, fmt.Sprintf("%d",nums[i]))
        }
        i = j+1
    }
    return res
}

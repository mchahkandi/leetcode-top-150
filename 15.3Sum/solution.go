func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	result := make([][]int, 0)

	for i := 0; i < len(nums) - 2; i++ {

		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		target := -nums[i]
		j := i+1
		k := len(nums) - 1

		for j < k {
			if j > i + 1 && nums[j] == nums[j - 1] {
				j++
				continue
			}
			if k < len(nums) - 1 && nums[k] == nums[k + 1] {
				k--
				continue
			}
			if nums[j] + nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}else{
				if nums[j] + nums[k] > target {
					k--
				}else{
					j++
				}
			}
		}
	}
	return result
}

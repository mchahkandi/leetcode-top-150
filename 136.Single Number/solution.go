func singleNumber(nums []int) int {
    num := 0
    for _,val := range nums {
        num = val ^ num
    }
    return num
}

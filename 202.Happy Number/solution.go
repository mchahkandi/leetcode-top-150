func isHappy(n int) bool {

    hashMap := map[int]bool{}

    for n != 1 {
        sum := 0
        str := strconv.Itoa(n)
        for i := 0 ; i < len(str) ; i++ {
            digit := int(str[i] - '0')
            sum += digit * digit
        }
        if _, exists := hashMap[sum] ; exists {
            return false
        }else{
            hashMap[sum] = true
        }
        n = sum
    }
    return true
}

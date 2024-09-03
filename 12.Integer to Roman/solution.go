func intToRoman(num int) string {

    maps := map[int]string{ 1000 : "M", 900 : "CM", 500 : "D", 400 : "CD", 100 : "C", 90 : "XC", 50 : "L", 40 : "XL", 10 : "X", 9 : "IX", 5 : "V", 4 : "IV", 1 : "I",}

    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

    str := ""

        for _, val := range values {
            for num >= val {
                num -= val
                str += maps[val]
            }
        }
    return str
}

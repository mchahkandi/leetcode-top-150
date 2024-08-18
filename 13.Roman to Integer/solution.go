func romanToInt(s string) int {

    pairs := map[byte]int{
        'I' : 1,
        'V' : 5,
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }

    sum := 0

    for i := 0 ; i < len(s) - 1 ; i++ {

        key, _ := pairs[s[i]]

        if pairs[s[i]] < pairs[s[i + 1]] {
            sum -= key
        }else{
            sum += key
        }

    }

    sum += pairs[s[len(s) - 1]]

    return sum
}

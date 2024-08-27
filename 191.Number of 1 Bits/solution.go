func hammingWeight(n int) int {

    count := 0

    for n > 0 {
        remainder := n % 2

        if remainder != 0 {
            count++
        }

        n = n/2
    }

    return count

}

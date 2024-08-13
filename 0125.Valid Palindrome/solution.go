func isPalindrome(s string) bool {
    
    wiped := removeNonAlphanumeric(s)

    leftPointer := 0
    rightPointer := len(wiped) - 1

    for i := leftPointer ; i <= rightPointer ; i++ {
        if wiped[leftPointer] == wiped[rightPointer] {
            leftPointer++
            rightPointer--
        }else{
            return false
        }
    }

    return true
}

func removeNonAlphanumeric(input string) string {
	var builder strings.Builder

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}

	return strings.ToLower(builder.String())
}

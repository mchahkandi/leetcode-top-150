func letterCombinations(digits string) []string {

	hashMap := map[rune][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	res := []string{}

	if len(digits) == 0 {
		return res
	}

	var recursive func(i int, str string)

	recursive = func(i int, str string) {
		if i == len(digits) {
			res = append(res, str)
			return
		}

		for _, val := range hashMap[rune(digits[i])] {
			recursive(i+1, str+string(val))
		}
	}

	recursive(0, "")
	return res
}

func canConstruct(ransomNote string, magazine string) bool {
   
   ransomTable := map[rune]int{}
   magazineTable := map[rune]int{}


   for _, item := range ransomNote {
        ransomTable[item]++
   }

   for _, str := range magazine {
        magazineTable[str]++
   }

    for char, frequency := range ransomTable {
		if magazineTable[char] < frequency {
			return false
		}
	}

	return true


}

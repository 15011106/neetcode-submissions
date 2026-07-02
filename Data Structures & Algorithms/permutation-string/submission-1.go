func checkInclusion(s1 string, s2 string) bool {
	
	stringMap := make(map[rune]int)
	comparisonMap := make(map[rune]int)
	long := ""
	short := ""

	if len(s1) > len(s2){
		return false
	}

			long = s2
		short = s1
		for _ ,v := range s1{
				stringMap[v]++
			}
		for i:=0 ;i <len(s1); i++{
			comparisonMap[rune(long[i])]++
		}

if compareMaps(stringMap, comparisonMap) {
    return true
}
	l := 0
	for r :=len(short); r < len(long); r++{
		
		comparisonMap[rune(long[l])]--
		l++
		comparisonMap[rune(long[r])]++
		
		if compareMaps(stringMap, comparisonMap){
		return true
		}
	}

	return false
}

func compareMaps(a, b map[rune]int) bool {

    for k, v := range a {
        if bv, ok := b[k]; !ok || bv != v {
            return false 
        }
    }
    return true 
}

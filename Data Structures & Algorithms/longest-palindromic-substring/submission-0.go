func longestPalindrome(s string) string {

	ans := ""
	for i:=0; i<len(s); i++{
		oddAns := expend(s, i, i)
		if len(oddAns) >= len(ans){
			ans = oddAns
		}

		evenAns := expend(s, i, i+1)
		if len(evenAns) >= len(ans){
			ans = evenAns
		}
	}    

	return ans
}

func expend(s string, left, right int) string{
	
	for left >=0 && right < len(s){
		if s[left] == s[right]{
			left--
			right++
		}else{
			break
		}
	}
	return string(s[left+1:right])
}

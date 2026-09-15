func isPalindrome(s string) bool {

s = strings.ToLower(s)
var sb strings.Builder

for _,c := range s{
	if ('a'<= c && c<='z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9'){
		sb.WriteString(string(c))
	}
}

s = sb.String()

l := 0
h := len(s)-1
for l < h{
    if s[l] != s[h]{
      return false
    }
      l++
      h--
}

return true
}
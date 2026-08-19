func letterCombinations(digits string) []string {

var keyboards map[byte]string
keyboards = make(map[byte]string)
keyboards['2'] = "abc"
keyboards['3'] = "def"
keyboards['4'] = "ghi"
keyboards['5'] = "jkl"
keyboards['6'] = "mno"
keyboards['7'] = "pqrs"
keyboards['8'] = "tuv"
keyboards['9'] = "wxyz"

ans := []string{}

var dfs func(idx int, str string)
dfs = func(idx int, str string){
    if idx == len(digits){
      ans = append(ans, str)
        return 
    }

    for i:=0;i<len(keyboards[digits[idx]]); i++{
      str += string(keyboards[digits[idx]][i])
      dfs(idx+1, str)
      str = str[:len(str)-1]
  }  
}

dfs(0, "")

if len(digits) == 0 {
		return []string{}
}

return ans
}


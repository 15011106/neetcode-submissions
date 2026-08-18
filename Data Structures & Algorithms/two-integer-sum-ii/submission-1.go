
func twoSum(numbers []int, target int) []int {
   
s := 0
l := len(numbers)
ans := []int{}
for s<l{

  if numbers[s] + numbers[l-1] == target{
  ans = append(ans, s+1)
  ans = append(ans, l)
  break
}
  if numbers[s] + numbers[l-1] < target{
      s++
  }else{
      l--
}
}

return ans
}

// 3Sum
// https://neetcode.io/problems/three-integer-sum

func threeSum(nums []int) [][]int {

    sort.Slice(nums, func(i, j int)bool {
       return nums[i] < nums[j]
    })
    
    ans :=[][]int{}
    
    for i:=0 ; i<=len(nums)-3; i++{
        if i > 0 && nums[i] == nums[i-1]{
            continue
            }
        
        l := i+1
        h := len(nums)-1
        
        for l < h{
            
            sum := nums[l]+nums[h] + nums[i]
        
            
            if sum == 0{
                ans = append(ans, []int{nums[i], nums[l], nums[h]})
                
                l++
                h--
                
                for l < h && nums[l] == nums[l-1]{
                    l++
                    }
                    
                for l < h && nums[h] == nums[h+1]{
                    h--
                    }
                      
                }else if sum > 0{
                    h--
                    }else{
                    l++       
                    }
            }
        }
        
        return ans
}
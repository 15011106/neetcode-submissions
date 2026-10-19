# Products of Array Except Self
# https://neetcode.io/problems/products-of-array-discluding-self

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        
        
        n = len(nums)
        left = [1] * n
        right = [1] * n 
        ans = [1] * n
        left[0] = nums[0]
        
        right[n-1] = nums[n-1]
        
        for i in range(1,n):
            left[i] = left[i-1]* nums[i]
       
        for i in range(n-2, -1, -1):
           right[i] = right[i+1] * nums[i]
        
        for i in range(1, n-1):
           ans[i] = left[i-1] * right[i+1]             
       
        ans[0] = right[1]
        ans[n-1] = left[n-2]    
        
        return ans

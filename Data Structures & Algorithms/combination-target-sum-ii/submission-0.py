# Combination Sum II
# https://neetcode.io/problems/combination-target-sum-ii

class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        ans = []
        
        def dfs(start,total, path):
            
            if total == target:
                ans.append(path.copy())
                return
                
            if total > target:
                return
                
            for i in range (start, len(candidates)):
                if i> start and candidates[i] == candidates[i-1]:
                    continue        
                path.append(candidates[i])            
                dfs(i+1, total+candidates[i], path) 
                path.pop()
        dfs(0,0,[])
        return ans
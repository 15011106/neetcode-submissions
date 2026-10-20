# Number of Connected Components in an Undirected Graph
# https://neetcode.io/problems/count-connected-components

class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        
    
        parents = [i for i in range(n+1)]
        size = [1] * (n+1)    
        
        def union(a, b):
            rootA = find(a)
            rootB = find(b)
            if rootA == rootB:
                return False
            
            if size[rootA] < size[rootB]:
                rootA, rootB = rootB, rootA
            parents[rootB] = rootA
            size[rootA] += size[rootB]
            return True
        
        def find(a):
            if parents[a] != a:
                parents[a] = find(parents[a])
            return parents[a]                
        
        ans = n
        for a,b in edges:
            if union(a,b):
                ans -=1
        return ans
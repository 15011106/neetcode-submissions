# Network Delay Time
# https://neetcode.io/problems/network-delay-time

import heapq

class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        inf = 1<<60
        dist = [inf] * (n+1)
        dist[k] = 0
        graph = [[] for i in range(n+1)]
        for u,v,t in times:
            graph[u].append((v,t))
        
        pq = [(0, k)] 
        while pq:
            d, cur = heapq.heappop(pq)
            if d > dist[cur]:
                continue
            
            for next, v in graph[cur]:
                tempDist = d + v
                if tempDist < dist[next]:
                    dist[next] = tempDist
                    heapq.heappush(pq, (tempDist, next))
        
        
        ans = max(dist[1:])
        if ans == inf:
            return -1
        
        
        return ans

                
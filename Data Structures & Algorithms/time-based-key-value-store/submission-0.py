# Time Based Key-Value Store
# https://neetcode.io/problems/time-based-key-value-store

class TimeMap:

    def __init__(self):
        self.data:dict[str, list[tuple[int, str]]] = {}

    def set(self, key: str, value: str, timestamp: int) -> None:
        
        if key not in self.data:
            self.data[key] = [(timestamp, value)]
        else:
            self.data[key].append((timestamp, value))
            
    def get(self, key: str, timestamp: int) -> str:
        if key not in self.data:
            return ""
        low = 0
        high = len(self.data[key])-1
        result = ""
        
        while low <= high :
            mid = (low+high)//2
            if self.data[key][mid][0] == timestamp:
                return self.data[key][mid][1]
            elif self.data[key][mid][0] > timestamp:
                  high = mid-1
            else :
                 result = self.data[key][mid][1] 
                 low = mid+1
        return result

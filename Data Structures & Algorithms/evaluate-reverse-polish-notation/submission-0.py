# Evaluate Reverse Polish Notation
# https://neetcode.io/problems/evaluate-reverse-polish-notation

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack =[]
        sum = 0
        
        for i in tokens:
            if i not in ('+','-','*','/') :
                stack.append(int(i))
            else:
                a = stack.pop()
                b = stack.pop()
                temp = 0
                
                if i == '+':
                    temp = b+a
                    stack.append(temp)
                elif i == '-':
                    temp = b-a
                    stack.append(temp)
                elif i == '*':
                    temp = b*a
                    stack.append(temp)
                else:
                    temp = int(b/a)
                    stack.append(temp)
        sum = stack.pop()
        return sum
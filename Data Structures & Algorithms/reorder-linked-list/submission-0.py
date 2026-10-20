# Reorder Linked List
# https://neetcode.io/problems/reorder-linked-list

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        fast = head
        slow = head
        
        
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        
        second = slow.next
        slow.next = None
        
        prev = None
        
        while second:
            temp = second.next
            second.next = prev
            prev = second
            second = temp
        
        first = head
        second = prev
        
        while second:
            firstTemp = first.next
            secondTemp = second.next
            
            first.next = second
            second.next = firstTemp
            
            first = firstTemp
            second = secondTemp
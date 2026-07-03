func reverseList(head *ListNode) *ListNode {

curNode := head
var prevNode *ListNode
prevNode = nil

for curNode != nil{
nextNode := curNode.Next
curNode.Next = prevNode

prevNode = curNode
curNode = nextNode
}

return prevNode
}



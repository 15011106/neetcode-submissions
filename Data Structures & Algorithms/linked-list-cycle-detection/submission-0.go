/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

	visited := make(map[*ListNode]bool)
	curNode := head
	visited[curNode] = true

	for curNode != nil{

		if curNode.Next != nil{
			if _, ok:= visited[curNode.Next]; ok{
				return true
			}else{
				visited[curNode.Next] = true
				curNode = curNode.Next
			}
		}else{
			break
	}
}
	return false
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {

ans := []int{}
if root == nil{
	return ans
}

queue := []*nodeWithDepth{}
queue = append(queue, &nodeWithDepth{treeNode:root, depth:0})


var bfs func(queue []*nodeWithDepth, depth int)
bfs = func(queue []*nodeWithDepth, depth int){

for len(queue) > 0{

cur := queue[0]

if cur.depth == len(ans){
ans = append(ans, cur.treeNode.Val)
}else{
ans[cur.depth] = cur.treeNode.Val
}

queue = queue[1:]

if cur.treeNode.Left != nil{
queue= append(queue, &nodeWithDepth{treeNode : cur.treeNode.Left, depth : cur.depth+1})}



if cur.treeNode.Right !=nil{
queue = append(queue, &nodeWithDepth{treeNode: cur.treeNode.Right, depth : cur.depth+1})}
}
}

bfs(queue, 0)

return ans

}



type nodeWithDepth struct{
treeNode *TreeNode
depth int

}
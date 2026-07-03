/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	var dfs func(root *TreeNode,depth int)
	dfs = func (root *TreeNode, depth int){
if root == nil{
return
}
if len(ans) == depth{
	ans = append(ans, []int{})
}
ans[depth] = append(ans[depth], root.Val)
dfs(root.Left, depth + 1)
dfs(root.Right, depth + 1)
}
	dfs(root, 0)

	return ans
}


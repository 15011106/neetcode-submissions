// Invert Binary Tree
// https://neetcode.io/problems/invert-a-binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    var dfs func(*TreeNode)

    dfs = func(node *TreeNode) {

        if node.Left == nil && node.Right == nil {
            return
        }
        
        if node.Left != nil && node.Right != nil{
            node.Left, node.Right = node.Right, node.Left
        }else if node.Left == nil{
            node.Left = node.Right
            node.Right = nil
        }else{
            node.Right = node.Left
            node.Left = nil
            }
        
        if node.Left != nil{
            dfs(node.Left)
            }
        if node.Right != nil{
            dfs(node.Right)
       }
    }
    
    if root == nil{
        return root
        }
    
    dfs(root)
    return root
}
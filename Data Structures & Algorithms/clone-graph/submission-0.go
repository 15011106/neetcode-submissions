/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    cloned := make(map[*Node]*Node)

	    var dfs func(*Node) *Node
		dfs = func(current *Node) *Node {
        
		if current == nil {
            return nil
        }

        if copy, ok := cloned[current]; ok {
            return copy
        }
		cloned[current] = &Node{Val: current.Val}
		
        for _, neighbor := range current.Neighbors {
            cloned[current].Neighbors = append(
                cloned[current].Neighbors,
                dfs(neighbor),
            )
        }

        return cloned[current]
    }

	ans := dfs(node)

    return ans

}

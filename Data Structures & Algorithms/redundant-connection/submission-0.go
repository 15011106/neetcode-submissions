// Redundant Connection
// https://neetcode.io/problems/redundant-connection

func findRedundantConnection(edges [][]int) []int {


    parents := make([]int,len(edges)+1)
    for i:=0; i<=len(edges);i++{
        parents[i ]= i
        }
    
    var union func([]int,int, int)
   var find func([]int, int) int    
     
   find = func(parents []int,i int) int{
       
       
       if parents[i] != i{
           parents[i] = find(parents, parents[i])
           }

       
       return parents[i]
       }


   union = func(parents []int, i,j int){
       i = find(parents, i)
       j = find(parents, j)
       
       if i<j{
           parents[j] = i
       }else{
           parents[i] = j
           }
   }
   
   ans := make([]int,2)
   for i:=0 ;i<len(edges);i++{
       if find(parents, edges[i][0]) == find(parents, edges[i][1]){
           ans[0] = edges[i][0]
           ans[1] = edges[i][1]
           }
       union(parents, edges[i][0], edges[i][1])
       
       }
       
       return ans[:]
}

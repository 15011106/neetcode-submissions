
func uniquePaths(m int, n int) int {

	arr := make([][]int, m)
	for i:=0; i<m;i++{
		arr[i] = make([]int, n)
}

	for i:=0; i<m; i++{
		if i == 0 {
			for j:=0 ;j< n;j++{
				arr[0][j] = 1
			}
		}
		arr[i][0] = 1
	}

	for i:=1; i <m ;i++{
		for j:=1; j<n ;j++{
			arr[i][j] = arr[i-1][j]+ arr[i][j-1]
		}
	}
	
	return arr[m-1][n-1]
}


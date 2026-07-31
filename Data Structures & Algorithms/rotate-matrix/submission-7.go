

func rotate(matrix [][]int)  {

n := len(matrix)

for i:=0; i<n/2;i++{
   for j:=i; j<(n-i)-1; j++{

    temp := matrix[j][n-i-1]
    
    // first loop
    matrix[j][n-i-1] = matrix[i][j]
    // second
    matrix[n-i-1][n-j-1], temp = temp, matrix[n-i-1][n-j-1]
    // third
    matrix[n-j-1][i], temp = temp, matrix[n-j-1][i]
    // fourth
    matrix[i][j] = temp
    
  }
}
}

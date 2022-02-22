package top100

import "fmt"

/*
n x n 的二维矩阵，我们需要将其顺时针旋转 90 度

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/

func Rotate(matrix [][]int) {
	/*
		可以比较简单的找到规律:
			matrix[i][j] = matrix[n-j-1][i]

		但是直接替换的话，会被覆盖掉，又要求只能在原地替换

		2.直接旋转，为避免保存的复杂，直接四个一组这样交换
	*/
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
			 matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1],matrix[n-i-1][n-j-1]
		}
	}

}



func Rotate2(matrix [][]int) {
	/*
		可以比较简单的找到规律，matrix[i][j] = matrix[n-j-1][i]
		但是直接替换的话，会被覆盖掉，又要求只能在原地替换

		1.先转置，再翻转每一行
	*/
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	fmt.Println(matrix)

	// 反转每一行
	for i := 0; i < n; i++ {
		fmt.Println(matrix[i])
		reverse(matrix[i])
	}
}

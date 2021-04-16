//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？
//
//
//
// 示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//
// 示例 2：
//
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//
//
// 示例 3：
//
//
//输入：m = 7, n = 3
//输出：28
//
//
// 示例 4：
//
//
//输入：m = 3, n = 3
//输出：6
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 109
//
// Related Topics 数组 动态规划
// 👍 950 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1];
}

func main() {
	res := uniquePaths(7,3)
	fmt.Println(res)
}

/**
题解：
动态规划：
	由题可知：机器人每次只能向下和向右移动，所以到达finish的所有路径 = 到达finish的上方方格路径 + 到达finish的左边方格路径
	到达finish的所有路径 记为f[i][j] 可得 f[i][j] = f[i][j-1] + f[i-1][j]
	因f[0][j]为方格的横向边缘，到达路径只能左边方格，所以是1 和 f[i][0] 为方格纵向边缘，到达路径只能上方方格，所以也是1
	即f[0][j] = 1 f[i][0] = 1
	因为f[0][0]为出发点，所以 f[0][0] = 1
	所以 可根据f[i][j] = f[i][j-1] + f[i-1][j] 得出到达finish的所有路径和

复杂度分析：
	时间复杂度：O(m*n)
	空间复杂度：O(m*n)
*/
//leetcode submit region end(Prohibit modification and deletion)

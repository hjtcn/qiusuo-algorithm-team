//这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。 
//
// 如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。 
//
// 请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：grid = [[1,0],[0,1]]
//输出：0
//解释：没有一台服务器能与其他服务器进行通信。 
//
// 示例 2： 
//
// 
//
// 输入：grid = [[1,0],[1,1]]
//输出：3
//解释：所有这些服务器都至少可以与一台别的服务器进行通信。
// 
//
// 示例 3： 
//
// 
//
// 输入：grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
//输出：4
//解释：第一行的两台服务器互相通信，第三列的两台服务器互相通信，但右下角的服务器无法与其他服务器通信。
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m <= 250 
// 1 <= n <= 250 
// grid[i][j] == 0 or 1 
// 
// Related Topics 图 数组 
// 👍 22 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func countServers(grid [][]int) int {
	col_arr := map[int]int{}
	row_arr := map[int]int{}
	connect_num := 0

	for row_key, row_vals := range grid {
		for col_key, val := range row_vals {
			if val == 1 {
				row_arr[row_key]++
				col_arr[col_key]++
			}
		}
	}

	for row_key, row_vals := range grid {
		for col_key, val := range row_vals {
			if val == 1 && (row_arr[row_key] > 1 || col_arr[col_key] > 1) {
				connect_num++
			}
		}
	}

	return connect_num
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {
	grid := [][]int{{1,0},{1,1}}
	res := countServers(grid)

	fmt.Println(res)
}

/**
题解：统计参与通信的服务器

耗时：72ms
内存：7M

要保证某个节点上的服务器能与其他服务器通信的条件：为这台服务器的所在节点的 行服务器数 > 1 或 列服务器数 > 1即可

若节点(r, w)存在服务器，则在row_arr[r]++, col_arr[c]++, 然后把grid每行每列遍历一下

然后在重新把grid遍历一遍，当查到一台服务器(r, w)，判断该服务器的行或列是否存在其他的服务器(即：row_arr[r] > 1 或 col_arr[c] > 1)，
	若成立则，能与其他服务器通信的服务器数量 +1
	若不成立则，继续向下遍历

复杂度分析：
	时间复杂度：O(m*n)
	空间复杂度：O(m+n)
 */

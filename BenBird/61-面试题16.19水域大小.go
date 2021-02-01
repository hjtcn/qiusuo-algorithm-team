//你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指
//相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
// 示例：
// 输入：
//[
//  [0,2,1,0],
//  [0,1,0,1],
//  [1,1,0,1],
//  [0,1,0,1]
//]
//输出： [1,2,4]
//
// 提示：
//
// 0 < len(land) <= 1000
// 0 < len(land[i]) <= 1000
//
// Related Topics 深度优先搜索 广度优先搜索
// 👍 39 👎 0
package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func pondSizes(land [][]int) []int {
	row, col := len(land), len(land[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			visited[i] = append(visited[i], false)
		}
	}
	res := []int{}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if land[r][c] == 0 && !visited[r][c] {
				t := 0
				dfSearch(land, visited, r, c, &t)
				res = append(res, t)
			}
		}
	}

	sort.Ints(res)
	return res
}

func dfSearch(land [][]int, visited [][]bool, r, c int, time *int) {
	if r >= len(land) || r < 0 || c >= len(land[0]) || c < 0 {
		return
	}

	if visited[r][c] {
		return
	}

	if land[r][c] != 0 {
		return
	}

	if land[r][c] == 0 {
		visited[r][c] = true
		*time += 1
		dfSearch(land, visited, r - 1, c - 1, time)	//左上
		dfSearch(land, visited, r + 1, c + 1, time)	//右下
		dfSearch(land, visited, r - 1, c + 1, time)	//右上
		dfSearch(land, visited, r + 1, c - 1, time)	//左下
		dfSearch(land, visited, r - 1, c - 1, time)	//上
		dfSearch(land, visited, r, c + 1, time)	//右
		dfSearch(land, visited, r + 1, c, time)	//下
		dfSearch(land, visited, r, c - 1, time)	//左
	}
}

func main() {
	land := [][]int {
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}
	res := pondSizes(land)

	fmt.Println(res)
}
//leetcode submit region end(Prohibit modification and deletion)

/**
题解：
DFS：深度优先搜索，遍历矩阵中所有元素，遇到水域时，就开始递归遍历该位置的八个方向是否存在水域，遇到水域时，记录水域个数并将该水域位置标记已访问
	复杂度分析：
		时间复杂度：O(m*n) 最坏情况下每个节点遍历两次
		空间复杂度：O(n)
*/

//3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。
//
// 给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。
//
//
//
// 示例：
//
// 输入: [[4,3,8,4],
//      [9,5,1,9],
//      [2,7,6,2]]
//输出: 1
//解释:
//下面的子矩阵是一个 3 x 3 的幻方：
//438
//951
//276
//
//而这一个不是：
//384
//519
//762
//
//总的来说，在本示例所给定的矩阵中只有一个 3 x 3 的幻方子矩阵。
//
//
// 提示:
//
//
// 1 <= grid.length <= 10
// 1 <= grid[0].length <= 10
// 0 <= grid[i][j] <= 15
//
// Related Topics 数组
// 👍 38 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func numMagicSquaresInside(grid [][]int) int {
	row := []int{-1,-1,-1,0,1,1,1,0}	//相对于中心元素周围8个数的顺时针的相对row位置
	col := []int{-1,0,1,1,1,0,-1,-1}	//相对于中心元素周围8个数的数时针的相对col位置
	count := 0

	for r := 1; r < len(grid) - 1; r++ {
		for c := 1; c < len(grid[0]) - 1; c++ {
			if grid[r][c] == 5 {
				around := []int{}
				for k := 0; k < 8; k++ {
					around = append(around, grid[r + row[k]][c + col[k]])
				}

				if isMagic(around) {
					count++
				}
			}
		}
	}

	return count
}

/**
将现获取的矩阵和标准幻方进行对比
 */
func isMagic(around []int) bool {
	var m = []int{8,1,6,7,2,9,4,3,8,1,6,7,2,9,4,3}
	for i := 0; i < 8; i += 2 {
		if m[i] == around[0] {
			flag_posit := true
			flag_rever := true
			for j := 0; j < 8; j++ {
				if around[j] != m[i+j] {
					flag_posit = false
				}

				if around[j] != m[i+8-j] {
					flag_rever = false
				}

				if !flag_posit && !flag_rever {
					return false
				}
			}

			return true
		}
	}

	return false
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {
	grid := [][]int{{4,3,8,4},{9,5,1,9},{2,7,6,2}}
	res := numMagicSquaresInside(grid)

	fmt.Println(res)
}

/**
题解：矩阵中的幻方

耗时：0ms
内存：2.2M

符合条件的幻方只有下面8种
8 1 6 	6 1 8 	2 9 4 	4 9 2
3 5 7 	7 5 3 	7 5 3 	3 5 7
4 9 2 	2 9 4 	6 1 8 	8 1 6

8 3 4 	6 7 2 	2 7 6 	4 3 8
1 5 9 	1 5 9 	9 5 1 	9 5 1
6 7 2 	8 3 4 	4 3 8 	2 7 6

要想满足每一行、每一列、两个对角线的和相等，则和都等于15，且中心数字必是5
再则就是这些满足条件的幻方，其实都可以有一种幻方进行旋转可得到，
所以只记录满足条件的顺时针数据为 81672943

所以首先循环遍历grid，查找中心元素为5的坐标，
	若发现存在中心元素为5的3x3矩阵，则获取改矩阵外围数时针数字，并以此存入切片around

将获取3x3矩阵和标准幻方进行对比
	若和标准幻方相等，则count++
	若无幻方与其相等，则继续遍历下一个中心元素

复杂度分析：
	时间复杂度：O(r*c)
	空间复杂度：O(1)
 */
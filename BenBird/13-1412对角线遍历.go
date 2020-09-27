//给你一个列表 nums ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，按顺序返回 nums 中对角线上的整数。
//
//
//
// 示例 1：
//
//
//
// 输入：nums = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,4,2,7,5,3,8,6,9]
//
//
// 示例 2：
//
//
//
// 输入：nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
//输出：[1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
//
//
// 示例 3：
//
// 输入：nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
//输出：[1,4,2,5,3,8,6,9,7,10,11]
//
//
// 示例 4：
//
// 输入：nums = [[1,2,3,4,5,6]]
//输出：[1,2,3,4,5,6]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10^5
// 1 <= nums[i].length <= 10^5
// 1 <= nums[i][j] <= 10^9
// nums 中最多有 10^5 个数字。
//
// Related Topics 排序 数组
// 👍 18 👎 0
package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findDiagonalOrder(nums [][]int) []int {
	m := make(map[int][]int, 0)
	res := []int{}
	max := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums[i]) - 1; j >= 0; j-- {
			if m[i + j] == nil {
				m[i + j] = []int{nums[i][j]}
			} else {
				m[i + j] = append(m[i + j], nums[i][j])
			}

			if (i + j) > max {
				max = i + j
			}
		}
	}

	for k := 0; k <= max; k++ {
		res = append(res, m[k]...)
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {
	//nums := [][]int{{1,2,3},{6},{7,8,9}}
	nums := [][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}}
	res := findDiagonalOrder(nums)

	fmt.Println(res)
}

/**
题解：对角线遍历

耗时：232ms
内存：22.4M

通过规律可发现每个对角线上纵横坐标和是相等的
所以在遍历nums时，将nums[i][j] 存入map m[i+j][] 中

因为Go没有自带的在数组/切片/map 头部插入元素的内置函数
所以为解决对m 展开遍历获取返回值中对角线元素排序的问题，在遍历nums时，是从最后一个元素开始向第一个元素进行遍历

所以，最后直接对m 进行遍历即可获取正确的返回值

复杂度分析：
	时间复杂度：O(C*R)
	空间复杂度：O(C*R)
 */

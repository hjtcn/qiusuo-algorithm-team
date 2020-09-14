//给定一个矩阵 A， 返回 A 的转置矩阵。
//
// 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
//
//
//
// 示例 1：
//
// 输入：[[1,2,3],[4,5,6],[7,8,9]]
//输出：[[1,4,7],[2,5,8],[3,6,9]]
//
//
// 示例 2：
//
// 输入：[[1,2,3],[4,5,6]]
//输出：[[1,4],[2,5],[3,6]]
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 1000
// 1 <= A[0].length <= 1000
//
// Related Topics 数组
// 👍 104 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

/**
one
耗时：20ms
内存：6.2M
**/
func transpose_one(A [][]int) [][]int {
	row_len := len(A[0])
	res := make([][]int, row_len)
	for i, _ := range res {
		res[i] = make([]int, len(A))
	}

	j := 0
	for _, val := range A {
		for i := 0; i < row_len; i++ {
			res[i][j] = val[i]
		}
		j++
	}

	return res
}

/**
two
耗时：4ms
内存：6.2M
 */
func transpose(A [][]int) [][]int {
	row_len := len(A[0])
	res := make([][]int, row_len)
	for i, _ := range res {
		res[i] = make([]int, len(A))
	}

	for row_key, val := range A {
		for clo_key, _ := range val {
			res[clo_key][row_key] = A[row_key][clo_key]
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)

//测试代码
func main()  {
	test := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	test_two := [][]int{{1,2,3},{4,5,6}}
	fmt.Println(test)
	res := transpose(test)
	fmt.Println(res)
	fmt.Println(test_two)
	res_two := transpose(test_two)
	fmt.Println(res_two)
}

/**
转置矩阵

one:
遍历切片A，获取key和val，然后在遍历val，将行每个元素变为列元素

复杂度分析：
	时间复杂度：O(C * R)
	空间复杂度：O(C * R)

two:
和第一种解法类似，但是不用刻意去计算列的长度，分别把行和列遍历，然后直接进行行列转换

复杂度分析：
	时间复杂度：O(C * R)
	空间复杂度：O(C * R)

遇到的问题:
	切片多维数组，刚开始卡在多维切片赋值，老报错，还没有报错信息，只有一串类似地址符的东西，
	最后发现是多维切片需要分步去申请内存空间的
 */

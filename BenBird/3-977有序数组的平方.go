//给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
//
//
//
// 示例 1：
//
// 输入：[-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//
//
// 示例 2：
//
// 输入：[-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 10000
// -10000 <= A[i] <= 10000
// A 已按非递减顺序排序。
//
// Related Topics 数组 双指针
// 👍 121 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(A []int) []int {
	a_len := len(A)
	j := a_len - 1

	if a_len == 0 {
		return []int{}
	}

	new_arr := make([]int, a_len, a_len)

	for i := 0; i <= j; {
		left_num := A[i] * A[i]
		right_num := A[j] * A[j]
		a_len--

		if left_num >= right_num {
			new_arr[a_len] = left_num
			i++
		} else {
			new_arr[a_len] = right_num
			j--
		}
	}

	return new_arr
}
//leetcode submit region end(Prohibit modification and deletion)

//测试
func main() {
	A := []int{-7,-3,2,3,11}
	new_arr := sortedSquares(A)
	fmt.Println(new_arr)
}

/**
有序数组的平方

耗时：36ms
内存：6.8M

双指针法

初始化一个新切片new_arr := make([]int, len(A), len(A))
使用左(i)右(j)指针，刚开始，左指针指向切片最左端，有指针指向切片最右端同时遍历该有序数组
直接将左右两端数值求平方进行对比，
若左边数值大，将左边求平方后数据存入new_arr，则左边指针向右移动，
若右边数值大，将左边求平方后数据存入new_arr，则右边指针向左移动，
继续左右指针指向的数值，求平方后对比，循环如此，直到左右指针相遇
最后切片new_arr就是所要获取的结果

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
 */

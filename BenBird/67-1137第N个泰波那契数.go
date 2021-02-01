//泰波那契序列 Tn 定义如下：
//
// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
//
// 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
//
//
//
// 示例 1：
//
// 输入：n = 4
//输出：4
//解释：
//T_3 = 0 + 1 + 1 = 2
//T_4 = 1 + 1 + 2 = 4
//
//
// 示例 2：
//
// 输入：n = 25
//输出：1389537
//
//
//
//
// 提示：
//
//
// 0 <= n <= 37
// 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
//
// Related Topics 递归
// 👍 56 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func tribonacciOld(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	x, y, z := 0, 1, 1
	for i := 3; i <= n; i++ {
		x, y, z = y, z, x + y + z
	}

	return z
}

var nums map[int]int
func tribonacciO(n int) int {
	nums = make(map[int]int)
	nums[0], nums[1], nums[2] = 0, 1, 1
	return dfsTribonacci(n)
}

func dfsTribonacci(n int) int {
	if _, ok := nums[n]; ok {
		return nums[n]
	}

	nums[n] = dfsTribonacci(n - 1) + dfsTribonacci(n - 2) + dfsTribonacci(n - 3)

	return nums[n]
}

func tribonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	nums := make([]int, n + 1)
	nums[0], nums[1], nums[2] = 0, 1, 1

	for i := 3; i <= n; i++ {
		nums[i] = nums[i - 1] + nums[i - 2] + nums[i - 3]
	}

	return nums[n]
}
//leetcode submit region end(Prohibit modification and deletion)

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
/**
空间优化：动态计算
n 为 0, 1, 2时，可以直接获取到值 0, 1, 1
n > 2时，先初始化前三个泰波那契数x, y, z = 0, 1, 1，
	然后进行n - 2次循环，循环内容为：x, y, z = y, z, x + y + z
循环结束后，z就为所有求的第n个泰波那契数

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(1)
 */

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
/**
记忆递归：
使用map用于保存泰波那契数, 初始化前三个泰波那契数n[0], n[1], n[2] = 0, 1, 1
使用递归：
	调用递归时，先判断map中是否已存在第n个泰波那契数
	若存在直接返回
	若不存在则继续进行递归调用计算第 n-1, n-2, n-3个泰波那契数
其实第n-1个泰波那契数计算出后, 第n-2和n-3个泰波那契数计算就是O(1)了

刚开始想到用递归，但是没有进行存储，导致计算第34个时，就超时了

复杂度分析：
	时间复杂度：O(n) 递归了n-2次
	空间复杂度：O(n)
 */

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
/**
动态计算并存储
这个相当于把前面两种做法中和了一下，感觉没啥意义
取第一种的动态计算的做法，使用递归调用的记忆，这里使用记忆没意义😂，与第一种相比反而额外使用了多余的空间
原意是想用这种思想，初始化前38个泰波那契数，然后获取第n个泰波那契数的时间和空间复杂度都是 O(1)但是，没这么做

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)

犯错：
刚开始，没有进行n 为0，1，1时的判断，认为只要对前三个泰波那契数进行初始化就行了，但是初始化nums时申请的空间是根据n的值来决定的
所以当n = 0或1时，就申请了1或2个空间，所以在初始化前三个泰波那契数时就会报错，因为nums就申请<= 2个的空间，但是需要初始化三个值，导致溢出报错
 */
//leetcode submit region end(Prohibit modification and deletion)

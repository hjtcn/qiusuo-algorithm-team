//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划
// 👍 1373 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairsOld(n int) int {
	dp := map[int]int{}

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}

	if val, ok := dp[n]; ok {
		return val
	}

	return 1
}

func climbStairs(n int) int {
	p, q, r := 0, 0, 1

	for i := 1; i <= n; i++ {
		p = q;
		q = r;
		r = p + q;
	}

	return r;
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := climbStairs(44)
	fmt.Println(res)
}

/**
题解：爬楼梯

复杂度分析：
	方法1：动态规划(申请额外空间)
	时间复杂度：O(n)
	空间复杂度：O(n)

	方法2：动态规划
	时间复杂度：O(n)
	空间复杂度：O(1)
 */

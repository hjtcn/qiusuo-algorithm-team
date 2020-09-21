//给你一个由 不同 整数组成的整数数组 arr 和一个整数 k 。
//
// 每回合游戏都在数组的前两个元素（即 arr[0] 和 arr[1] ）之间进行。比较 arr[0] 与 arr[1] 的大小，较大的整数将会取得这一回合的
//胜利并保留在位置 0 ，较小的整数移至数组的末尾。当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家 。
//
// 返回赢得比赛的整数。
//
// 题目数据 保证 游戏存在赢家。
//
//
//
// 示例 1：
//
// 输入：arr = [2,1,3,5,4,6,7], k = 2
//输出：5
//解释：一起看一下本场游戏每回合的情况：
//
//因此将进行 4 回合比赛，其中 5 是赢家，因为它连胜 2 回合。
//
//
// 示例 2：
//
// 输入：arr = [3,2,1], k = 10
//输出：3
//解释：3 将会在前 10 个回合中连续获胜。
//
//
// 示例 3：
//
// 输入：arr = [1,9,8,2,3,7,6,4,5], k = 7
//输出：9
//
//
// 示例 4：
//
// 输入：arr = [1,11,22,33,44,55,66,77,88,99], k = 1000000000
//输出：99
//
//
//
//
// 提示：
//
//
// 2 <= arr.length <= 10^5
// 1 <= arr[i] <= 10^6
// arr 所含的整数 各不相同 。
// 1 <= k <= 10^9
//
// Related Topics 数组
// 👍 13 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func getWinner(arr []int, k int) int {
	max_num := arr[0]
	times := 0

	for i := 1; i < len(arr); i++ {
		if (arr[i] > max_num) && times < k {
			max_num = arr[i]
			times = 1
			continue
		}

		times++

		if times == k {
			return max_num
		}
	}

	return max_num
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {
	arr := []int{1,25,35,42,68,70}
	k := 1
	res := getWinner(arr, k)
	fmt.Println(res)
}

/**
题解：

找出数组游戏的赢家

耗时：116ms
内存：9.3M

这道题把题分析清楚，就很容易做了，首先把最大数存入变量max_num，后面在遍历arr时，只需更新最大值的变量即可，无需移动元素位置
遍历过程中，如果max_num大于当前值
	则获胜次数times 自增1，并判断获胜次数是否等于k
		若等于k，则直接返回max_num
若max_num小于当前值
	max_num 更改为当前值，获胜次置为1

若循环结束(即times < k)，还未有返回值，则获胜的元素为max_num

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(1)
 */

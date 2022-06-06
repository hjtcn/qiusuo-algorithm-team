//你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有
//版本都是错的。
//
// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误
//的版本。你应该尽量减少对调用 API 的次数。
//
//
// 示例 1：
//
//
//输入：n = 5, bad = 4
//输出：4
//解释：
//调用 isBadVersion(3) -> false
//调用 isBadVersion(5) -> true
//调用 isBadVersion(4) -> true
//所以，4 是第一个错误的版本。
//
//
// 示例 2：
//
//
//输入：n = 1, bad = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= bad <= n <= 231 - 1
//
// Related Topics 二分查找 交互
// 👍 717 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var bad int = 2342234234

func isBadVersion(n int) bool {
	return n >= bad
}

func firstBadVersion(n int) int {
	low, high, mid := 1, n, 0

	for low <= high {
		mid = low + (high-low)/2

		if isBadVersion(mid) {
			high = mid - 1
			if !isBadVersion(high) {
				return mid
			}
		} else {
			low = mid + 1
			if isBadVersion(low) {
				return low
			}
		}
	}

	return 0
}

func main() {
	n := 3234234234234
	res := firstBadVersion(n)
	fmt.Println(res)
}

/**
总结：
思路：本题可抽象出一个从1递增的序列，从序列中找出第一个是错误的数
第一个错误的数：当前的数是错误的，前一个数是正确的

注意，该题的二分，开始和结尾，是1，n，做多了数组中的二分，开始和结尾容易整为 0， n-1
二分法，因为n 最大可为

复杂度分析：
	时间复杂度：O(logN) 因为n最大是2的31次方，所以做多计算约31次
	空间复杂度：O(1) 没有开辟额外的空间
 */


//leetcode submit region end(Prohibit modification and deletion)

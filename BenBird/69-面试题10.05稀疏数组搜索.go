//稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
//
// 示例1:
//
//  输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""],
// s = "ta"
// 输出：-1
// 说明: 不存在返回-1。
//
//
// 示例2:
//
//  输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""],
//s = "ball"
// 输出：4
//
//
// 提示:
//
//
// words的长度在[1, 1000000]之间
//
// Related Topics 二分查找
// 👍 40 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

//非递归
func findString(words []string, s string) int {
	low, high := 0, len(words) - 1

	for low <= high {
		for low < high && words[low] == "" {
			low++
			continue
		}

		for high > low && words[high] == "" {
			high--
			continue
		}

		if words[low] > s || words[high] < s {
			return -1
		}

		mid := low + ((high - low) >> 1)
		for mid > low && words[mid] == "" {
			mid--
			continue
		}

		if words[mid] == s {
			return mid
		} else if words[mid] < s {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
/**
本题，不适合直接套用二分模板, 但可以在二分模板基础上进行优化，如改有序数组存在空字符串, 所以需要想办法将空字符串给处理一下, 先从左右两边过滤
空字符串, 缩小左右边界, 再判断左右要查询的值是否在左右边界内，如果不在，直接退出
若在, 计算mid值, 如果words[mid] 为空，更新mid值, mid--, 向左偏移(其实这一步, 自我感觉向左向右都可), words[mid]不为空的话，直接使用mid
然后将mid位置的值和 s进行比较, words[mid] = s直接返回; words[mid] < s 更新low: low = mid + 1; words[mid] > s更新high: high = mid - 1
直到找到s的值，否则返回 -1

复杂度分析：
	时间复杂度：O(logN)
	空间复杂度：O(1)
 */

//递归
func findStringNew(words []string, s string) int {
	return findRecursion(words, s, 0, len(words) - 1)
}

func findRecursion(words []string, s string, low int, high int) int {
	for low < high && words[low] == "" {
		low++
		continue
	}

	for high > low && words[high] == "" {
		high--
		continue
	}

	if words[low] > s || words[high] < s {
		return -1
	}

	mid := low + ((high - low) >> 2)
	for mid > low && words[mid] == "" {
		mid--
		continue
	}

	if words[mid] == s {
		return mid
	} else if words[mid] < s {
		return findRecursion(words, s, mid + 1, high)
	} else {
		return findRecursion(words, s, low, mid - 1)
	}
}

/**
递归：思路和非递归思路类似，先处理空字符串，然后判断做递归调用，更新high和low的值
复杂度分析：
	时间复杂度：O(logN)
	空间复杂度：O(logN)

 */

func main()  {
	str := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	//fmt.Println(findString(str, "ball"))
	fmt.Println(findStringNew(str, "ball"))
}
//leetcode submit region end(Prohibit modification and deletion)

/**
二分查找：是一个有序的数据集合，每次都通过跟区间的中间元素对比，将待查找的区间缩小为之前的一半，直到找到要查找的元素，或者区间被缩小为 0。

二分模板：
func dichotomy(words []string, s string) int {
	low, high := 0, len(words) - 1
	mid := low + ((high - low) >> 1)

	for low <= high {
		if words[mid] == s {
			return mid
		} else if words[mid] < s {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

重点：
循环退出条件：low <= high, 不是 low < high, 当low == high时，还需要进行比较一下
mid的取值：mid = (low + high) / 2
	low 和 high 如果太大，之和可能会溢出(如：low 和 high 都为 int, 如果low 和 high的值都无限接近int 的最大值, 之和会超过int最大值的, 导致溢出),
	所以可改为，mid = low + (high - low) / 2, 而除以2, 可以使用位运算, 因为据说相比除法，计算机处理位运算快的多, 所以最终可以把mid值的
	更新优化为：mid = low + ((high - low) >> 2), 因为运算符 ">>" 的优先级低于 "+" 所以注意加括号
low 和 high 的更新：如果words[mid] > s, 更新high的值：high = mid - 1, 如果words[mid] < s, 更新low的值：low = mid + 1
 */

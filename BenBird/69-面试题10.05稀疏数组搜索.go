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

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findString(words []string, s string) int {
	low, height := 0, len(words) - 1

	for low <= height {
		for low < height && words[low] == "" {
			low++
			continue
		}

		for height > low && words[height] == "" {
			height--
			continue
		}

		if words[low] > s || words[height] < s {
			return -1
		}

		mid := low + ((height - low) >> 1)
		for mid > low && words[mid] == "" {
			mid--
			continue
		}

		if words[mid] == s {
			return mid
		} else if words[mid] < s {
			low = mid + 1
		} else {
			height = mid - 1
		}
	}

	return -1
}

func main()  {
	str := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	fmt.Println(findString(str, "ball"))
}
//leetcode submit region end(Prohibit modification and deletion)

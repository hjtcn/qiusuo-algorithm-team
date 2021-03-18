//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
//
//
// 示例 2:
//
// 输入: "race a car"
//输出: false
//
// Related Topics 双指针 字符串
// 👍 348 👎 0
package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	str_len := len(s)

	if str_len <= 1 {
		return true
	}

	left, right := 0, str_len - 1

	for left < right {
		for left < right && !isAlNum(s[left]) {
			left++
		}

		for left < right && !isAlNum(s[right]) {
			right--
		}

		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}

	return true
}

func isAlNum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

/**
双指针
双指针，设置left和right指针，分别指向字符串的首和尾，并比较两指针指向字符串的值
若相等left指针则想右移，即left++,right指针向左移动，即right--
若不等直接返回false，代表改字符串不是回文串
然后依次移动left和right指针，然后进行对比
在left和right指针移动过程中，遇到非字母或数字的其他字符，直接跳过，接着移动，直到移动到字母或数字字符。

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(1)
*/

func main() {
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)

	fmt.Println(res)
}
//leetcode submit region end(Prohibit modification and deletion)

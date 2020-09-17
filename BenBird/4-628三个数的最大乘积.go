//给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
//
// 示例 1:
//
//
//输入: [1,2,3]
//输出: 6
//
//
// 示例 2:
//
//
//输入: [1,2,3,4]
//输出: 24
//
//
// 注意:
//
//
// 给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
// 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
//
// Related Topics 数组 数学
// 👍 171 👎 0
package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func maximumProduct(nums []int) int {
	min_one, min_two := 1000, 1000
	max_one, max_two, max_three := -1000, -1000, -1000

	for _, val := range nums {
		if val < min_one {
			min_two = min_one
			min_one = val
		} else if val < min_two {
			min_two = val
		}

		if val > max_one {
			max_three = max_two
			max_two = max_one
			max_one = val
		} else if val > max_two {
			max_three = max_two
			max_two = val
		} else if val > max_three {
			max_three = val
		}
	}

	left_max := min_one * min_two * max_one
	right_max := max_one * max_two * max_three

	if left_max > right_max {
		return left_max
	}

	return right_max
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{1,2,3}
	max := maximumProduct(nums)
	fmt.Println(max)
}

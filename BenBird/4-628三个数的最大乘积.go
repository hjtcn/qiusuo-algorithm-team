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
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func maximumProduct_one(nums []int) int {
	sort.Ints(nums)
	nums_len := len(nums)
	left_one := nums[0]
	left_two := nums[1]
	right_last := nums[nums_len-1]
	right_two := nums[nums_len-2]
	right_one := nums[nums_len-3]

	left_max := right_last * left_one * left_two
	right_max := right_last * right_one * right_two

	if left_max > right_max {
		return left_max
	} else {
		return right_max
	}
}

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

/**
题解

最大三数乘积

one:
耗时：76ms
内存：6.6M

先对数组进行正序排列，然后根据规律可得，最大乘积之和的三个数必有排在最后一位数，另外两个数则存在第一、二位数 或 倒数第二、三位,
然后分别对这两种可能的三个数，进行乘积比较，然后获取最大三数乘积之和

复杂度分析：
	时间复杂度：O(NlogN)
	空间复杂度：O(logN)

two：
耗时：40ms
内存：6.6M

线性扫描，找出最小的两位数，和最大的三位数，然后进行比较，同上排序法，将最小的两位数和最大的数的乘积 与 最大的三位数乘积比较
获取最大三数乘积

复杂度分析：
	时间复杂度：O(N)
	空间复杂度：O(1)
 */

//数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。 
//
// 示例 1： 
//
// 输入：[1,2,5,9,5,9,5,5,5]
//输出：5 
//
// 
//
// 示例 2： 
//
// 输入：[3,2]
//输出：-1 
//
// 
//
// 示例 3： 
//
// 输入：[2,2,1,1,1,2,2]
//输出：2 
//
// 
//
// 说明： 
//你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？ 
// Related Topics 位运算 数组 分治算法 
// 👍 27 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement_one(nums []int) int {
	m := map[int]int{}
	max_len := len(nums) / 2

	for _, val := range nums {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}

		if m[val] > max_len {
			return val
		}
	}

	return -1
}

func majorityElement(nums []int) int {
	major := -1
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num;
			count++
		} else {
			if major == num {
				count++
			} else {
				count--
			}
		}
	}

	counter := 0
	if count <= 0 {
		return -1
	} else {
		for _, num := range nums {
			if num == major {
				counter++
			}
		}
	}

	if counter > (len(nums) / 2) {
		return major
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{1,2,1,3,1,4,1,5,1}
	a := majorityElement(nums)
	fmt.Println(a)
}

/**
one
耗时: 44ms
内存: 6.1M
开辟一个新的内存空间，空间大小 最坏情况 为 nums的长度，也就是 n, 应应与不符合题目 时间复杂度O(n) 空间复杂度O(1)的要求
遍历nums, 用map存放nums中的元素和其出现的次数, 当出现次数大于nums长度的一半，则返回该元素, 否则返回 -1

复杂度分析
	时间复杂度：O(n)
	空间复杂度：O(n)


two 摩尔算法
耗时: 16ms
内存: 6.1M
主要使用相邻两个数进行两两抵消的方法，因为所求数值必出现次数大于nums总元素个数的一半，所以：
假设一定存在这个数，则一定存在相邻两个数相同 或 这种不相邻但首尾呼应{1,2,1,3,1,4,1,5,1}，这种前面的数都两两抵消，只剩下1，然而验证的时候
刚好符合
假设不存在这个数，则在进行两两抵消后，有两种结果，1. 全部被抵消 2. 剩下一个数字 然后经验证不成立

就很棒

复杂度分析
	时间复杂度：O(n)
	空间复杂度：O(1)
*/
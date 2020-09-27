//给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k 个元素的顺序。我们要执行零次或多次煎饼翻转（
//按顺序一次接一次地进行）以完成对数组 A 的排序。
//
// 返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。
//
//
//
// 示例 1：
//
// 输入：[3,2,4,1]
//输出：[4,2,4,3]
//解释：
//我们执行 4 次煎饼翻转，k 值分别为 4，2，4，和 3。
//初始状态 A = [3, 2, 4, 1]
//第一次翻转后 (k=4): A = [1, 4, 2, 3]
//第二次翻转后 (k=2): A = [4, 1, 2, 3]
//第三次翻转后 (k=4): A = [3, 2, 1, 4]
//第四次翻转后 (k=3): A = [1, 2, 3, 4]，此时已完成排序。
//
//
// 示例 2：
//
// 输入：[1,2,3]
//输出：[]
//解释：
//输入已经排序，因此不需要翻转任何内容。
//请注意，其他可能的答案，如[3，3]，也将被接受。
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 100
// A[i] 是 [1, 2, ..., A.length] 的排列
//
// Related Topics 排序 数组
// 👍 73 👎 0
package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func pancakeSort(arr []int) []int {
	res := []int{}
	recursiveSort(arr, len(arr), &res)
	return res
}

/**
递归排序
*/
func recursiveSort(arr []int, n int, res *[]int) {
	if n == 1 {
		return
	}

	k := getMaxIndex(arr[:n])
	reverse(arr, 0, k)
	*res = append(*res, k + 1)
	reverse(arr, 0, n - 1)
	*res = append(*res, n)

	recursiveSort(arr, n - 1, res)
}

/**
获取切片中最大元素的下标
*/
func getMaxIndex(nums []int) int {
	res := 0
	index := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
			index = i
		}
	}

	return index
}

/**
旋转切片中指定范围内的元素
*/
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {
	arr := []int{3,2,4,1}
	res := pancakeSort(arr)

	fmt.Println(res)
}

/**
题解：煎饼排序

耗时：0ms
内存：2.5M

首先找到需要排序元素的最大值，并获取其下标
然后将首位元素到最大值中间的所有元素进行反转一次，将最大值置于第一位
需要排序元素个数为n，将前n个元素进行反转一次，此时最大值就在它应该在位置上了

按照此思路依次循环往复，直到需要排序元素个数为1，即排序完成

复杂度分析：
	时间复杂度：O(n^2)
	空间复杂度：O(n)

 */
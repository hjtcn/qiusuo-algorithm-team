//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 
//
// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2
//, 1, 1, 0, 0]。 
//
// 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。 
// Related Topics 栈 哈希表 
// 👍 544 👎 0
package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(T []int) []int {
	temp_stack := []int{}
	res := make([]int, len(T))
	stack_top := 0
	for key := 0; key < len(T); key++ {
		for len(temp_stack) > 0 && T[key] > T[temp_stack[len(temp_stack) - 1]] {
			stack_top = temp_stack[len(temp_stack) - 1]
			res[stack_top] = key - stack_top
			temp_stack = temp_stack[:len(temp_stack) - 1]
		}

		temp_stack = append(temp_stack, key)
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	arr := []int{73,74,75,71,69,72,76,73}
	res := dailyTemperatures(arr)

	fmt.Println(res)
}

/**
题解：每日温度

刚开始题就理解错了，折腾了半天做出了一个错解，最后看懂题又做不出来了
看题解，压根看不懂，昨天就放弃了
今天清醒了一下，静下心去看题解+视频讲解，才整明白

单调栈：先初始化一个空数组res，和一个空栈temp_stack，temp_stack存储当前气温的下标
遍历气温列表T，
	若temp_stack栈为空或栈顶元素stack_top所对应的气温值 > 当前遍历气温值，则将当前气温下标入栈
	若temp_stack非空 并且 栈顶元素stack_top对应的气温 < 当前气温，则将栈顶元素stack_top出栈，并将气温信息存入数组res
	即：res[stack_top] = key - stack_top


复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)

 */

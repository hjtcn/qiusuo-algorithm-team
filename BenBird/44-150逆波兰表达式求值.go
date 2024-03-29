//根据 逆波兰表示法，求表达式的值。
//
// 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//
//
// 说明：
//
//
// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//
//
//
//
// 示例 1：
//
// 输入: ["2", "1", "+", "3", "*"]
//输出: 9
//解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
//
//
// 示例 2：
//
// 输入: ["4", "13", "5", "/", "+"]
//输出: 6
//解释: 该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
//
//
// 示例 3：
//
// 输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
//输出: 22
//解释:
//该算式转化为常见的中缀算术表达式为：
//  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22
//
//
//
// 逆波兰表达式：
//
// 逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
//
//
// 平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
// 该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
//
//
// 逆波兰表达式主要有以下两个优点：
//
//
// 去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
// 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。
//
// Related Topics 栈
// 👍 201 👎 0
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
//没有多余的判断，耗时较少，单不易维护
func evalRPN(tokens []string) int {
	temp_stack := []int{}

	for _, val := range tokens {
		length := len(temp_stack)
		switch val {
		case "+":
			temp_stack = append(temp_stack[:length - 2], temp_stack[length - 2] + temp_stack[length - 1])
		case "-":
			temp_stack = append(temp_stack[:length - 2], temp_stack[length - 2] - temp_stack[length - 1])
		case "*":
			temp_stack = append(temp_stack[:length - 2], temp_stack[length - 2] * temp_stack[length - 1])
		case "/":
			temp_stack = append(temp_stack[:length - 2], temp_stack[length - 2] / temp_stack[length - 1])
		default:
			num, _ := strconv.Atoi(val)
			temp_stack = append(temp_stack, num)
		}
	}

	return temp_stack[0]
}

//易维护，单运算耗时稍稍增加
func evalRPNNew(tokens []string) int {
	temp_stack := []int{}

	for _, val := range tokens {
		if len(val) > 0 && strings.ContainsAny(val, "+-*/") {
			length := len(temp_stack)
			elem_one := temp_stack[length - 1]
			elem_two := temp_stack[length - 2]
			temp_stack = temp_stack[:len(temp_stack) - 2]

			switch val {
			case "+":
				temp_stack = append(temp_stack, elem_two + elem_one)
			case "-":
				temp_stack = append(temp_stack, elem_two - elem_one)
			case "*":
				temp_stack = append(temp_stack, elem_two * elem_one)
			case "/":
				temp_stack = append(temp_stack, elem_two / elem_one)
			}
		} else {
			num, _ := strconv.Atoi(val)
			temp_stack = append(temp_stack, num)
		}
	}

	return temp_stack[0]
}
//leetcode submit region end(Prohibit modification and deletion)

/**
测试代码
 */
func main() {
	str := []string{"4","13","5","/","+"}
	res := evalRPN(str)

	fmt.Println(res)
}

/**
题解：逆波兰表达式

初始化一个空切片作为栈temp_stack使用, 遍历给定切片tokens,
	遍历的当前值若是数字, 则入栈
	若是运算符, 则将temp_stack连续两次出栈, 第一次出栈的数据在运算符右边，第二次出栈的数据在运输符左边，然后进行运算
	运算结果入栈temp_stack
当tokens遍历完成后, 返回栈temp_stack第一个元素, 也就是整个tokens的运算结果

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
 */

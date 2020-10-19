//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
// 👍 1925 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	brackets_map := map[string]string{
		")" : "(",
		"]" : "[",
		"}" : "{",
	}

	temp_stack := []string{}
	for _, str := range s {
		temp_str := string(str)
		stack_len := len(temp_stack)

		if (stack_len > 0) && (brackets_map[temp_str] == temp_stack[stack_len - 1]) {
			temp_stack = temp_stack[:stack_len - 1]
			continue
		}

		temp_stack = append(temp_stack, temp_str)
	}

	if len(temp_stack) == 0{
		return true
	}

	return false
}
//leetcode submit region end(Prohibit modification and deletion)

func main()  {
	str := "()}"
	res := isValid(str)

	fmt.Println(res)
}

/**
题解：有效的括号

耗时：0ms
内存：2.2M

利用栈：将给定字符串s从左到右依次入栈temp_stack,
入栈前先判断将要入栈元素与栈顶元素是否匹配
	若不匹配：则将元素入栈
	若匹配：则该元素不入栈，且弹出栈顶元素
当字符串s遍历完成后，判断栈temp_stack 是否为空
	若为空：则字符串s为有效括号
	若不为空：则字符串s为无效括号

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
 */


//给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。 
//
// 注意：如果对空文本输入退格字符，文本继续为空。 
//
// 
//
// 示例 1： 
//
// 输入：S = "ab#c", T = "ad#c"
//输出：true
//解释：S 和 T 都会变成 “ac”。
// 
//
// 示例 2： 
//
// 输入：S = "ab##", T = "c#d#"
//输出：true
//解释：S 和 T 都会变成 “”。
// 
//
// 示例 3： 
//
// 输入：S = "a##c", T = "#a#c"
//输出：true
//解释：S 和 T 都会变成 “c”。
// 
//
// 示例 4： 
//
// 输入：S = "a#c", T = "b"
//输出：false
//解释：S 会变成 “c”，但 T 仍然是 “b”。 
//
// 
//
// 提示： 
//
// 
// 1 <= S.length <= 200 
// 1 <= T.length <= 200 
// S 和 T 只含有小写字母以及字符 '#'。 
// 
//
// 
//
// 进阶： 
//
// 
// 你可以用 O(N) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？ 
// 
//
// 
// Related Topics 栈 双指针 
// 👍 232 👎 0
package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompareOld(S string, T string) bool {
	s_stack := []string{}
	t_stack := []string{}

	for _, val := range S {
		s_str := string(val)
		if s_str == "#" {
			if len(s_stack) > 0 {
				s_stack = s_stack[:len(s_stack) - 1]
			}
		} else {
			s_stack = append(s_stack, s_str)
		}
	}

	for _, val := range T {
		t_str := string(val)
		if t_str == "#" {
			if len(t_stack) > 0 {
				t_stack = t_stack[:len(t_stack) - 1]
			}
		} else {
			t_stack = append(t_stack, t_str)
		}
	}

	if len(s_stack) != len(t_stack) {
		return false
	}

	for i := 0; i < len(s_stack); i++ {
		s_str := s_stack[i]
		t_str := t_stack[i]

		if s_str != t_str {
			return false
		}
	}

	return true
}

//第一种解法优化 start
func backspaceComparePrefect(S string, T string) bool {
	return build(S) == build(T)
}

func build(str string) string {
	s := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s) - 1]
		}
	}

	return string(s)
}
//第一种解法优化 end

/**
都是使用辅助栈，遍历字符串依次
若遇到非#字符串，直接将该字符入栈
若遇到#，先判断栈中是否为空
	若为空，继续遍历
	若非空，弹出栈顶元素，继续遍历
最后把简化后的字符串进行对比，相同则 返回true，不同则false

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)

上面两种解法都是用同一种思想，但代码的可读性、可维护性和代码量却差别很大
第一种相较于第二种：
1. 没有把重复代码封装抽出来，导致代码臃肿
2. 没有合理使用字符类型，使字符串对比较为复杂
 */


func backspaceCompare(S string, T string) bool {
	s_pointer, t_pointer := len(S) - 1, len(T) - 1
	s_skip, t_skip := 0, 0

	for s_pointer >= 0 || t_pointer >= 0 {
		s_pointer, s_skip = traver(S, s_pointer, s_skip)
		t_pointer, t_skip = traver(T, t_pointer, t_skip)

		if s_pointer >= 0 && t_pointer >= 0 {
			if S[s_pointer] != T[t_pointer] {
				return false
			}
		} else if s_pointer >= 0 || t_pointer >= 0 {
			return false
		}
		s_pointer--
		t_pointer--
	}

	return true
}

func traver(str string, pointer int, skip int) (int, int) {
	for pointer >= 0 {
		if str[pointer] == '#' {
			pointer--
			skip++
		} else if skip > 0 {
			pointer--
			skip--
		} else {
			break
		}
	}

	return pointer, skip
}
//leetcode submit region end(Prohibit modification and deletion)

/**
双指针：

定义两个指针s_pointer, t_pointer指向字符串尾部，在定义两个变量s_skip, t_skip记录当前遇到且未被使用的#个数
两指针由字符串尾部向前移动
当遇到 # 时指针向前移动1位，# 个数加 1,
当遇到非 # 且 # 的个数大于 0, 则指针向前移动1位, # 个数减1
其他情况，也就是当遇到非 # 且 # 个数等于 0, 记录当前遇到的字符, 用来进行比较,
	若两字符串此时都未遍历完，且两字符不同，则直接返回false
	若两字符串此时只有一个字符串遍历完成，也返回false
否则，指向字符串的两指针继续向前遍历，若两字符串都遍历完成时, 则返回true


复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
 */

func main() {
	s_str := "bbbextm"
	t_str := "bbb#extm"
	res := backspaceCompare(s_str, t_str)
	fmt.Println(res)
}

/**
总结：
	缺少抽象意识、数据类型掌握较差
 */


/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

1. Clearfication:

找重复性
看了一会发现并没有思路就去看题解了：
题解里面方法一：
栈操作，将字母、数字和括号看成是独立的TOKEN,并用栈来维护这些TOKEN。
具体做法，遍历这个栈：
如果当前的字符为数位，解析出一个数字（连续的多个数位）并进栈
如果当前的字符为字母或者左括号，直接进栈
如果当前的字符为右括号，开始出栈，一直到左括号出栈，出栈序列反转后拼接成一个字符串，此时取出栈顶的数字，就是这个字符串应该出现的次数
func decodeString(s string)string {
    stk := []string{}
    ptr := 0
    
    for ptr < len(s) {
        cur := s[ptr]
        
        if cur >= '0' && cur <= '9' {
            digits := getDigits(s, &ptr)
            stk = append(stk, digits)
        }else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
            stk = append(stk, string(cur))
            ptr++
        }else {
            ptr++
            sub := []string{}
            for stk[len(stk) - 1] != "[" {
                sub = append(sub, stk[len(stk) - 1])
                stk = stk[:len(stk) - 1]
            }
            
            for i := 0;i < len(sub) / 2;i++ {
                sub[i],sub[len(sub) - i - 1] = sub[len(sub)-i-1],sub[i]
            }
            // 左括号出栈
            stk = stk[:len(stk) - 1]
            repTime,_ := strconv.Atoi(stk[len(stk) - 1])
            // 数字出栈
            stk = stk[:len(stk)-1]
            t := strings.Repeat(getString(sub),repTime)
            stk = append(stk,t)
        }
    }
    
    return getString(stk)
}

func getDigits(s string,ptr *int)string {
    ret := ""
    
    for ;s[*ptr] >= '0' && s[*ptr] <= '9';*ptr++ {
        ret += string(s[*ptr])
    }
    
    return ret
}
func getString(v []string) string {
    ret := ""
    for _,s := range v {
        ret += s
    }
    return ret
}


递归

var (
    src string
    ptr int
)
func decodeString(s string) string {
    src = s
    ptr = 0
    return getString
}
func getString() string {
    // terminator
    if ptr == len(src) || src[ptr] == ']' {
        return ""
    }
    cur := src[ptr]
    repTime := 1
    ret := ""
    if cur >= '0' && cur <= '9' {
        repTime = getDigits()
        ptr++
        str := getString()
        ptr++
        // process current logic
        ret = strings.Repeat(str, repTime)
    }else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
        ret = string(cur)
        ptr++
    }
    // drill down
    return ret + getString()
}
func getDigits() int {
    ret := 0
    for ; src[ptr] >= '0' && src[ptr] <= '9';ptr++ {
        ret = ret * 10 + int(src[ptr] - '0')
    }
    return ret
}

复杂度分析：
时间复杂度：O(n)
空间复杂读：O(n)

总结：
1. 题目中有很多细节要处理

2. 思维要完整

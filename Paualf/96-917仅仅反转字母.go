/**
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
示例 1：
输入："ab-cd"
输出："dc-ba"
示例 2：
输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：
输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"
 
提示：
S.length <= 100
33 <= S[i].ASCIIcode <= 122 
S 中不包含 \ or "
**/

1. Clearfication:
    不是字母的字符保留在原地
    ab-cd
    dc-ba
    a-bC-dEf-ghIj
    j-Ih-gfE-dCba
    双指针，进行交换，如果不是字符则跳过

2. Coding:
/*
    不是字母的字符保留在原地
    ab-cd
    dc-ba
    a-bC-dEf-ghIj
    j-Ih-gfE-dCba
    双指针，进行交换，如果不是字符则跳过
*/
func reverseOnlyLetters(S string) string {
    i,j := 0,len(S)-1
    for i < j {
        for !isChar(S[i]) {
            i++
        }
        for !isChar(S[j]) {
            j--
        }
        S[i],S[j] = S[j],S[i]
    }
    return S
}
func isChar(s byte) bool {
    if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
        return true
    } 
    return false
}
报错：Line 25: Char 19: cannot assign to S[i] (strings are immutable) (solution.go) Line 25: Char 19: cannot assign to S[j] (strings are immutable) (solution.go)
string 是不可以修改的，需要转换成byte 的

3. 看题解：
func reverseOnlyLetters(S string) string {
    buf := []byte(S)
    
    for i,j := 0,len(buf) - 1;i < j; {
        for i < j && !isLetter(buf[i]) {
            i++
        }
        
        for i < j && !isLetter(buf[j]) {
            j--
        }
        buf[i],buf[j] = buf[j],buf[i]
        i++
        j--
    }
    
    return string(buf)
}
func isLetter(c byte) bool {
    if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
        return true
    }
    
    return false
}
使用栈实现：
func reverseOnlyLetters(S string) string {
    stack := make([]rune,0)
    buf := []rune{}
    
    for _,v := range S {
        if isLetter(v) {
            stack = append(stack, v)
        }
    }
    
    for _,v := range S {
        if isLetter(v) {
            bytes := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            
            buf = append(buf, bytes)
        }else {
            buf = append(buf, v)   
        }
    }
    
    return string(buf)
}
func isLetter(c rune) bool {
    if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
        return true
    }
    
    return false
}

4. 复杂度分析：
时间复杂度：O(n) : 遍历字符串长度
空间复杂度：O(n)：使用byte对字符串进行存储

5. 总结：
5.1: String 不能进行交换
5.2. 对Go中string、rune、byte 的区别不知道 https://blog.golang.org/strings:这篇文章明天早上读一下
5.3: 第一时间想到双指针，解决方法上要依赖原有字符串的队列，可以借用栈实现，栈这个思路还是蛮好的

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
示例 1：
输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：
输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 
提示：
输入的字符串长度不会超过 1000 。

1. Clearfication:
计算这个字符串中有多少个回文子串
        abc
            a  b  c
        aaa
            a a a aa aa aaa
枚举字符串，然后判断是否是回文子串
            ret := 0
            for i : 0;i < len(s);i++ {
                for j := i + 1;j <= len(s);j++ {
                    if check(s[i:j]) {
                        ret++
                    }
                }
            }
            return ret
            func check(s string)bool {
                for i := 0,j := len(s) - 1;i < j;i++,j-- {
                    if s[i] != s[j] {
                        return false
                    }
                }
                return true
            }
            s[2:3] => c
            回文中心是两个位置？
 动态规划
            寻找子问题，然后判断回文子串的数量
            子问题是什么呢？

2. Coding:
func countSubstrings(s string) int {
    ret := 0
     for i := 0;i < len(s);i++ {
        for j := i + 1;j <= len(s);j++ {
             if check(s[i:j]) {
                    ret++
                }
            }
        }
    return ret
}

 func check(s string)bool {
    for i,j := 0,len(s) - 1;i < j;i,j = i+1,j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
 }

3. 看题解：
看了下动态规划，没有看到那个 dp[i][j] 的意思，也看了自己原来写的题解，感觉还是怪怪的

4. 复杂度分析：
时间复杂度：O(n*n*n): for 循环两层遍历+一层check
空间复杂度：O(1)

5. 总结：
5.1: 动态规划感觉还是它的表达和含义还是蛮重要的

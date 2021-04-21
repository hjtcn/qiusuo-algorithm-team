给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
致谢：
特别感谢 @pbrother 添加此问题并且创建所有测试用例。
示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false
 
提示：
0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。
1. Clearfication:
一开始想动态规划，感觉和最长公共子序列有点像，不过自己没有想出来
后来想真的需要动态规划么，我一个一个比不行么，发现也是可以的，一开始有疑问的地方是 a b c，会不会有其他意外情况，后来发现自己想多了，想找到a,然后再去找 b,c
 具体实现：
       比对s和t,s 是从 0开始，t是从0开始
          如果
            s[i] == t[i]
                i++

2. Coding:
func isSubsequence(s string, t string) bool {
    i := 0
    j := 0
    for ;i < len(s) && j < len(t); {
        if s[i] == t[j] {
            i++
            j++
        }
        j++
    }
    return i == len(s)
}

上面的代码提交的时候wrong 掉了，自己打印了一下 i 和 len(s) 比对了一下：发现 s[i] 和 t[j] 相等的时候 j 的游标会向后面走两次
修改后的代码
func isSubsequence(s string, t string) bool {
    i := 0
    j := 0
    for ;i < len(s) && j < len(t); {
        if s[i] == t[j] {
            i++
            j++
        }else {
           j++
        }
    }
    
    return i == len(s)
}

3. 看题解：
看了下题解，发现自己写的代码确实不是很好,并需要加 else 的
func isSubsequence(s string,t string) bool {
    m,n := len(s),len(t)
    i,j := 0,0
    
    for i < m && j < n {
        if s[i] == t[j] {
            i++
        }
        j++
    }
    
    return i == m
}

4. 复杂度分析：
时间复杂度：O(n + m)
空间复杂度：O(1)

5. 总结：
5.1: 真的要做很难的题才能锻炼自己的代码能力么，easy 题中的 if 和 else 逻辑有时候都没有分析清楚，所以不要好高骛远，脚踏实地去写代码，然后偶尔仰望下星空

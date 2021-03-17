/*
给定一个字符串 s，计算具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是连续的。
重复出现的子串要计算它们出现的次数。
示例 1 :
输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
请注意，一些重复出现的子串要计算它们出现的次数。
另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
示例 2 :
输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
提示：
s.length 在1到50,000之间。
s 只包含“0”或“1”字符。
*/

1. Clearfication:
    思维不要太发散：
        找出所有的子串
        然后判断它们是否具有相同数量的连续1和0
            怎么判断它们具有相同数量0和1的非空（连续）子字符串的数量呢？    

2. Coding:

3. 看题解:

func countBinarySubstrings(s string) int {
    counts := []int{}
    ptr,n := 0,len(s)
    
    for ptr < n {
        c := s[ptr]
        count := 0
        for ptr < n && s[ptr] == c {
            ptr++
            count++
        }
        counts = append(counts, count)
    }
    ans := 0
    for i := 1;i < len(counts);i++ {
        ans += min(counts[i], counts[i-1]
    }
    
    return ans
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

里面控制代码的小细节还是蛮厉害的， 一开始和自己比,ptr 肯定是 ++ 的
对于某一个位置i,我们其实只关心 i - 1 位置的counts值是多少，所以可以用一个last变量来维护

func countBinarySubstrings(s string) int {
    var ptr,last,ans int
    n := len(s)
    for ptr < n {
        c := s[ptr]
        count := 0
        for ptr < n && s[ptr] == c {
            ptr++
            count++
        }
        ans += min(count,last)
        last = count
    }
    
    return ans
}

func min(x,y int)int {
    if x < y {
        return x
    }
    return y
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：使用数组存储复杂度是O(n), 使用last变量空间复杂度是 O(1)

5. 总结：
5.1: 将字符串s按照0和1的连续段分组转化为相邻字符相等的数量进行解决  将问题转换为计算机可以解决的问题这种思维和方式还是要多锻炼的

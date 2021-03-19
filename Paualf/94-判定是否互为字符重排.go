给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
示例 1：
输入: s1 = "abc", s2 = "bca"
输出: true 
示例 2：
输入: s1 = "abc", s2 = "bad"
输出: false
说明：
0 <= len(s1) <= 100
0 <= len(s2) <= 100

1. Clearfication:
s1="abc",s2="bca"
    使用hash存储每个字符串出现的次数，hash 的结构怎么定义和存储呢？
    如果都是小写字符串是不是可以将它缩小空间到 0 ～ 26 中间
    题目中没有给出字符串的范围，所以map的大小我们还没有办法很好的确定，所以map的空间我们可以给它一个动态值

2. Coding:
func CheckPermutation(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    m := make(map[rune]int, 0)
    for _,v := range s1 {
        m[v] = m[v] + 1
    }
    for _,v := range s2 {
        m[v] = m[v] - 1
    }
    for _,v := range s1 {
        if m[v] > 0 {
            return false
        }
    }
    return true
}

3. 看题解：
使用 string.Spilt() 函数拆分成单个字母的字符串切片，然后再利用sort.Strings() 函数将字符串切片进行排序，最后利用strings.Join() 函数将排序后的字符串组合成字符串，最后判断两个字符是否相等
func CheckPermutation(s1 string,s2 string) bool {
    tmp1 := strings.Split(s1, "")
    tmp2 := strings.Split(s2, "")
    sort.Strings(tmp1)
    res1 := strings.Join(tmp1, "")
    sort.Strings(tmp2)
    res2 := strings.Join(tmp2, "")
    return res1 == res2
}

4. 复杂度分析：
时间复杂度：O(n) : 遍历整个字符串
空间复杂度：O(n) : 使用map存储字符出现次数

5. 总结：
5.1: 这道题目主要考查对hash的理解

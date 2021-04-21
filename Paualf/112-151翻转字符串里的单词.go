给定一个字符串，逐个翻转字符串中的每个单词。
说明：
无空格字符构成一个 单词 。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 
示例 1：
输入："the sky is blue"
输出："blue is sky the"
示例 2：
输入："  hello world!  "
输出："world! hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：
输入："a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
示例 4：
输入：s = "  Bob    Loves  Alice   "
输出："Alice Loves Bob"
示例 5：
输入：s = "Alice does not even like bob"
输出："bob like even not does Alice"
 
提示：
1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词
 
进阶：
请尝试使用 O(1) 额外空间复杂度的原地解法。
1. Clearfication:
    1. 逐个翻转字符串中的每个单词
    2. 输入字符串在前面或者后面包含多余的空格，反转后的字符不能包括
    3. 两个单词间有多余的空格，反转后单词的空格减少到一个
知道题目里面的关键以后，发现还是没有怎么写代码的思路。。。

2. Coding:

3. 看题解:
Fields: 按空格将字符串切分成字符串切片
字符串切片反转
Join: 将字符串切片以空格为分隔符合并成字符串
func reverseWords(s string) string {
    res := strings.Fields(s)
    length := len(res)
    
    for i := 0;i < length / 2;i++ {
        res[i],res[length-1-i] = res[length - 1 - i],res[i]
    }
    
    return strings.Join(res, " ")
}

手动实现字符串中的空格过滤，将返回的字符串然后使用空格分隔成切片数组，然后进行交换处理
func reverseWords(s string) string {
    if s == "" {
        return ""
    }
    
    s = PreProcess(s)
    wordarry := strings.Split(s, " ")
    left := 0
    right := len(wordarry) - 1
    
    for left < right {
        wordarry[left],wordarry[right] = wordarry[right],wordarry[left]
        left++
        right--
    }
    
    res := strings.Join(wordarry, " ")
    return res
}

func PreProcess(s string) string {
    l := len(s)
    var res []byte
    
    flag := 1
    for l != 0 && s[l - 1] == ' ' {
        l--
    }
    
    for i := 0;i < l;i++ {
        if s[i] != ' ' {
            res = append(res, s[i])
            flag = 0
        }
        
        if s[i] == ' ' && flag == 0 {
            res = append(res,s[i])
            flag = 1
        }
    }
    
    return string(res)
}

flag 用的蛮巧妙的，初始化设置成为1，然后如果放进入一个字符串后置为0，然后添加一个空格，也就是当加入一个字符串后，需要添加空格的时候，flag这个标记才为0
4. 复杂度分析：
时间复杂度：O(n) : 遍历和处理整个字符串
空间复杂度：O(n)：使用切片进行存储

5. 总结：
5.1: 不了解Go语言中的字符串库，需要去看一下标准库的使用

5.2: debug 的时候要知道自己想要程序实现什么，程序真实实现的是什么？然后去对比和排查

给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
示例：
输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
 
提示：
在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
1. Clearfication:
反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
    怎么思考这个问题呢？
        全局 局部 怎么去处理呢？

2. Coding:

3. 看题解：
使用额外空间：

func reverseWords(s string) string {
    length := len(s)
    ret := []byte{}
    
    for i := 0;i < length; {
        start := i
        for i < length && s[i] != ' ' {
            i++
        }
        
        for p := start; p < i;p++ {
            ret = append(ret, s[start + i - 1 - p])
        }
        
        for i < length && s[i] == ' ' {
            i++
            ret = append(ret, ' ')
        }
    }
    
    return string(ret)
}

里面的细节点在 start + i - 1 -p 这个点如果写不好和先不出来，还是蛮容易出错的

双指针反转单词 链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/solution/go-zi-fu-chuan-bian-li-shuang-zhi-zhen-fan-zhuan-d/
1. 遍历字符串遇到空格就停止
2. 利用双指针 实现字符串的反转
3. 最后一个字符串是没有空格的，所以当是最后一个word的时候需要特殊处理一下结束的索引位置

func reverseWords(s string) string {
    b := []byte(s)
    l := 0
    for i,v := range s {
        if v == ' ' || i == len(s) - 1 {
            r := i - 1
            if i == len(s) - 1 {
                r = i
            }
            
            for l < r {
                b[l],b[r] = b[r],b[l]
                l++
                r--
            }
            
            l = i + 1
        }
    }
    
    return string(b)
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 使用byte 数组进行存储

5. 总结：
5.1：细节点1: start + i - 1 -p  
5.2：细节点2: 最后一个字符串没有空格，需要将索引下标减回来 r := i -1 这个点
5.3:  代码设计还是蛮重要的，整体设计完了以后，细节点的设计同样是很重要的

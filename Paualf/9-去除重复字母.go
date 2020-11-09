给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同
 
示例 1：
输入：s = "bcabc"
输出："abc"
示例 2：
输入：s = "cbacdcbc"
输出："acdb"
 
提示：
• 1 <= s.length <= 10^4
• s 由小写英文字母组成

没思路：看到labuladong的题解 https://leetcode-cn.com/problems/remove-duplicate-letters/solution/you-qian-ru-shen-dan-diao-zhan-si-lu-qu-chu-zhong-/
就跟着用Go语言实现了一遍
要求：
1. 要去重
2. 去重字符串中的字符顺序不能打乱s中字符出现的相对顺序
3. 在所有符合上一条要求的去重字符串中，字典序最小的作为最终结果
忽略要求3，用栈来实现要求1和要求2:

func removeDuplicateLetters(s string) string {
    // 存放去重的结果
    stack := []rune{}
    inStack := make([]bool, 256)
    for _,c := range s {
        // 如果字符c存在栈中，直接跳过
        if inStack[c] {
            continue
        }
        stack = append(stack, c)
        inStack[c] = true
    }
     return string(stack)
}

输入 s="bcabc" 返回的是 bca,已经符合要求1和要求2了，题目希望的答案是 abc
满足要求3，保证字典序，需要做些什么修改？
在向栈 stack中插入字符 'a'的时候，我们的算法需要直到，字符 'a'的字典序和之前的两个字符'b'和'c' 相比，谁大谁小？
如果当前字符'a' 比之前的字符字典序小，就有可能需要把前面的字符pop 出栈，让'a' 排在前面，对吧？
我们修改一版代码：

func removeDuplicateLetters(s string) string {
    // 存放去重的结果
    stack := []rune{}
    inStack := make([]bool, 256)
    for _,c := range s {
        // 如果字符c存在栈中，直接跳过
        if inStack[c] {
            continue
        }
        
        // 插入之前，和之前的元素比较一下大小
        // 如果字典序比前面的小，pop前面的元素
        for len(stack) > 0 && stack[len(stack) - 1] > c {
            // 弹出栈顶元素，并把该元素标记为不在栈中
            inStack[stack[len(stack) - 1]] = false
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, c)
        inStack[c] = true
    }
     return string(stack)
}

 连续pop出比当前字符小的栈顶字符，直到栈顶元素比当前元素的字典序还小为止
这样，对于输入 s="bcabc"，我们可以得出正确结果为 "abc"了
如果改一下输入，假设s="bcac",按照刚才的算法逻辑，返回的结果是"ac",而正确的答案是“bac"，分析一下这是怎么回事？
很容易发现，因为s中只有唯一一个'b',即便字符'a'的字典序比字符'b'要小，字符'b'也不应该被pop 出去
那问题出在哪里？
我们的算法在 stack[len(stack)-1] > c 时才会 pop元素，其实这时候应该分两种情况：
1. 如果 stack[len(stack)-1] 这个字符之后还会出现，那么可以把它pop出去，反正后面还有嘛，后面再push到栈里，刚好符合字典序的要求
2. 如果 stack[len(stack)-1] 这个字符之后不会出现了，前面说了栈中不会存在重复的元素，那么就不能把它pop出去，否则你就永远失去了这个字符
s="bcac" 的例子，插入字符'a'的时候，发现前面的字符'c'的字典序比'a'大，且在'a'之后还存在字符'c'，那么栈顶的这个'c'就会被pop掉
for 循环继续判断，发现前面的字符'b'的字典序还是比'a'大，但是在'a'之后再没有字符'b'了，所以不应该把'b' pop出去。
那么关键在于：如何让算法直到字符'a'之后有几个'b'有几个'c'呢？
也不难，我们再改一版代码：

func removeDuplicateLetters(s string) string {
    // 存放去重的结果
    stack := []rune{}
    
    // 维护一个计数器记录字符串中字符的数量
    // 因为输入为ASCII 字符，大小256 够用了
    count := make([]int,256)
    
    for _,c := range s {
        count[c]++
    }
    
    inStack := make([]bool, 256)
    for _,c := range s {
        // 每遍历过一个字符，将对应的计数器减1
        count[c]--
        
        // 如果字符c存在栈中，直接跳过
        if inStack[c] {
            continue
        }
        
        // 插入之前，和之前的元素比较一下大小
        // 如果字典序比前面的小，pop前面的元素
        for len(stack) > 0 && stack[len(stack) - 1] > c {
            // 如果之后不存在栈顶元素，则停止pop
            if count[stack[len(stack) - 1 ]] == 0 {
                break
            }  
           
            // 弹出栈顶元素，并把该元素标记为不在栈中
            inStack[stack[len(stack) - 1]] = false
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, c)
        inStack[c] = true
    }
    return string(stack)
}

用了一个计数器count,当字典序较小的字符试图[挤掉]栈顶元素的时候，在count中检查栈顶元素是否是唯一的，只有当后面还存在栈顶元素的以后才能挤掉，否则不能挤掉。
至此，这个算法就结束了，时间空间复杂度都是O(N).
要求1: 通过inStack 这个布尔数组做到栈中不存在重复元素
要求2:我们顺序遍历字符串s,通过[栈]这种顺序结构的push/pop 操作记录结果字符串，保证了字符出现的顺序和s中出现的顺序一致
这里也可以想到为什么要用[栈]这种数据结构，因为先进后出的结构允许我们立即操作刚插入的字符，如果使用[队列]的话肯定是做不到的。
要求3:我们用类似单调栈的思路，配合计数器count 不断pop掉不符合最小字典序的字符，保证得到的结果字典需序最小。
在leetcode 英文站上看到漂亮的代码：

func removeDuplicateLetters(s string) string {
    dict := make(map[byte]int)
    
    for i := 0;i < len(s);i++ {
        dict[s[i]]++
    }
    
    set := make(map[byte]bool)
    var stack []byte
    
    for i := 0;i < len(s);i++ {
        dict[s[i]]--
        if !set[s[i]] {
            for len(stack) > 0 && stack[len(stack) - 1] > s[i] && dict[stack[len(stack) - 1]] > 0 {
                set[stack[len(stack)-1]] = false
                stack = stack[:len(stack) - 1]
            }
            stack = append(stack, s[i])
            set[s[i]] = true
        }
    }
    
    return string(stack)
}

复杂度分析：
时间复杂度：O(n) 遍历字符串
空间复杂度：O(n) 使用单调栈进行遍历

总结：
1. 去重，不打乱字符的相对位置，保证字典序 ，去重，使用 set 记录是否使用过，不打乱相对位置，使用栈进行解决，栈的巧妙之处先进后出，可以立即操作刚插入的字符，然后使用hash记录使用次数，来保证可以使用，同时保证字典序的前提

2. 单个知识点都不是很难，组合起来就会有点难思考，可能还是见这种题目见的少，见的多了就见怪不怪了

3. 很多优雅和优秀的解决方式和代码都是有一个演进和演化的过程的，我们如果只看到结果，而去模仿和追逐，个人可能会在追逐的过程中焦虑和迷失的。因为我们没有看到问题的本质，也没有看到中间的过程，所以培养自己看到问题本质的能力，独立思考，独立解决问题的能力是非常重要的。


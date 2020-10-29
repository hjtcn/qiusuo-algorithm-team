// 316. 去除重复字母
// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同

 

// 示例 1：

// 输入：s = "bcabc"
// 输出："abc"
// 示例 2：

// 输入：s = "cbacdcbc"
// 输出："acdb"
 

// 提示：

// 1 <= s.length <= 104
// s 由小写英文字母组成



/*
 * @lc app=leetcode.cn id=316 lang=javascript
 *
 * [316] 去除重复字母
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var removeDuplicateLetters = function(s) {
    let stack=[]
    for(let i=0;i<s.length;i++){
        //这里一直有疑惑？后来想通了。
        //前面已经判断过某个栈顶元素(出现过比当前元素大，而且后面还存在这元素的可能性)，当时就会出栈。
        //因此如果遍历到某个元素，栈中已存在，说明当前这个元素本身就打不过栈里面的元素，就不做操作，直接continue
        if(stack.indexOf(s[i])>-1) continue
        //循环判断当前元素小于栈顶元素，且字符串后面存在比栈顶元素大的，则出栈
        while(stack.length&&stack[stack.length-1]>s[i]&&s.indexOf(stack[stack.length - 1], i) > i){
            stack.pop()
        }
        stack.push(s[i])
    }
    return stack.join('')
};
// @lc code=end

/**
    刚看到这个题目真的没思路，发现贪心确实是牛，没思路感觉难死，有思路就感觉贼简单。
    思路：当当前元素<栈顶元素，且字符串后面还会出现栈顶元素，则这个栈顶元素出栈。当前元素入栈。
    复杂度分析：
    时间复杂度：O(nlogn)
    遍历查询，且会循环判断当前位置之后是否存在更合适的值。了解到indexOf第二个参数出发位置ß

    空间复杂度：O(N)
    变量声明数组，模拟栈的操作。
 */

    //递归今天没弄出来，明天再试试
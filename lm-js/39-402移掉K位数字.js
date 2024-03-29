// 402. 移掉K位数字
// 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

// 注意:

// num 的长度小于 10002 且 ≥ k。
// num 不会包含任何前导零。
// 示例 1 :

// 输入: num = "1432219", k = 3
// 输出: "1219"
// 解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
// 示例 2 :

// 输入: num = "10200", k = 1
// 输出: "200"
// 解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
// 示例 3 :

// 输入: num = "10", k = 2
// 输出: "0"
// 解释: 从原数字移除所有的数字，剩余为空就是0。


/**
 * @param {string} num
 * @param {number} k
 * @return {string}
 */
var removeKdigits = function (num, k) {
    let stack = []
    for (let i = 0; i < num.length; i++) {
        //遍历，当当前元素小于栈顶元素，stack弹出栈顶元素。使栈保持从小到大的排列
        while (k > 0 && stack.length > 0 && num[i] < stack[stack.length - 1]) {
            stack.pop()
            k--
        }
        stack.push(num[i])
    }
    //如果k没有删除完指定数量的元素，就从栈接着弹出栈顶元素
    while (k > 0) {
        stack.pop()
        k--
    }
    //此时其实已经得到结果了，但是不要忘了去除前导零
    while (stack.length && stack[0] == '0') {
        stack.shift()
    }
    //  如果栈为空，返回0，反之返回栈拼成的字符串
    return stack.join('') || '0'

};

/** 题解
 * 这个题是查看题解后得到了思路。
 贪心。通过出栈入栈来保持栈作为从小到大的排列，如果k删除完了，就只管把后面的元素入栈即可。
 如果循环结束，k仍然没有删除完毕，就需要出栈操作了，从后面移除较大元素，剩余元素就会是较小的数字组成了。
 由于是数字，也要考虑前导零问题，把前面的零都给去掉，此时的栈再拼接为字符串即为结果。
 若字符串为空，则返回0
复杂度分析：
  时间复杂度：O(N)
  遍历一遍字符串,元素最多有一次入栈出栈操作

  空间复杂度：O(N)
  变量声明数组，模拟栈的操作。
*/
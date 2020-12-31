<?php

/*
 * @lc app=leetcode.cn id=844 lang=php
 *
 * [844] 比较含退格的字符串
 *
 * https://leetcode-cn.com/problems/backspace-string-compare/description/
 *
 * algorithms
 * Easy (51.04%)
 * Likes:    232
 * Dislikes: 0
 * Total Accepted:    59K
 * Total Submissions: 112.9K
 * Testcase Example:  '"ab#c"\n"ad#c"'
 *
 * 给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
 * 
 * 注意：如果对空文本输入退格字符，文本继续为空。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：S = "ab#c", T = "ad#c"
 * 输出：true
 * 解释：S 和 T 都会变成 “ac”。
 * 
 * 
 * 示例 2：
 * 
 * 输入：S = "ab##", T = "c#d#"
 * 输出：true
 * 解释：S 和 T 都会变成 “”。
 * 
 * 
 * 示例 3：
 * 
 * 输入：S = "a##c", T = "#a#c"
 * 输出：true
 * 解释：S 和 T 都会变成 “c”。
 * 
 * 
 * 示例 4：
 * 
 * 输入：S = "a#c", T = "b"
 * 输出：false
 * 解释：S 会变成 “c”，但 T 仍然是 “b”。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S 和 T 只含有小写字母以及字符 '#'。
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 你可以用 O(N) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？
 * 
 * 
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $S
     * @param String $T
     * @return Boolean
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了75.29%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了71.01%的用户
    function backspaceCompare1($S, $T) {
        // 获取两个字符串长度
        $len1 = strlen($S);
        $len2 = strlen($T);
        // 声明两个数组，来作为栈
        $stack1 = [];
        $stack2 = [];
        // 分别遍历两个字符串
        for($i=0; $i<$len1; $i++) {
            // 如果是#号，就出栈，
            if($S[$i] == '#') {
                array_pop($stack1);
            }else{
                array_push($stack1, $S[$i]); // 如果是字母就进栈
            }
        }
        for($i=0; $i<$len2; $i++) {
            // 如果是#号，就出栈，
            if($T[$i] == '#') {
                array_pop($stack2);
            }else{
                array_push($stack2, $T[$i]); // 如果是字母就进栈
            }
        }
        // 获取两个数组
        return $stack1 == $stack2;
        /*
        $size1 = count($stack1);
        $size2 = count($stack2);
        if($size1 == 0 && $size2 == 0) {
            return true;// 最后都是空字符串
        }elseif($size1 != $size2) {
            return false;// 长度都不相等，肯定不等
        }
        // 逐一比较两个字符串
        foreach($stack1 as $key => $value) {
            if($value != $stack2[$key]) {
                return false;
            }
        }
        return true;
        */
    }

    // 双指针
    // 执行用时：4 ms, 在所有 PHP 提交中击败了97.65%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了62.13%的用户
    function backspaceCompare($S, $T) {
        // 定义i j 双指针
        $i     = strlen($S)-1; 
        $j     = strlen($T)-1;
        // 定义需要跳过的字符数量
        $skipS = 0;
        $skipT = 0;
        // 遍历条件
        while($i >= 0 || $j >= 0) {
            // 遍历S
            while($i >= 0) {
                if($S[$i] == '#') {
                    // 如果是#号退格字符
                    $skipS++;
                    $i--;
                } elseif($skipS >0 ) {
                    // 如果是普通字符，并且需要跳过
                    $skipS--;
                    $i--;
                } else {
                    // 如果是普通字符，而且不需要跳过的
                    break;
                }
            }
            // 遍历T,同理
            while($j >= 0) {
                if($T[$j] == '#') {
                    $skipT++;
                    $j--;
                } elseif($skipT >0) {
                    $skipT--;
                    $j--;
                } else {
                    break;
                }
            }
            // 如果是普通字符，而且不需要跳过的，就进行比较
            if($i >=0 && $j>=0 ) {
                if($S[$i] != $T[$j]) {
                    return false; // 不相等直接退出
                }
            }elseif($i >=0 || $j>=0 ) {
                return false;
            }
            $i--;
            $j--;
        }
        return true;
    }
}
// @lc code=end

/*

【1】使用数组模拟栈

最容易想到的，还原字符串的格式。

- 遍历字符串
- 如果遇到#，就对栈顶元素进行弹出
- 反之，就将元素压入栈中
- 最后栈中的元素就是，退格后的字符串

然后再比较两个栈内容是否相等，从而比较最后的字符串是否相等。

申请了两个栈空间，遍历字符串
空间复杂度O(N+M)
时间复杂度O(N+M)


【2】双指针

双指针 真的牛逼 每次都能更好的降低空间复杂度

例如字符串ab#c#d
我们可以知道当前字符是否被删除，是取决于后面的#符号的，因此可以逆序解决这个问题

定义skip变量，标识当前待删除的字符的数量
- 逆序遍历字符串
- 如果当前为#，我们就skip++，表示要多删除一个字符
- 如果当前为字母，
    - 我们判断skip==0 ，当前字符不需要删除
    - skip>0 ,当前字符需要删除，并且skip--

时间复杂度：O(N+M) 我们需要遍历两字符串各一次。
空间复杂度：O(1) 我们采用指针跟计数器
*/
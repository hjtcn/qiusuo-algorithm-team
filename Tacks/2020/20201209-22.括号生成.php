<?php
/*
 * @Descripttion: 周三 括号生成
 * @Author: wangtao10@qianxin.com
 * @Date: 2020-12-09 09:50:05
 * @LastEditTime: 2020-12-09 11:36:56
 */


/*
 * @lc app=leetcode.cn id=22 lang=php
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.45%)
 * Likes:    1460
 * Dislikes: 0
 * Total Accepted:    205.8K
 * Total Submissions: 268.9K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 * 
 * 
 */

// @lc code=start
class Solution {

    private $strings = [];
    /**
     * @param Integer $n
     * @return String[]
     */
    function generateParenthesis($n) {
        $this->_generate($n, 0, 0, '');
        return $this->strings;
    }

    /**
     * 递归子程序
     *
     * @param [int] $n 生成括号对数
     * @param [int] $left  左括号数
     * @param [int] $right 右括号数
     * @param [string] $s  括号
     */
    function _generate($n, $left, $right, $s)
    {
        // 1、递归终止条件
        if($left == $n && $right == $n) {
            $this->strings[] = $s;
            return ;
        }

        // 2、处理当前层逻辑

        // 3、下一层
        if($left < $n) {
            // 增加左括号
            $this->_generate($n, $left+1, $right, $s.'(');
        }
        if($right < $left) {
            // 右括号+1
            $this->_generate($n, $left, $right+1, $s. ')');
        }
        
        // 4、递归结束清理当前层
    }

}
// @lc code=end



/*

递归四件套

1、递归终止条件，防止陷入死循环
2、递归内部处理逻辑
3、递归进下一层
4、递归清除，比如清楚一些无用的变量，分情况是否需要

*/
<?php
/*
 * @Descripttion: 周三的题 补一下
 * @Author: tacks321@qq.com
 * @Date: 2021-01-17 16:43:15
 */



/*
 * @lc app=leetcode.cn id=590 lang=php
 *
 * [590] N叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (75.17%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    44.5K
 * Total Submissions: 59.1K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的后序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其后序遍历: [5,6,3,2,4,1].
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * class Node {
 *     public $val = null;
 *     public $children = null;
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->children = array();
 *     }
 * }
 */

class Solution {
    /**
     * @param Node $root
     * @return integer[]
     */
    // 执行用时：20 ms, 在所有 PHP 提交中击败了65.00%的用户
    // 内存消耗：18.8 MB, 在所有 PHP 提交中击败了75.00%的用户
    function postorder($root) {
        $result = [];
        $this->_helper($root, $result);
        $result[] = $root->val;

        return $result;
    }

    // 辅助函数 递归处理
    private function _helper($node, &$result)
    {   
        if($node == null) {
            return [];
        }
        foreach($node->children as $child) {
            $this->_helper($child, $result);
            $result[] = $child->val; // 后序遍历 最后访问根节点
        }
    }
}
// @lc code=end

/*
    也是有点类似原题，然后处理就是 2 => N  如何去解决多叉树的一个思想，
*/
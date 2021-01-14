<?php
/*
 * @Descripttion: 补充一下 周一的题，今天还没来得及写
 * @Author: tacks321@qq.com
 * @Date: 2021-01-14 08:58:29
 */


/*
 * @lc app=leetcode.cn id=103 lang=php
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (56.99%)
 * Likes:    368
 * Dislikes: 0
 * Total Accepted:    109.6K
 * Total Submissions: 192.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回锯齿形层序遍历如下：
 * 
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[][]
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了26.83%的用户
    // 内存消耗：15.5 MB, 在所有 PHP 提交中击败了88.62%的用户
    function zigzagLevelOrder($root) {
        $res = [];
        $this->_helper($root, 0, $res);
        return $res;
    }

    function _helper($node, $level, &$res) {
        if($node == null) {
            return ;
        }
        if($level == count($res)) {
            $res[$level] = [];
        }
        // 核心代码
        if($level % 2 == 0 ) {
            array_push($res[$level], $node->val);
        } else {
            array_unshift($res[$level], $node->val);
        }

        $this->_helper($node->left, $level+1, $res);
        $this->_helper($node->right, $level+1, $res);
    }
}
// @lc code=end


/*

可以对比 【102. 二叉树的层序遍历】 这道题。https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

其实就是在控制层序遍历的时候，是从左到右，还是从右到左 增加一个 level%2 的做法


==================

*/

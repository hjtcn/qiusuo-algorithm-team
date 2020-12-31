<?php
/*
 * @Descripttion: 周二 🙂
 * @Author: Tacks
 * @Date: 2020-12-01 09:56:31
 * @LastEditTime: 2020-12-01 10:51:50
 */
/*
 * @lc app=leetcode.cn id=145 lang=php
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Medium (73.44%)
 * Likes:    483
 * Dislikes: 0
 * Total Accepted:    166.4K
 * Total Submissions: 226.1K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]  
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3 
 * 
 * 输出: [3,2,1]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    // 迭代 执行用时：8 ms, 在所有 PHP 提交中击败了67.01%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了26.29% 的用户
    function postorderTraversal1($root) {
        if(!$root) {
            return [];
        }
        $stack = [$root]; // 树的根节点压入栈中
        $res   = [];
        while($stack) {
            $node = array_pop($stack);
            // 倒着输入输出
            array_unshift($res, $node->val); 
            // 若栈不为空，将根节点推出栈，并将栈的不为空的左右节点推入栈中
            $node->left  && array_push($stack, $node->left);
            $node->right && array_push($stack, $node->right);
        }
        return $res;
    }

    // 递归
    // 执行用时：8 ms, 在所有 PHP 提交中击败了67.01%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了9.79%的用户
    function postorderTraversal($root) {
        $res = [];
        $this->traversal($root, $res);
        return $res;
    }

    private function traversal($node, &$result)
    {
        if(!$node) {
            return null;
        }
        // 左 右 中
        $this->traversal($node->left, $result);
        $this->traversal($node->right, $result);
        $result[] = $node->val;
    }


}
// @lc code=end


/*
二叉树的后序遍历：按照访问左子树——右子树——根节点的方式遍历这棵树。


给定一个二叉树，返回它的后序遍历。 两种思路


【迭代】
【递归】
    时间复杂度：O(n)，其中 n 是二叉搜索树的节点数。每一个节点恰好被遍历一次。
    空间复杂度：O(n)，为递归过程中栈的开销，平均情况下为 O(log n)最坏情况下树呈现链状，为 O(n)。

    


*/
<?php
/*
 * @Descripttion: 补一下周四的题
 * @Author: wangtao10@qianxin.com
 * @Date: 2020-12-04 18:56:48
 * @LastEditTime: 2020-12-04 19:02:56
 */


/*
 * @lc app=leetcode.cn id=404 lang=php
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (56.19%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    63.1K
 * Total Submissions: 112.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 计算给定二叉树的所有左叶子之和。
 * 
 * 示例：
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 * 
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

    protected $sum = 0;
    
    /**
     * @param TreeNode $root
     * @return Integer
     */
    function sumOfLeftLeaves($root) {
        if($root == null) {
            return $this->sum;
        }

        $this->dfs($root);

        return $this->sum;
        
    }

    private function dfs($root) {
        if($root == null) {
            return;
        }

        if($root->left !== null) {
            if($root->left->left == null && $root->left->right == null) {
                $this->sum += $root->left->val;
            }

            $this->dfs($root->left);
        }

        if($root->right !== null) {
            $this->dfs($root->right);
        } 

    }

}
// @lc code=end

/*

一个节点为「左叶子」节点，当且仅当它是某个节点的左子节点，并且它是一个叶子结点

当我们遍历到节点 node 时，如果它的左子节点是一个叶子结点，那么就将它的左子节点的值累加计入答案

深度dfs

广度bfs


*/
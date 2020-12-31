<?php
/*
 * @Descripttion: 周一冲冲冲
 * @Author: wangtao10@qianxin.com
 * @Date: 2020-12-07 08:54:33
 * @LastEditTime: 2020-12-07 18:37:30
 */


/*
 * @lc app=leetcode.cn id=112 lang=php
 *
 * [112] 路径总和
 *
 * https://leetcode-cn.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (51.38%)
 * Likes:    473
 * Dislikes: 0
 * Total Accepted:    155.9K
 * Total Submissions: 302.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例: 
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \      \
 * ⁠       7    2      1
 * 
 * 
 * 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
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
     * @param Integer $sum
     * @return Boolean
     */
    // 执行用时：16 ms, 在所有 PHP 提交中击败了45.07%的用户
    // 内存消耗：17.2 MB, 在所有 PHP 提交中击败了5.71%的用户
    function hasPathSum($root, $sum) {
        if($root === null) {
            return false;
        }
        // 如果当前是叶子节点 $sum减去叶子节点的值为0 满足条件
        if($root->left === null && $root->right === null) {
            return $root->val == $sum;
        }
        // 递归大问题转化为小问题
        // 根据输入值（22）-当前节点的值（5）得出剩余数（17）作为下一次递归的参数
        // 继续递归 左右节点
        return $this->hasPathSum($root->left, $sum-$root->val) || $this->hasPathSum($root->right, $sum-$root->val);
    }
}
// @lc code=end

/*

核心思想是对树进行一次遍历，在遍历时记录从根节点到当前节点的路径和。

【1】DFS 真的牛逼 递归用起来很爽 代码量又少
时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次
空间复杂度：O(H)，其中 H 是树的高度。空间复杂度主要取决于递归时栈空间的开销，最坏情况下，树呈现链状，空间复杂度为 O(N)。平均情况下树的高度与节点数的对数正相关，空间复杂度为 O(logN)。


*/
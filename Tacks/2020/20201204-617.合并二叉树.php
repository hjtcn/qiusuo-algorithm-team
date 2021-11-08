<?php
/*
 * @Descripttion: 周五的题
 * @Date: 2020-12-04 19:04:52
 * @LastEditTime: 2021-11-08 09:57:38
 */


/*
 * @lc app=leetcode.cn id=617 lang=php
 *
 * [617] 合并二叉树
 *
 * https://leetcode-cn.com/problems/merge-two-binary-trees/description/
 *
 * algorithms
 * Easy (78.36%)
 * Likes:    583
 * Dislikes: 0
 * Total Accepted:    110.3K
 * Total Submissions: 140.6K
 * Testcase Example:  '[1,3,2,5]\n[2,1,3,null,4,null,7]'
 *
 * 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
 * 
 * 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL
 * 的节点将直接作为新二叉树的节点。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 
 * Tree 1                     Tree 2                  
 * ⁠         1                         2                             
 * ⁠        / \                       / \                            
 * ⁠       3   2                     1   3                        
 * ⁠      /                           \   \                      
 * ⁠     5                             4   7                  
 * 输出: 
 * 合并后的树:
 * 3
 * / \
 * 4   5
 * / \   \ 
 * 5   4   7
 * 
 * 
 * 注意: 合并必须从两个树的根节点开始。
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
     * @param TreeNode $t1
     * @param TreeNode $t2
     * @return TreeNode
     */
    function mergeTrees($t1, $t2) {
        // 临界条件
        if($t1 == null) return $t2;
        if($t2 == null) return $t1;
        // 新的值
        $t1->val += $t2->val;
        // 新的左树

        // 可以使用深度优先搜索合并两个二叉树。从根节点开始同时遍历两个二叉树，并将对应的节点进行合并
        $t1->left = $this->mergeTrees($t1->left, $t2->left);
        $t1->right = $this->mergeTrees($t1->right, $t2->right);
        
        return $t1;
    }
}
// @lc code=end


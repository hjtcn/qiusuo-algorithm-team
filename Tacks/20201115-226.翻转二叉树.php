<?php

/*
    2020/11/15 周日
*/



/*
 * @lc app=leetcode.cn id=226 lang=php
 *
 * [226] 翻转二叉树
 *
 * https://leetcode-cn.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (77.44%)
 * Likes:    683
 * Dislikes: 0
 * Total Accepted:    160.1K
 * Total Submissions: 206.5K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 翻转一棵二叉树。
 * 
 * 示例：
 * 
 * 输入：
 * 
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 * 
 * 输出：
 * 
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 * 
 * 备注:
 * 这个问题是受到 Max Howell 的 原问题 启发的 ：
 * 
 * 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
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
     * @return TreeNode
     */
    // 递归交换两颗子树
    // 执行用时：4 ms, 在所有 PHP 提交中击败了94.47%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了13.39%的用户
    function invertTree($root) {
        if(!$root) {
            return null;
        }
        $left = $this->invertTree($root->left);
        $right= $this->invertTree($root->right);
        // 交换两颗子树位置
        $root->left = $right;
        $root->right= $left;
        return $root;
    }
}
// @lc code=end

/*
【递归】
    递归地对树进行遍历
        从叶子结点先开始翻转
        如果当前遍历到的节点 root 的左右两棵子树都已经翻转，那么我们只需要交换两棵子树的位置
        即可完成以 root 为根节点的整棵子树的翻转
时间复杂度：O(N)
    其中 N 为二叉树节点的数目。
    我们会遍历二叉树中的每一个节点，对每个节点而言，我们在常数时间内交换其两棵子树。
空间复杂度：O(N)
    使用的空间由递归栈的深度决定，它等于当前节点在二叉树中的高度。
    在平均情况下，二叉树的高度与节点个数为对数关系，即 O(logN)。
    而在最坏情况下，树形成链状，空间复杂度为 O(N)。
 

*/
<?php
/*
 * @Description: 今天周四了，偷懒了三天没刷题，罪恶感十足
 * @Author: Tacks
 * @Date: 2020-11-19 14:45:26
 */
/*
 * @lc app=leetcode.cn id=110 lang=php
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (54.79%)
 * Likes:    524
 * Dislikes: 0
 * Total Accepted:    150.4K
 * Total Submissions: 274.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 * 
 * 本题中，一棵高度平衡二叉树定义为：
 * 
 * 
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：root = []
 * 输出：true
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数在范围 [0, 5000] 内
 * -10^4 
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
     * @return Boolean
     */
    // 自顶向下
    // 执行用时：16 ms, 在所有 PHP 提交中击败了74.29%的用户
    // 内存消耗：17.3 MB, 在所有 PHP 提交中击败了11.63%的用户
    // function isBalanced($root) {
    //     if(!$root) {
    //         return true;
    //     } else {
    //         // 对当前的节点 前序遍历
    //         // 如果当前节点的左右子树的高度差不超过1，再去分别遍历左子树跟右子树
    //         return abs($this->height($root->left) - $this->height($root->right)) <= 1 && $this->isBalanced($root->left) && $this->isBalanced($root->right);
    //     }
    // }

    // // 计算节点的高度
    // function height($p) {
    //     if(!$p) {
    //         return 0;
    //     } else {
    //         // p是非空节点
    //         return max($this->height($p->left), $this->height($p->right)) + 1;
    //     }
    // }


    // 自底向上
    // 执行用时：20 ms, 在所有 PHP 提交中击败了43.43%的用户
    // 内存消耗：17.7 MB, 在所有 PHP 提交中击败了5.23%的用户
    function isBalanced($root) {
        return $this->height($root) >=0 ;
    }

    // 计算节点的高度
    function height($p) {
        if(!$p) {
            return 0;
        } 
        $leftHeight = $this->height($p->left);// 左子树高度
        $rightHeight= $this->height($p->right);// 右子树高度
        // 如果一棵子树是平衡的，则返回其高度（高度一定是非负整数），否则返回 −1。
        // 如果存在一棵子树不平衡，则整个二叉树一定不平衡。
        if(abs($leftHeight - $rightHeight) > 1 || $leftHeight == -1 || $rightHeight == -1) {
            return -1;
        } else {
            return max($leftHeight, $rightHeight) + 1;
        }
    }

}
// @lc code=end

/*

平衡二叉树：二叉树的每个节点的左右子树的高度差绝对值不超过1

=》 根据定义，我们知道一颗是平衡二叉树，那么他的左右子树肯定也是平衡二叉树，因此可以采用递归的形式

【递归】递归的顺序，可以分为自顶向下或者自底向上。（之前Paualf给我提的建议，可以有这两种顺序来进行解决递归）

【自顶向下】
    对当前的节点 前序遍历
    如果当前节点的左右子树的高度差不超过1，再去分别遍历左子树跟右子树
    
    时间复杂度：O(n^2) 其中 n 是二叉树中的节点个数
            对于节点 p，如果它的高度是 d，height(p) 最多会被重复调用 d 次（即遍历到它的每一个祖先节点时）
    空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过 n。
【自底向上】
    对于每个节点，函数 height 只会被调用一次
    后序遍历，对于当前遍历到的节点，先递归地判断其左右子树是否平衡，再判断以当前节点为根的子树是否平衡

    时间复杂度：O(n)，其中 n 是二叉树中的节点个数。使用自底向上的递归
    空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过 n。



*/
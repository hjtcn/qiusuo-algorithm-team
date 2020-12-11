<?php
/*
 * @Descripttion: 周二牛逼，我今天是自己的带的午饭
 * @Author: wangtao10@qianxin.com
 * @Date: 2020-12-08 14:04:52
 * @LastEditTime: 2020-12-08 15:49:36
 */

/*
 * @lc app=leetcode.cn id=235 lang=php
 *
 * [235] 二叉搜索树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (65.86%)
 * Likes:    498
 * Dislikes: 0
 * Total Accepted:    112.5K
 * Total Submissions: 170.8K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 * 
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 * 
 * 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
 * 
 * 
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * 输出: 6 
 * 解释: 节点 2 和节点 8 的最近公共祖先是 6。
 * 
 * 
 * 示例 2:
 * 
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * 输出: 2
 * 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 * 
 * 
 * 
 * 说明:
 * 
 * 
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉搜索树中。
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
     * @param TreeNode $p
     * @param TreeNode $q
     * @return TreeNode
     */
    // 执行用时：32 ms, 在所有 PHP 提交中击败了93.01%的用户
    // 内存消耗：18.7 MB, 在所有 PHP 提交中击败了10.22%的用户
    function lowestCommonAncestor1($root, $p, $q) {
        // 如果当前节点的值大于 p 和 q 的值，
        // 说明 p 和 q 应该在当前节点的左子树，因此将当前节点移动到它的左子节点；
        if($root->val > $p->val && $root->val > $q->val) {
            return $this->lowestCommonAncestor($root->left, $p, $q);
        }    
        // 如果当前节点的值小于 p 和 q 的值，
        // 说明 p 和 q 应该在当前节点的右子树，因此将当前节点移动到它的右子节点；
        if($root->val < $p->val && $root->val < $q->val) {
            return $this->lowestCommonAncestor($root->right, $p, $q);
        }
        // 否则就是分岔口，也就是当前最近公共祖先
        return $root;
    }

    // 执行用时：36 ms, 在所有 PHP 提交中击败了76.34%的用户
    // 内存消耗：18.6 MB, 在所有 PHP 提交中击败了12.37%的用户
    function lowestCommonAncestor($root, $p, $q) {
        // 肯定都位于根节点的同一侧
        while( ($root->val - $p->val) * ($root->val - $q->val)  >0 ) {  
            $root = $p->val < $root->val ? $root->left : $root->right;
        }
        return $root;
    }


}
// @lc code=end


/*
【二叉搜索树】Binary Search Tree  BST

递归处理
迭代处理


*/

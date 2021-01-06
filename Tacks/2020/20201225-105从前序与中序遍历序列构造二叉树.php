<?php
/*
 * @Descripttion: 周三的题
 * @Author: tacks321@qq.com
 * @Date: 2020-12-25 18:09:03
 * @LastEditTime: 2020-12-25 18:34:13
 */

/*
 * @lc app=leetcode.cn id=105 lang=php
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (68.41%)
 * Likes:    805
 * Dislikes: 0
 * Total Accepted:    138.7K
 * Total Submissions: 202K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
class TreeNode {
    public $val   = null;
    public $left  = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}
class Solution {

    /**
     * @param Integer[] $preorder
     * @param Integer[] $inorder
     * @return TreeNode
     */
    // 执行用时：96 ms, 在所有 PHP 提交中击败了37.04%的用户
    // 内存消耗：305.1 MB, 在所有 PHP 提交中击败了16.05%的用户
    // todo:: 笨办法
    function buildTree($preorder, $inorder) {
        if(empty($preorder)) {
            return null;
        }
        // 前序遍历  第一个节点就是根节点
        $preRootValue = $preorder[0];
        // 先根据根节点建立起来二叉树
        $root = new TreeNode($preRootValue);
        // 在中序遍历中找到根节点的位置
        $inRootIndex = array_search($preRootValue, $inorder);

        // 用于构造左子树的中序遍历序列
        $leftInOrder = array_slice($inorder, 0, $inRootIndex);
        // 用于构造左子树的中序遍历序列
        $rightInOrder = array_slice($inorder, $inRootIndex+1);

        // 用于构造左子树的前序遍历序列
        $leftPreOrder = array_slice($preorder, 1, $inRootIndex);
        $rightPreOrder = array_slice($preorder, $inRootIndex+1);

        // 递归 构造左子树
        $root->left  = $this->buildTree($leftPreOrder, $leftInOrder);
        // 递归 构造右子树
        $root->right = $this->buildTree($rightPreOrder, $rightInOrder);
        return $root;
    }
}
// @lc code=end



/*

二叉树

前序遍历： 根节点-> 左子树 -> 右子树
中序遍历： 根节点-> 左子树 -> 右子树


前序遍历中第一个节点就是根节点




*/
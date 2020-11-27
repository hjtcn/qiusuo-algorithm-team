<?php
/*
 * @Description: 周三的题
 * @Author: Tacks
 * @Date: 2020-11-19 18:03:29
 */
/*
 * @lc app=leetcode.cn id=94 lang=php
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (73.99%)
 * Likes:    780
 * Dislikes: 0
 * Total Accepted:    301.3K
 * Total Submissions: 406.3K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：root = [1]
 * 输出：[1]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：root = [1,2]
 * 输出：[2,1]
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：root = [1,null,2]
 * 输出：[1,2]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数目在范围 [0, 100] 内
 * -100 
 * 
 * 
 * 
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

    public $res = [];
    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    // 递归 中序遍历
    // 执行用时：8 ms, 在所有 PHP 提交中击败了64.74%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了13.55%的用户
    function inorderTraversal1($root) {
        if($root) {
            $this->inorderTraversal1($root->left);
            array_push($this->res, $root->val);
            $this->inorderTraversal1($root->right);
        }
        return $this->res;
    }

    // 迭代 中序遍历
    // 执行用时：8 ms, 在所有 PHP 提交中击败了64.74%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了16.77%的用户
    function inorderTraversal2($root) {
        $stack = [];
        $res   = [];
        $curr  = $root;
        while($curr != null || $stack) {
            while($curr != null) {
                // 左子树节点全部放到栈里
                array_push($stack, $curr);
                $curr = $curr->left; // 左子树
            }
            $curr = array_pop($stack);
            $res[] = $curr->val;// 中间
            $curr = $curr->right;// 右边
        }
        return $res;
    }

}
// @lc code=end



/*

【二叉树中序遍历】
    按照访问左子树——根节点——右子树的方式遍历这棵树
    
【递归】

    时间复杂度：O(n)，其中 n 为二叉树节点的个数。
    空间复杂度：O(n)。空间复杂度取决于递归的栈深度。
【迭代】
    借助栈来存储
    时间复杂度：O(n)，其中 n 为二叉树节点的个数。
    空间复杂度：O(n)。空间复杂度取决于递归的栈深度。

【Morris 中序遍历】
    这个算法还没看明白？？？待做
*/
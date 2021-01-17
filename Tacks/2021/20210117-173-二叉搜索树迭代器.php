<?php
/*
 * @Descripttion: 周四的题 补一下
 * @Author: tacks321@qq.com
 * @Date: 2021-01-17 16:53:46
 */



/*
 * @lc app=leetcode.cn id=173 lang=php
 *
 * [173] 二叉搜索树迭代器
 *
 * https://leetcode-cn.com/problems/binary-search-tree-iterator/description/
 *
 * algorithms
 * Medium (75.59%)
 * Likes:    314
 * Dislikes: 0
 * Total Accepted:    36.1K
 * Total Submissions: 47.7K
 * Testcase Example:  '["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]\n' +
  '[[[7,3,15,null,null,9,20]],[],[],[],[],[],[],[],[],[]]'
 *
 * 实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
 * 
 * 调用 next() 将返回二叉搜索树中的下一个最小的数。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // 返回 3
 * iterator.next();    // 返回 7
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 9
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 15
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 20
 * iterator.hasNext(); // 返回 false
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
 * 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
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
class BSTIterator {

    // 维护一个栈
    private $stack = [];

    /**
     * @param TreeNode $root
     */
    // 执行用时：44 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：22.9 MB, 在所有 PHP 提交中击败了100.00%的用户
    function __construct($root) {
        while($root != null) {
            $this->stack[] = $root;
            $root = $root->left;
        }

    }

    /**
     * @return Integer
     */
    function next() {
        $root = array_pop($this->stack);
        $val = $root->val;
        $root= $root->right;

        while($root != null) {
            $this->stack[] = $root;
            $root = $root->left;
        }

        return $val;
    }

    /**
     * @return Boolean
     */
    function hasNext() {
        return !empty($this->stack);
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * $obj = BSTIterator($root);
 * $ret_1 = $obj->next();
 * $ret_2 = $obj->hasNext();
 */
// @lc code=end


/*
二叉搜索树
    二叉搜索树的中序序列是升序序列
*/
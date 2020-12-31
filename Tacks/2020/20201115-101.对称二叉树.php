<?php
/*
    2020-11-15 周日
*/

/*
 * @lc app=leetcode.cn id=101 lang=php
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (53.10%)
 * Likes:    1116
 * Dislikes: 0
 * Total Accepted:    229.1K
 * Total Submissions: 431.2K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 
 * 
 * 
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 
 * 
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 你可以运用递归和迭代两种方法解决这个问题吗？
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
     * @return Boolean
     */
    // // 递归
        // 执行用时：4 ms, 在所有 PHP 提交中击败了98.84%的用户
        // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了14.71%的用户
    // function isSymmetric($root) {
    //     return $this->_check($root, $root);
    // }

    // // 检查对称性
    // private function _check($p, $q) {
    //     if($p == null && $q == null) {
    //         return true;
    //     }
    //     if($p == null || $q == null ) {
    //         return false;
    //     }

    //     return $p->val == $q->val && $this->_check($p->left, $q->right) && $this->_check($p->right, $q->left);
    // }


    // 迭代 
    
     


    

}
// @lc code=end

/*
对称概念
    如果一个树的左子树与右子树镜像对称，那么这个树是对称的

对称条件
    它们的两个根结点具有相同的值
    每个树的右子树都与另一个树的左子树镜像对称

【递归】

    时间复杂度：
        这里遍历了这棵树，渐进时间复杂度为 O(n)。

    空间复杂度：
        这里的空间复杂度和递归使用的栈空间有关，这里递归层数不超过 n，故渐进空间复杂度为 O(n)。
    
【迭代】

    递归改成迭代，可以借助队列

    通过 把根节点入队两次，然后每次提取两个结点比较对应的值

    然后将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，或者我们检测到树不对称，结束程序

*/
<?php
/*
 * @Description: 周四的题
 * @Author: Tacks
 * @Date: 2020-11-19 18:41:17
 */

/*
 * @lc app=leetcode.cn id=222 lang=php
 *
 * [222] 完全二叉树的节点个数
 *
 * https://leetcode-cn.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (73.10%)
 * Likes:    279
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 53.1K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给出一个完全二叉树，求出该树的节点个数。
 * 
 * 说明：
 * 
 * 
 * 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第
 * h 层，则该层包含 1~ 2^h 个节点。
 * 
 * 示例:
 * 
 * 输入: 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 * 
 * 输出: 6
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
     * @return Integer
     */
    // 暴力破解
    // 执行用时：40 ms, 在所有 PHP 提交中击败了38.24%的用户
    // 内存消耗：22.6 MB, 在所有 PHP 提交中击败了14.71%的用户
    function countNodes1($root) {
        return $root==null ? 0 : $this->countNodes($root->left) + $this->countNodes($root->right) + 1;
    }

    // 完全二叉树特点
    // 执行用时：28 ms, 在所有 PHP 提交中击败了91.18%的用户
    // 内存消耗：22.4 MB, 在所有 PHP 提交中击败了26.47%的用户
    function countNodes($root) {
        if(!$root) return 0;
        
        $ld = $this->getDepth($root->left);
        $rd = $this->getDepth($root->right);

        if($ld == $rd) {
            return $this->countNodes($root->right) + (1 << $ld);
        }else {
            return $this->countNodes($root->left) +  (1 << $rd);
        }
    }

    // 求高度
    function getDepth($root) {
        $depth = 0;
        while($root) {
            $depth++;
            $root = $root->left;
        }
        return $depth;
    }
}
// @lc code=end

/*
【暴力破解】
    一层一层递归，进行计数
【完全二叉树的特性】
完全二叉树中，除了最后一层外，其余每层节点都是满的，并且最后一层的节点全部靠向左边。

完全二叉树分为
    满二叉树： 一个二叉树的层数为K，且结点总数是(2^k) -1 ，则它就是满二叉树。
    不满二叉树

完全二叉树求解：
    左右自述高度一致，则左子树一定是满二叉树,反之，右子树就是满二叉树

*/
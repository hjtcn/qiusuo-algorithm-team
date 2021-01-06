<?php
/*
 * @Descripttion: 周一的题  冲冲冲，跟小马姐老黑小珍一起吃的饺子
 * @Author: tacks321@qq.com
 * @Date: 2020-12-24 16:52:29
 * @LastEditTime: 2020-12-24 17:30:24
 */

/*
 * @lc app=leetcode.cn id=98 lang=php
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (32.65%)
 * Likes:    868
 * Dislikes: 0
 * Total Accepted:    203.3K
 * Total Submissions: 614.8K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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
    // 内存消耗：18 MB, 在所有 PHP 提交中击败了28.97%的用户
    // 执行用时：16 ms, 在所有 PHP 提交中击败了65.42%的用户
    function isValidBST1($root) {
        // 结果数组
        $result = [];
        // 辅助函数 中序遍历二叉树
        $this->_helper1($root, $result);        
        
        // 判断是否是增数组
        for($i=0; $i<count($result)-1; $i++) {
            if($result[$i] >= $result[$i+1]) {
                return false;
            }
        }
        return true;
    }


    // 中序遍历
    function _helper1($node, &$result)
    {
        // 终止递归条件
        if($node == null) {
            return $result;
        }
        // 中序
        $this->_helper1($node->left, $result);
        $result[] = $node->val; // 根节点
        $this->_helper1($node->right, $result);
        
    }




    // 执行用时：12 ms, 在所有 PHP 提交中击败了92.52%的用户
    // 内存消耗：17.8 MB, 在所有 PHP 提交中击败了61.68%的用户
    function isValidBST($root) {
        return $this->_helper($root, null, null);
    }

    // 中序遍历
    function _helper($node, $lower, $upper)
    {
        // 终止递归条件
        if($node === null) {
            return true;
        }
        // 当前值
        $val = $node->val;

        if($lower !== null && $val <= $lower) return false;
        if($upper !== null && $val >= $upper) return false;

        // 递归调用左子树 ，上界upper 改为 root.val，左子树里所有节点的值均小于它的根节点的值
        if(!$this->_helper($node->left, $lower, $val)) {
            return false;
        }
        // 递归调用右子树
        if(!$this->_helper($node->right, $val, $upper)) {
            return false;
        }
        return true;
    }
    


}
// @lc code=end



/*
真不错！！！

上周休息周总结的中序遍历遇到了



【方法1】

题目给了思路： 节点的左子树只包含小于当前节点的数。节点的右子树只包含大于当前节点的数。而且左右子树也是二叉搜索树
由此可得=>
二叉搜索树的性质：二叉搜索树，可以通过中序遍历得到一个递增的有序序列


所以我可以中序遍历一次，然后判断这个序列是否是有序的。

【方法2】

当然也可以一步到位，在递归的时候就判断值的范围

二叉搜索树所有节点的范围都在 (left, right)

如果某个节点node的值val不在这个范围，则说明不满足条件直接返回false，否则我们将递归遍历这个左右子树，都满足才能符合。

显然第二种方法效率更高一些。

*/
<?php
/*
 * @Descripttion: 周二的题 请一天假 在家休息了一天
 * @Author: tacks321@qq.com
 * @Date: 2020-12-24 17:33:49
 * @LastEditTime: 2020-12-24 18:16:30
 */


/*
 * @lc app=leetcode.cn id=199 lang=php
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (64.41%)
 * Likes:    377
 * Dislikes: 0
 * Total Accepted:    78.5K
 * Total Submissions: 121.3K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 * 
 * 示例:
 * 
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
     * @return Integer[]
     */
    // 执行用时：0 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了31.71%的用户
    function rightSideView($root) {
        // 结果数组
        $result = [];
        $this->_helper($root, 0, $result);
        return $result;
    }

    function _helper($node, $depth, &$result)
    {
        if($node == null) {
            return false;
        }
        // 同一层，如果遍历到右子树，可以获取到值，那就不用遍历左子树
        if($depth == count($result)) {
            $result[] = $node->val;
        }
        // 不断向更深层遍历
        $this->_helper($node->right, $depth+1, $result);
        $this->_helper($node->left, $depth+1, $result);

    }


    // 层序遍历
    // 执行用时：8 ms, 在所有 PHP 提交中击败了70.73%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了19.51%的用户
    function rightSideView2($root)
    {
        $result = []; // 结果数组
        $queue  = []; // 队列
        // 边界判断
        if($root == null) {
            return $result;
        }
        // 入队
        array_push($queue, $root);
        while($count = count($queue)) {
            $tmp = '';
            for($i=0; $i<$count; $i++) {
                $node = array_shift($queue); // 先入先出
                $tmp  = $node->val;
                // 左子树
                if($node->left != null) {
                    array_push($queue, $node->left);
                }
                // 右子树
                if($node->right != null) {
                    array_push($queue, $node->right);
                }
            }
            $result[] = $tmp; // 只保留最后一个
        }
        return $result;
    }
}
// @lc code=end

/*


*/

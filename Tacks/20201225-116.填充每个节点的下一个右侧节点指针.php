<?php
/*
 * @Descripttion: 周四的题  周四是平安夜哦，女巫用了解药哈哈哈哈哈哈
 * @Author: tacks321@qq.com
 * @Date: 2020-12-25 18:37:05
 * @LastEditTime: 2020-12-25 18:49:12
 */



/*
 * @lc app=leetcode.cn id=116 lang=php
 *
 * [116] 填充每个节点的下一个右侧节点指针
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/description/
 *
 * algorithms
 * Medium (67.59%)
 * Likes:    362
 * Dislikes: 0
 * Total Accepted:    89K
 * Total Submissions: 130.4K
 * Testcase Example:  '[1,2,3,4,5,6,7]'
 *
 * 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
 * 
 * 
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 * 
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 * 
 * 初始状态下，所有 next 指针都被设置为 NULL。
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 * 
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * 
 * 输入：root = [1,2,3,4,5,6,7]
 * 输出：[1,#,2,3,#,4,5,6,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B
 * 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点的数量少于 4096
 * -1000 
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a Node.
 * class Node {
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->left = null;
 *         $this->right = null;
 *         $this->next = null;
 *     }
 * }
 */

class Solution {
    /**
     * @param Node $root
     * @return Node
     */
    // 递归处理
    // 执行用时：28 ms, 在所有 PHP 提交中击败了76.92%的用户
    // 内存消耗：19.1 MB, 在所有 PHP 提交中击败了18.13%的用户
    public function connect($root) {
        // 如果子节点没有
        if($root->left == null) {
            return $root;
        }
        // 把右节点赋值给左节点的next, 进行连接
        $root->left->next = $root->right;
        // 如果当前节点有next节点，就能根据next的节点找到对应的左节点
        // 可以画图，有点像一个倒着的提醒，就是用来夸父节点进行连接
        if($root->next != null) {
            $root->right->next = $root->next->left;
        }

        // 继续前序遍历
        $this->connect($root->left);
        $this->connect($root->right);
        return $root;
    }
}
// @lc code=end

/*
题目要求
1、常量级额外空间
2、可以用递归


可以前序遍历  处理根节点下面左右节点  寻找next节点



目前只用了一个办法 之后可以用其他的
*/
<?php
/*
 * @Descripttion: 补一下周二的题目
 * @Author: tacks321@qq.com
 * @Date: 2021-01-17 16:25:26
 */


/*
 * @lc app=leetcode.cn id=117 lang=php
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (59.22%)
 * Likes:    345
 * Dislikes: 0
 * Total Accepted:    60.9K
 * Total Submissions: 102.8K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * 给定一个二叉树
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
 * 输入：root = [1,2,3,4,5,null,7]
 * 输出：[1,#,2,3,#,4,5,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数小于 6000
 * -100 <= node.val <= 100
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start

//  Definition for a Node.
class Node {
    function __construct($val = 0) {
        $this->val = $val;
        $this->left = null;
        $this->right = null;
        $this->next = null;
    }
}


class Solution {
    /**
     * @param Node $root
     * @return Node
     */
    // BFS (层序遍历)
    // 执行用时：28 ms, 在所有 PHP 提交中击败了23.81%的用户
    // 内存消耗：18.3 MB, 在所有 PHP 提交中击败了100.00%的用户
    public function connect($root) {
        if($root == null) {
            return $root;
        }

        $queue = [$root]; // 队列
        while(!empty($queue)) {
            // 每一层的数量
            $level_count = count($queue);
            // 获取前一个结点
            $pre = new Node(0);
            // 遍历当前层
            for($i=0; $i<$level_count; $i++) {
                $node = array_shift($queue); // 先进先出
                // 当前层，pre不是空的，那么next就指向当前节点
                if($pre != null) {
                    $pre->next = $node;
                }
                // 将当前节点作为下一次的前节点
                $pre = $node;
                if($node->left != null ) {
                    array_push($queue, $node->left);
                } 
                if($node->right != null) {
                    array_push($queue, $node->right);
                }

            }
        }
        return $root;
    }

}
// @lc code=end

/*
    貌似见过这个题，应该是变型
*/
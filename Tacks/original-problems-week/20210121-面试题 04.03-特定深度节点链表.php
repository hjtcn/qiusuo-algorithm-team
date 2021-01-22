<?php

/*
 * @Descripttion: 周四冲冲冲
 * @Author: tacks321@qq.com
 * @Date: 2021-01-21 18:52:59
 * @LastEditTime: 2021-01-21 19:50:42
 */


/*
【面试题 04.03. 特定深度节点链表】

给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

示例：

输入：[1,2,3,4,5,null,7,8]

    1
    /  \ 
    2    3
    / \    \ 
4   5    7
/
8

输出：[[1],[2,3],[4,5,7],[8]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/list-of-depth-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/


/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution1 {

    public $list = []; // 结果数组

    /**
     * @param TreeNode $tree
     * @return ListNode[]
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了50.00%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了16.67%的用户
    // 递归处理 DFS
    function listOfDepth($tree) {
        // 递归处理
        $this->dfs($tree, 0);
        return $this->list;
    }

    function dfs($node, $level) {
        if($node == null) {
            return [];
        }
        //  层序遍历变形
        if(!isset($this->list[$level])) {
            $this->list[$level] = new ListNode($node->val);
        }else{
            // 链表的拼接 如果已经加入，那么就头插入节点链表
            $tmp = $this->list[$level];
            $this->list[$level] = new ListNode($node->val); // 当前值
            $this->list[$level]->next = $tmp;
        }
        // 递归处理 从右向左
        $this->dfs($node->right, $level+1);
        $this->dfs($node->left, $level+1);

    }
}

class Solution2 {

    /**
     * @param TreeNode $tree
     * @return ListNode[]
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了50.00%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了8.33%的用户
    // 迭代 BFS 广度优先遍历
    function listOfDepth($tree) {
        $list = [];
        if($tree == null){
            return $list;
        }
        // 队列
        $queue = [$tree];
        while($count = count($queue)) {
            $dummyHead = new ListNode(); // 哑节点
            $cur = $dummyHead;
            for($i=0; $i<$count; $i++) {
                $node = array_shift($queue);// 队首元素
                // 链表的生成
                $cur->next = new ListNode($node->val);
                $cur = $cur->next;

                if($node->left  != null) array_push($queue, $node->left);
                if($node->right != null) array_push($queue, $node->right);
            }
            // 每一层赋值
            $list[] = $dummyHead->next;
        }
        return $list;
    }
}

/*
[题目] 
    给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

[思考]
    某一个深度？ ==>  每一层的节点 => 层序遍历
    创建链表？   ==>  链表的生成

[考察点]
    层序遍历 生成 + 链表的 创建

    容易出错的还是，链表在生成的时候，有时候next指针会弄错 !

【DFS】
    递归处理

【BFS】
    迭代处理 ，借助栈
*/
// 给定一个二叉树

// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

// 初始状态下，所有 next 指针都被设置为 NULL。

//  

// 进阶：

// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


/*
 * @lc app=leetcode.cn id=117 lang=javascript
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function (root) {
    if(!root){
        //这里很过分，返回了一下[]就报错了。
        return root
    }
    let queue = [root]
    while (queue.length) {
        let len = queue.length
        for (let i = 0; i < len; i++) {
            let curNode = queue.shift()

            if (curNode.left) {
                queue.push(curNode.left)
            }
            if (curNode.right) {
                queue.push(curNode.right)
            }
            if (i == len - 1) {
                curNode.next = null
            }
            else {
                curNode.next = queue[0]
            }
        }
    }
    return root
};

/** 
 广搜结构还是很清晰的，每一行如果是尾部，直接添加null,如果不是结尾，则队首作为next节点
 时间复杂度：O(N)
 取决于树的节点数
 空间复杂度：O(N)
 取决于树的节点数，队列模拟遍历树节点
*/




var connect = function (root) {
    if(!root){
        return root
    }
    let dfs=(root)=>{
        if(root==null) return;
        if(root.left){
            if(root.right){
                root.left.next=root.right
            }
            else{
                root.left.next=findNext(root.next)
            }
        }

        if(root.right){
            root.right.next=findNext(root.next)
        }
        // 递归顺序，必须先右后左！！！！！！！！
        // 确保每一层从左向右扫描时能够扫描到
        dfs(root.right)
        dfs(root.left)
    }
    let findNext=(root)=>{
        while(root){         
            if(root.left) return root.left
            if(root.right) return root.right        
            root=root.next
        }
        return null
    }
    dfs(root)
    return root
};

/** 
 递归。思路就有点难搞了，上次做的题是完整二叉树会存在root.right.next=root.next.left这种情况
 而这次没有了，怎么办呢，遍历去找，直到找到root.next.next.next。。。。出现值。
 但又一个位置卡了我好久好久。就是递归顺序，必须从右向左，确保每一层从左向右能扫描到
 时间复杂度：O(N)
 取决于树的节点数。还要乘每层树叶子节点的个数。
 空间复杂度：O(1)
 本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/
// 107. 二叉树的层次遍历 II
// 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其自底向上的层次遍历为：

// [
//   [15,7],
//   [9,20],
//   [3]
// ]

/*
 * @lc app=leetcode.cn id=107 lang=javascript
 *
 * [107] 二叉树的层次遍历 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrderBottom = function(root) {
    if(!root)
    return []
    let queue=[root],res=[]
    while(queue.length){
        let len=queue.length,floorArr=[]
        //一层一层处理节点，当前节点出队，其左右子节点入队。
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            floorArr.push(curNode.val)
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
        }
        //从头部插入结果数组
        res.unshift(floorArr)
    }

    return res
};
// @lc code=end


/**
  这道题我做起来还是蛮顺利的，因为刷104二叉树的最大深度的时候，熟悉了一遍bfs
  不过也遇到了两个问题，1.首先就是返回的数组结构，我懵了一下，我以为还要维持一定的节点的结构
                     2.没有考虑root节点为空的情况
  题解：首先root节点为空，返回[];然后循环队列(while)，记录队列的长度(进行for循环)，用来进行每一层节点的出队操作。
  当然出队的同时，左右节点要及时入队。每一层for循环结束，将结果从头部unshift入结果数组，彻底跳出循环后返回结果数组
    时间复杂度：O(N)
    每个节点只会执行入队出队一次
    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 * 
 */
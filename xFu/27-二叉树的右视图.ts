// 199. 二叉树的右视图
// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 示例:

// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:

//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---


/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function rightSideView(root: TreeNode | null): number[] {
    let mapArray = [];
    let queue = [];

    if(!root) return [];

    queue.push(root);

    while(queue.length){
        // 某一次遍历开始时，表示当成层的数据个数
        let len = queue.length;
        while(len){
            const shiftQueue = queue.shift();
            // 当队列中只剩下一个数的时候（就是刚取出的数），就取这个值
            if(len === 1) mapArray.push(shiftQueue.val);
            // 遍历取到的数，将他们的左右存在的子节点添加到队列中
            shiftQueue.left && queue.push(shiftQueue.left);
            shiftQueue.right && queue.push(shiftQueue.right);
            len--;
        }
    }

    return mapArray;
};
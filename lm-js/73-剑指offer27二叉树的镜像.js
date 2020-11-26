// 剑指 Offer 27. 二叉树的镜像
// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。

// 例如输入：

//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// 镜像输出：

//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1

 

// 示例 1：

// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]


//我这个出题人出到原题啦，第二次做大家是不是飞快呀

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var mirrorTree = function(root) {
    if(!root) return root
    let left=mirrorTree(root.left)
    let right=mirrorTree(root.right)
    root.left=right
    root.right=left
    return root
};

/**
 * 递归，先递归调用记录左右节点，然后直接进行交换，返回root节点
 *  时间复杂度：O(N)
    取决于树的深度。
    空间复杂度：O(N)
    取决于树节点数，不会超过n
 */


var mirrorTree = function(root) {
    if(!root) return root
    let queue=[root]
    while(queue.length>0){
        let curNode=queue.shift()
        let swap=curNode.left
        curNode.left=curNode.right
        curNode.right=swap
        curNode.left&&queue.push(curNode.left)
        curNode.right&&queue.push(curNode.right)
    }
    return root
};

/**
 * BFS.这里踩了一个坑，提前记录了左右节点的值，进行交换和push，结果交换过程并没有作用到链表上
 *  时间复杂度：O(N)
    每个节点只会执行入队出队一次
    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 */
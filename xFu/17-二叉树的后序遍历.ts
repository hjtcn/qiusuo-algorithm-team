
// 145. 二叉树的后序遍历
// 给定一个二叉树，返回它的 后序 遍历。

// 示例:

// 输入: [1,null,2,3]  
//    1
//     \
//      2
//     /
//    3 

// 输出: [3,2,1]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？


// ============================================================
// ===                                                      ===
// ===       状态：通过,执行用时: 96 ms,内存消耗: 40.3 MB     ===
// ============================================================

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

function postorderTraversal(root: TreeNode | null): number[] {
    let nodeList: TreeNode[] = [];
    let mapList: number[] = [];

    if(!root) return [];

    nodeList.push(root);

    while(nodeList.length){

        const node = nodeList.pop();

        if(node){
            mapList.unshift(node.val);

            node.left && nodeList.push(node.left);
            node.right && nodeList.push(node.right); 
        }
    }

    return mapList;
};
// 剑指 Offer 54. 二叉搜索树的第k大节点
// 给定一棵二叉搜索树，请找出其中第k大的节点。

 

// 示例 1:

// 输入: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// 输出: 4
// 示例 2:

// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// 输出: 4
 

// 限制：

// 1 ≤ k ≤ 二叉搜索树元素个数


// ============================================================
// ===                     自己想的代码                      ===
// ===       状态：通过,执行用时: 112 ms,内存消耗: 45.5 MB     ===
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

function kthLargest(root: TreeNode | null, k: number): number {
    let numList:number[] = [];

    function mapTree (node: TreeNode | null){
        if(node !== null){
            numList.push(node.val)
            mapTree(node.left);
            mapTree(node.right);
        }
    }

    mapTree(root);
    numList = numList.sort((a,b) => b-a);
    return numList[k-1];
};
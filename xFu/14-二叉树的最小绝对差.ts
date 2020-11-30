
// 530. 二叉搜索树的最小绝对差
// 给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

 

// 示例：

// 输入：

//    1
//     \
//      3
//     /
//    2

// 输出：
// 1

// 解释：
// 最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 

// 提示：

// 树中至少有 2 个节点。
// 本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同




// ============================================================
// ===                     自己想的代码                      ===
// ===       状态：通过,执行用时: 288 ms,内存消耗: 45.7 MB     ===
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

function getMinimumDifference(root: TreeNode | null): number {
    let nodeList:number[] = [];
    let min = Infinity;

    function mapTree (node: TreeNode | null) {
        if(node !== null){

            // 遍历所有已处理过的结点，找到绝对值最小的一个值
            nodeList.forEach(nodeNum => {
                const abs = Math.abs(node.val - nodeNum)
                min = abs < min ? abs : min;
            });

            // 插入当前结点
            nodeList.push(node.val);
            
            mapTree(node.left);
            mapTree(node.right);
        }
    }

    mapTree(root);

    return min;

};
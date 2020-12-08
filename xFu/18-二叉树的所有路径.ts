
// 257. 二叉树的所有路径
// 给定一个二叉树，返回所有从根节点到叶子节点的路径。

// 说明: 叶子节点是指没有子节点的节点。

// 示例:

// 输入:

//    1
//  /   \
// 2     3
//  \
//   5

// 输出: ["1->2->5", "1->3"]

// 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3


// ============================================================
// ===                                                      ===
// ===       状态：通过,执行用时: 84 ms,内存消耗: 40.4 MB     ===
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

function binaryTreePaths(root: TreeNode | null): string[] {
    let pathList:string[] = [];

    function mapTree(node: TreeNode | null, valueList: number[]){
        if(node){
            valueList.push(node.val);
            if(node.left || node.right){
                node.left && mapTree(node.left, [ ...valueList ]);
                node.right && mapTree(node.right, [ ...valueList ]);
            }else{
                if(valueList.length === 1){
                    pathList.push(''+valueList[0]);
                }else if(valueList.length > 1){
                    pathList.push(valueList.join('->'));
                }
            }
        }
    }

    mapTree(root, []);

    return pathList;
};
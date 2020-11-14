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


// ============================================================
// ===                     自己想的代码                      ===
// ===       状态：通过,执行用时: 88 ms,内存消耗: 40.9 MB     ===
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

interface TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
 }

function levelOrderBottom(root: TreeNode | null): number[][] {
    let mapTree:number[][] = [];

    function getTreeValue (level: number, node: TreeNode | null){
        if(node !== null){
            // 若数组的长度大于等于遍历的层级，则表示当前遍历到的层级已经在别的子树中创建过了，就需要倒序找到需要修改的位置
            if(mapTree.length >= level){
                mapTree[mapTree.length - level].push(node.val);
            // 若数组的长度小于遍历的层级，就表示当前遍历的层级下还没有创新新的数组，就需要创建一个新的数组
            }else{
                mapTree.unshift([node.val]);
            }

            level++;
            getTreeValue(level, node.left);
            getTreeValue(level, node.right);
        }
    }

    getTreeValue(1, root);

    return mapTree;
};
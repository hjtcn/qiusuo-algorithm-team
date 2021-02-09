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

function pathSum(root: TreeNode | null, sum: number): number[][] {
    let result:number[][] = [];

    function mapTree(node: TreeNode | null, link = []) {
        if(node){
            link.push(node.val);
            if(node.left || node.right){
                mapTree(node.left, link);
                mapTree(node.right, link);
            }else{
                const prevTotal = link.reduce((a, b) => a+b, 0);
                if(prevTotal === sum){
                    result.push([...link]);
                }
            }
            // 删除按当前结点，实现回溯
            link.pop()
        }
    }

    mapTree(root);
    return result;
};
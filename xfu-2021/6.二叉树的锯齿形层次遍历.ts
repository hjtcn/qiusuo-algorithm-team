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

function zigzagLevelOrder(root: TreeNode | null): number[][] {
    const mapResult:number[][] = [];

    if(root){
        mapResult.push([root.val]);
        mapLayout(root);
    }
    
    function mapLayout(node: TreeNode | null, flag: 'right' | 'left' = 'right', level = 1) {
        if(node) {
            let layout = mapResult[level] || [];
            const reverseFlag = flag === 'left' ? 'right' : 'left'
            const startNode = node.left, endNode = node.right;

            // 判断是需要如何操作数据
            if(flag === 'right'){
                startNode && layout.unshift(startNode.val);
                endNode && layout.unshift(endNode.val);
            }else{
                startNode && layout.push(startNode.val);
                endNode && layout.push(endNode.val);
            }
            
            if(startNode || endNode){
                mapResult[level] = layout;
            }

            startNode && mapLayout(startNode, reverseFlag, level + 1);
            endNode && mapLayout(endNode, reverseFlag, level + 1);
        }
    }
    return mapResult;
};
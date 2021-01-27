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

function increasingBST(root: TreeNode | null): TreeNode | null {
    if(!root) return null;

    let prevNode = increasingBST(root.left);
    
    const currentNode = new TreeNode(root.val);

    let nextNode = increasingBST(root.right);

    // 将右子树的遍历结果挂载到当前结点上
    currentNode.right = nextNode;

    if(prevNode){
        let tempNode = prevNode;
        // 遍历链表，找到最后一个结点
        while(tempNode.right){
            tempNode = tempNode.right;
        }
        // 在最后一个结点后面挂载当前结点
        tempNode.right = currentNode;
    }else{
        // 左子树的遍历结果为空，就将当前结点作为主结点
        prevNode = currentNode;
    }

    return prevNode;
};
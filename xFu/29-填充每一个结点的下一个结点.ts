/**
 * Definition for Node.
 * class Node {
 *     val: number
 *     left: Node | null
 *     right: Node | null
 *     next: Node | null
 *     constructor(val?: number, left?: Node, right?: Node, next?: Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function connect(root: Node | null): Node | null {
    // 题设为完美二叉树，但是为了TS检查还是需要进行判断
    if(root && root.left && root.right){
        // 任一结点的左结点的下一个一定是此结点的右结点
        // 任一结点的右结点的下一个是，如果当前结点有右结点，则右结点就是当前结点的下一个结点的左结点
        root.left.next = root.right;
        if(root.next)
        root.right.next = root.next.left;

        // 递归判断左右子树
        connect(root.left);
        connect(root.right);
    }
    
    return root;
};
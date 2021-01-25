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

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function listOfDepth(tree: TreeNode | null): Array<ListNode | null> {
    const queue: TreeNode[] = [];
    let deepArr:Array<ListNode | null> = [];
    let index = 0;

    queue.push(tree);

    while(queue.length) {
        let len = queue.length;
        
        let pNode;
        while(len--){
            const lNode = new ListNode();
            const node = queue.shift();

            lNode.val = node.val;

            if(pNode){
                // 上一个结点的next是当前
                pNode.next = lNode;
            }else{
                // 保存第一个结点位置
                deepArr.push(lNode);
            }

            // 将当前结点设置为下一个结点的‘上一个结点’
            pNode = lNode;

            node.left && queue.push(node.left);
            node.right && queue.push(node.right);
        }
        
    }

    return deepArr;
};
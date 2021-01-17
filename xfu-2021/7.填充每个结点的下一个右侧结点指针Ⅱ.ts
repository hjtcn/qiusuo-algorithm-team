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
    let nodeList: Node[] = [];

    function mapNode (node: Node | null, level: number) {
        // node必须不存在才可以进行处理
        if(!node) return;
        const levelItem = nodeList[level];

        // 当前层级下有结点信息了，就表示当前结点层级左侧有一个未处理的结点，就把那个结点的next指向自身
        if(levelItem){
            levelItem.next = node;
        }
        
        // 最后将当前层级的结点替换为当前结点
        nodeList[level] = node;

        // 递归判断左、右结点
        mapNode(node.left, level + 1);
        mapNode(node.right, level + 1);
    }

    mapNode(root, 0);
    
    return root;
};
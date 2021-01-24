// 思路1

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

function isSymmetric(root: TreeNode | null): boolean {
    let isSymmetric = true;

    function mapTree (nodeLeft: TreeNode | null, nodeRight: TreeNode | null) {
        // 如果已经被标记过了不对称，或者结点的左、右结点都不存在就不继续进行了
        if(!isSymmetric || (!nodeRight && !nodeLeft)) return;

        const leftVal = nodeLeft ? nodeLeft.val : null;
        const rightVal = nodeRight ? nodeRight.val : null;
        // 判断若两个对称位置的值不相同就直接标记为false
        if(leftVal !== rightVal){
            isSymmetric = false;
            return;
        }else{
            mapTree(nodeLeft ? nodeLeft.left : null, nodeRight ? nodeRight.right : null);
            mapTree(nodeLeft ? nodeLeft.right : null, nodeRight ? nodeRight.left : null);
        }
    }

    mapTree(root ? root.left : null, root ? root.right : null);

    return isSymmetric;
};

// 思路2


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

function isSymmetric1(root: TreeNode | null): boolean {
    const nodeList:string[] = [];

    function mapTree (node: TreeNode | null, level = 0) {
        if(node){
            node.left && mapTree(node.left, level+1);
            nodeList.push(''+node.val+level);
            node.right && mapTree(node.right, level+1);
        }
    }

    mapTree(root);

    let start = 0;
    let end = nodeList.length-1;
    // 从两头开始判断，当start比end大或相等的时候说明已经判断完了
    while(start < end){
        if(nodeList[start] !== nodeList[end]) return false;

        start++;
        end--;
    }

    return true;
};
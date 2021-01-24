// 解法1

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

class BSTIterator {
    constructor(root: TreeNode | null) {
        this.root = root;
        this.nextMinNode = null;
        this.mapTree(root);
    }

    root: TreeNode | null;

    nextMinNode: TreeNode | null;

    /**
     * 计算下一个最小值
     */
    mapTree (node: TreeNode | null, pNode?: TreeNode | null, isEnd = true): boolean {
        // 左、中、右
        if(node){
            if(node.left) {
                // 左子树找到了最小值就不进行判断了，因为中、后的值一定会更大
                if(this.mapTree(node.left, node, false)){
                    return true;
                }
            };
            
            // 最小值不存在，当前结点就是最小值
            if(!this.nextMinNode){
                this.nextMinNode = node;
                return true;
            }else{
                const minVal = this.nextMinNode.val;
                const nodeVal = node.val;
                
                // 当前结点的值大于最小结点的值，那么当前结点就是需要的下一个最小结点
                // 当结点值和最小结点的值相等时就判断是否是最后一个结点了，是的话就没有下一个结点了
                if(nodeVal > minVal){
                    this.nextMinNode = node;
                    return true;
                }else if(nodeVal === minVal) {
                    if(isEnd || !pNode) {
                        this.nextMinNode = null;
                    }
                }
            }
            
            if(node.right) {
                if(this.mapTree(node.right, node)){
                    return true;
                }
            };

        }
    }

    next(): number {
        const node = this.nextMinNode;
        this.mapTree(this.root);

        return node && node.val;
    }

    hasNext(): boolean {
        return !!this.nextMinNode;
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */


 // 解法2

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

class BSTIterator1 {
    constructor(root: TreeNode | null) {
        this.arr = this.mapTree(root);
    }

    arr: number[] = [];

    /**
     * 遍历树，获取搜索树数组
     */
    mapTree (node: TreeNode | null): number[] {
        if(node) {
            return [...this.mapTree(node.left), node.val, ...this.mapTree(node.right)]
        }else{
            return [];
        }
    }

    next(): number {
        return this.arr.shift();
    }

    hasNext(): boolean {
        return !!this.arr.length;
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */
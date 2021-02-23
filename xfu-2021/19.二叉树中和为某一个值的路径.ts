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
    const result = [];
    // DFS
    const mapTree = function (node: TreeNode | null, link = []) {
        if(!node) return;

        link.push(node.val);
        const total = link.reduce((prev, next) => prev + next, 0);

        if(node.left || node.right){
            
            /**
             * 这里试图在每次赋值之后，判断一下结果是否符合要求
             * 期间遇到了 val值为负数，需要判断小值。val有正有负，未完整解决
             */
            // 不是叶结点，就判断添加后的总数是否超过了最大值，超过了就结束
            // if(Math.abs(total) > Math.abs(sum)) return;
            mapTree(node.left, [...link]);
            mapTree(node.right, [...link])
        }else{
            // 叶结点就判断总数是否等于最大值
            if(total === sum) result.push(link);
        }
    }

    mapTree(root, []);
    return result;
};

/**
 * 做题的时候想了挺多东西,但是调试通不过去，最终修改完了发现和自己原来的代码差不多 :(
 * 
 * 有时候做算法题也会让我有个想法，用某一个方式或语言特性实现一个什么功能，虽然最终可能实现不了，但这种过程也会激励让我继续坚持把
 * 
 * 继续加油吧~
 */

/**
 * 做完题太晚了，没来得及写MarkDown
 * 写的时候，想多做点和之前不一样的，想了两个点：
 * 1. 不全部遍历，只要某一次遍历的结果不符合要求了，就直接结束
 * 2. 使用ES6的Generator函数，尝试能不能让一个递归遍历提前终止（感觉可以，又感觉不可以，有时间了要尝试实践下）
 * 今天（:(已经是第二天了）下班了再解决上方的思考问题，补充MarkDown
 */
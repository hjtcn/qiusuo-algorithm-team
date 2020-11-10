// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。

// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3 。


/*
 * @lc app=leetcode.cn id=104 lang=javascript
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
//递归，比较左右子树的最大深度。
var maxDepth = function (root) {
    if (!root)
        return 0
    //递归真牛呀，哈哈哈
    return 1 + Math.max(maxDepth(root.left), maxDepth(root.right))
};
// @lc code=end

/*
    做树的题，一旦理解递归的思路，确实满嗨的，代码量也少
    比较左右子树的最大深度。递归更新最大值
    复杂度分析：
    时间复杂度：O(N)
    每一个节点在递归中只会被遍历一次

    空间复杂度：O(N)
    N取决与树的深度
*/



//BFS
var maxDepth = function (root) {
    if (!root) return 0
    let queue = [root], res = 1
    while (queue.length) {
        //当前层节点个数
        let len=queue.length
        //将每一层的节点出队
        for(let i=0;i<len;i++){
            let curNode = queue.shift()
            //左右节点入队
            if (curNode.left) {
                queue.push(curNode.left)
            }
            if (curNode.right) {
                queue.push(curNode.right)
            }
        }
        //节点都出队，队列还不为空，则说明有子子树
        if(queue.length){
            res++;
        }
    }
    return res
}

/*
    广度优先搜索，横向记录节点个数，将当前层每个节点出队，然后其左右节点入队。
    一次循环结束，队列不为空，说明还有子子树
    
    复杂度分析：
    时间复杂度：O(N)
    每一个节点只会被遍历一次

    空间复杂度：O(N)
    N取决于队列存储元素的数量，最坏情况会达到O(n)
*/
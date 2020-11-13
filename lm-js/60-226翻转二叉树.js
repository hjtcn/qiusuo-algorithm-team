// 226. 翻转二叉树
// 翻转一棵二叉树。

// 示例：

// 输入：

//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// 输出：

//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1

/*
 * @lc app=leetcode.cn id=226 lang=javascript
 *
 * [226] 翻转二叉树
 */

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
//BFS
var invertTree = function(root) {  
    if(!root)
    return root;////这个地方踩坑了，直接返回[]报错，误以为返回类型数组即可了
    let queue=[root],res=[]
    while(queue.length>0){
            let curNode=queue.shift()
            //交换过程
            let flag=curNode.left
            curNode.left=curNode.right
            curNode.right=flag
            //队列push左右子树，等待下一次它作为curNode
            curNode.left&&queue.push(curNode.left)
            curNode.right&&queue.push(curNode.right)
    }
    return root
};
// @lc code=end

/*
    bfs.刚开始写的有点复杂了，while循环中又加入了for循环，想要一层一层交换。
    后来AC之后拨丝抽茧。
    题解：发现这个翻转只需要两步：1.当前元素的左右子树交换
                             2.队列push左右子树，更新队列头，重获当前元素
    时间复杂度：O(N)
    每个节点只会执行入队出队一次
    空间复杂度：O(N)
    取决于队列节点数，不会超过n


*/

//递归
var invertTree = function(root) {
    //如果节点为空，返回null,看题解了解的边界判断
    if(root==null)
    return null;
    //递归获取左右子树
    let right=invertTree(root.right)
    let left=invertTree(root.left)
    //开始交换喽
    root.left=right
    root.right=left
    //返回原始节点，有点链表的意思了。
    return root
};

/**
    递归。做这道题时我知道怎么交换，但是想不到怎么去深层调用，推动递归的产生了。
    递归题就应该以小见大，直接将左右子节点进行调用invertTree方法，返回翻转左右子节点后的节点。
    时间复杂度：O(N)
    取决于树的深度。
    空间复杂度：O(N)
    取决于树节点数，不会超过n
 * 
 */
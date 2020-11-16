// 101. 对称二叉树
// 给定一个二叉树，检查它是否是镜像对称的。

 

// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
 

// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

//     1
//    / \
//   2   2
//    \   \
//    3    3
 

// 进阶：

// 你可以运用递归和迭代两种方法解决这个问题吗？


/*
 * @lc app=leetcode.cn id=101 lang=javascript
 *
 * [101] 对称二叉树
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
 * @return {boolean}
 */

 //递归
var isSymmetric = function(root) {
    if(!root)
    return true
    //定义一个check方法，对比左右节点。
    //核心为左右节点value相等&&左节点的左孩子==右节点的右孩子&&左节点的右孩子==右节点的左孩子
    const check=(left,right)=>{
        if(!left&&!right){
            return true
        }
        if(left&&right){
            return left.val==right.val&&check(left.left,right.right)&&check(left.right,right.left)
        }
        return false
    }
    return check(root.left,root.right)
};
/**
  题解：递归，首先找到可以递归的点.
  1.如果root为空，返回true
  2.如果root有左右子节点，则左右子节点相等才返回true
  3.如果root的左右子节点有左右子节点，则root的左右节点值相等&&左节点的左孩子==右节点的右孩子&&左节点的右孩子==右节点的左孩子，才会返回true
  ...
  总结。检查左右节点是否存在且相等&&左节点的左孩子==右节点的右孩子&&左节点的右孩子==右节点的左孩子，才会返回true
    时间复杂度：O(N)
    取决于树的最长深度
    空间复杂度：O(N)
    取决于树节点数，不会超过n
 * 
 */
// @lc code=end

var isSymmetric = function(root) {
    if(!root)
    return true
    let queue=[root,root]
    while(queue.length>0){
        let left=queue.shift(),right=queue.shift()
        if(!left&&!right){
            continue
        }
        if(!left||!right||left.val!==right.val){
            return false
        }
        //这个地方很酷，自动构建队列
        queue.push(left.left,right.right,left.right,right.left)
    }
    return true
}

/**
  bfs.这个解法是看题解解决的，我觉得挺巧妙的，自己去构建队列格式，让需要对比的节点紧挨着放入队列。
  如：构建原队列，还自动填充了[root,root],而核心对比即左节点的左孩子==右节点的右孩子&&左节点的右孩子==右节点的左孩子
  故构建队列时，queue.push(left.left,right.right,left.right,right.left)

    时间复杂度：O(N)
    取决于树的最长深度。
    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 * 
 */
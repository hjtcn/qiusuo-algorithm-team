// 222. 完全二叉树的节点个数
// 给出一个完全二叉树，求出该树的节点个数。

// 说明：

// 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

// 示例:

// 输入: 
//     1
//    / \
//   2   3
//  / \  /
// 4  5 6

// 输出: 6

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
var countNodes = function(root) {
    if(!root){
        return 0
    }
    let queue=[root],res=0
    while(queue.length>0){
        let len=queue.length
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            curNode.left&&queue.push(curNode.left)
            curNode.right&&queue.push(curNode.right)
            res++;    
        }
    }
    return res
};

/*
    bfs用的越来越顺手啦，啦啦啦啦

    完全二叉树，就一层每个节点入一次队，for循环内加1就好啦


    时间复杂度：O(N)
    取决于树的节点数
    
    空间复杂度：O(N)
    取决于树节点数
*/

var countNodes = function(root) {
    if(!root){
        return 0
    }
    return 1+countNodes(root.left)+countNodes(root.right)
};

/*
    嘿嘿，今天的题有点简单哟，得瑟得瑟，因为是完全二叉树，就不用考虑很多边界问题

    递归。计算当前节点个数(1)+左子树的个数(countNodes(root.left))+右子树的个数(countNodes(root.right))


    时间复杂度：O(N)
    取决于树的节点数
    
    空间复杂度：O(N)
    取决于树节点数
*/


//看到题解上有二分的方法了，但是木有看懂，随后问问各位老友，再进行补充
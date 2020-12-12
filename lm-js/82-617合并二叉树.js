// 617. 合并二叉树
// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

// 示例 1:

// 输入: 
// 	Tree 1                     Tree 2                  
//           1                         2                             
//          / \                       / \                            
//         3   2                     1   3                        
//        /                           \   \                      
//       5                             4   7                  
// 输出: 
// 合并后的树:
// 	     3
// 	    / \
// 	   4   5
// 	  / \   \ 
// 	 5   4   7
// 注意: 合并必须从两个树的根节点开始。

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} t1
 * @param {TreeNode} t2
 * @return {TreeNode}
 */
var mergeTrees = function(t1, t2) {
        if(!t1) return t2
        if(!t2) return t1
        let root=new TreeNode(t1.val+t2.val)
        root.left=mergeTrees(t1.left,t2.left)
        root.right=mergeTrees(t1.right,t2.right)
        return root
};
/** 
 * 递归的思路还是非常容易想到的，熟能生巧
 * 时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数
*/


var mergeTrees = function(t1, t2) {
    if(!t1||!t2) return t1||t2
    let queue=[[t1,t2]]
    while(queue.length){
        let [h1,h2]=queue.shift()
        let {left:l1,right:r1}=h1
        let {left:l2,right:r2}=h2
        if(h1&&h2) h1.val=h1.val+h2.val
        if(l1&&l2) queue.push([l1,l2])
        if(r1&&r2) queue.push([r1,r2])
      
        if(!l1&&l2) h1.left=h2.left
        if(!r1&&r2) h1.right=h2.right
    }
    return t1
};
/**
 * BFS这个思路是看题解发现还是蛮巧妙的。
 * 首先t1，t2节点出现一个为空，则返回t1||t2
 * 然后再开始队列模式，广搜的框架其实都熟悉的差不多了，
 * 但是自己写的时候没想到将两个节点放到数组里进行并列对比，而是想到设置两个队列，从这个时候就有点跑偏了，最后没弄出来
 * 当然这里es6的解构赋值用的也是很巧妙，其实它们也相当于声明了变量。
 * 接下来才是核心逻辑，如果两个节点都有，将h2的值加到h1上；h1没有，h2有，将h2赋值给h1;然后控制左右节点来进行队列的入队出队，最后返回t1节点
 * 合并核心就在于将t2链表放到t1上，然后返回t1
 * 时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数
 */
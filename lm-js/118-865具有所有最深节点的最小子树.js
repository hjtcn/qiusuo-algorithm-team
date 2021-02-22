// 865. 具有所有最深节点的最小子树
// 给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。

// 如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。

// 一个节点的 子树 是该节点加上它的所有后代的集合。

// 返回能满足 以该节点为根的子树中包含所有最深的节点 这一条件的具有最大深度的节点。

 

// 注意：本题与力扣 1123 重复：https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/

 

// 示例 1：



// 输入：root = [3,5,1,6,2,0,8,null,null,7,4]
// 输出：[2,7,4]
// 解释：
// 我们返回值为 2 的节点，在图中用黄色标记。
// 在图中用蓝色标记的是树的最深的节点。
// 注意，节点 5、3 和 2 包含树中最深的节点，但节点 2 的子树最小，因此我们返回它。
// 示例 2：

// 输入：root = [1]
// 输出：[1]
// 解释：根节点是树中最深的节点。
// 示例 3：

// 输入：root = [0,1,3,null,2]
// 输出：[2]
// 解释：树中最深的节点为 2 ，有效子树为节点 2、1 和 0 的子树，但节点 2 的子树最小。
 

// 提示：

// 树中节点的数量介于 1 和 500 之间。
// 0 <= Node.val <= 500
// 每个节点的值都是独一无二的。

/*
思路：仔细读读题意--具有所有最深节点的最小子树
     1.首先最深节点在哪？
     2.到达最深节点路径中都有那些根节点？
     3.距离可以通往所有最深节点的(最近的)根在哪？
回到解法：也很巧妙，不考虑那么多，就考虑递归左右子节点。
        1.左深==右深，返回root根节点
        2.左深<右深，返回右节点
        3.左深>右深，返回左节点。

自己写的时候一直想用bfs去做一遍，但是倒着去找父亲的过程，总是想不到怎么处理，唉
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var subtreeWithAllDeepest = function(root) {
    if(!root) return root
    let dfs=(root,depth)=>{
        if(!root) return [root,depth];
        const [lRoot,lDepth]=dfs(root.left,depth+1);
        const [RRoot,RDepth]=dfs(root.right,depth+1);
        if(lDepth==RDepth) return [root,lDepth]
        if(lDepth>RDepth) return [lRoot,lDepth]
        if(lDepth<RDepth) return [RRoot,RDepth]
    }
    return dfs(root,-1)[0]
};


/*
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数
*/
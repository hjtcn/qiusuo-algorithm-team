/*
538. 把二叉搜索树转换为累加树
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同

 

示例 1：



输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
示例 2：

输入：root = [0,null,1]
输出：[1,null,1]
示例 3：

输入：root = [1,0,2]
输出：[3,3,2]
示例 4：

输入：root = [3,2,4,1]
输出：[7,9,4,10]
 

提示：

树中的节点数介于 0 和 104 之间。
每个节点的值介于 -104 和 104 之间。
树中的所有值 互不相同 。
给定的树为二叉搜索树。
*/


/*
    思路：这个题首先是看图，了解它属于反向中序遍历。

    然后就拿出我们的框架。递归还好。迭代我又给忘记了。每次一用迭代，脑子就冲到层序遍历上了。多练习练习。
    题目核心：
    1:右根左
    2:记录和
    3:node进行val赋值

*/

var convertBST = function(root) {
    let prevNode=new TreeNode()
    prevNode.next=root
    let sum=0
    let dfs=(root)=>{
        if(!root){
            return 0
        }
        dfs(root.right)
        sum+=root.val
        root.val=sum
        dfs(root.left)
        return root   
    }
    dfs(root)
    return prevNode.next
};
// @lc code=end

var convertBST = function(root) {
    if(!root) return null
    let queue=[]
    node=root,sum=0
    while(node||queue.length){
        while(node){
            queue.push(node)
            node=node.right
        }
        node=queue.pop()
        sum+=node.val
        node.val=sum
        node=node.left
    } 
    return root
  }
  /*
    时间复杂度：O(N)
    空间复杂度：O(N)
  */

  /*
    思考：一旦写的少就忘记了。多练习
  */